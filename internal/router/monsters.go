package router

import (
	"dqix/internal/types"
	gin_utils "dqix/pkg/gin"
	"dqix/pkg/htmx"
	"dqix/web/templ/components/base"
	"dqix/web/templ/pages"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *app) MonsterRoutes(engine *gin.Engine) {
	router := engine.Group("/monsters")

	router.GET("/")
	router.GET("/:family", a.MonsterFamilyHandler)
	router.GET("/:family/:id", a.MonsterHandler)
}

func (a *app) MonsterFamilyHandler(ctx *gin.Context) {
	familyId := ctx.Param("family")

	monsters := a.data.monsterMap.GetFamilySlice(familyId, ctx.Query("sort"))

	pageTitle := "DQIX | " + types.ToFamilyTitle(familyId) + " Family"
	htmx.SetTitle(ctx, pageTitle)
	htmx.SetIcon(ctx, "/static/favicon.ico")
	params := pages.MonsterFamilyParams{
		Family:          familyId,
		FamilyTitle:     types.ToFamilyTitle(familyId),
		Monsters:        monsters,
		DisplayMode:     ctx.Query("display"),
		SortPathGetter:  types.PrepareSimpleSortPath(*ctx.Request.URL),
		SortOrderGetter: types.GetSortOrder(ctx.Request.URL),
		LayoutParams: base.LayoutParams{
			PageTitle:  pageTitle,
			Page:       familyId,
			IsDarkMode: gin_utils.IsDarkMode(ctx),
			CSSVersion: a.cssVersion,
		},
	}

	switch htmx.GetHxSwapTarget(ctx) {
	// case "monster-table":
	// 	ctx.HTML(http.StatusOK, "", pages.MonsterFamilyTable(params))
	case "page-content":
		ctx.HTML(http.StatusOK, "", pages.MonsterFamilyContent(params))
	case "sidenav-page-wrapper":
		if htmx.HasMatchingPath(ctx) || htmx.IsRequestingParentPath(ctx) {
			htmx.Retarget(ctx, "#page-content")
			htmx.Reswap(ctx, "innerHTML")
			ctx.HTML(http.StatusOK, "", pages.MonsterFamilyContent(params))
		} else {
			ctx.HTML(http.StatusOK, "", pages.MonsterFamilyContentWithSideNav(params))
		}
	default:
		ctx.HTML(http.StatusOK, "", pages.MonsterFamilyPage(params))
	}
}

func (a *app) MonsterHandler(ctx *gin.Context) {
	familyId := ctx.Param("family")
	id := ctx.Param("id")
	monster := a.data.monsterMap.GetMonster(familyId, id)

	pageTitle := "DQIX | " + monster.Title
	htmx.SetTitle(ctx, pageTitle)
	htmx.SetIcon(ctx, "/static/favicon.ico") // htmx.SetIcon(ctx, inventory.ImageSrc())
	params := pages.MonsterParams{
		Monster: monster,
		Getter:  a.data.GetQuickThing,
		LayoutParams: base.LayoutParams{
			PageTitle:  pageTitle,
			Page:       monster.GetFamilyID(),
			IsDarkMode: gin_utils.IsDarkMode(ctx),
			CSSVersion: a.cssVersion,
		},
	}

	switch htmx.GetHxSwapTarget(ctx) {
	case "page-content":
		if htmx.HasMatchingParentPath(ctx) {
			ctx.HTML(http.StatusOK, "", pages.MonsterContent(params))
		} else {
			htmx.Retarget(ctx, "#sidenav-page-wrapper")
			ctx.HTML(http.StatusOK, "", pages.MonsterContentWithSideNav(params))
		}
	case "sidenav-page-wrapper":
		ctx.HTML(http.StatusOK, "", pages.MonsterContentWithSideNav(params))
	default:
		ctx.HTML(http.StatusOK, "", pages.MonsterPage(params))
	}
}
