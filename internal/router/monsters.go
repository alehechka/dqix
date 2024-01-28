package router

import (
	"dqix/internal/types"
	"dqix/internal/types/params"
	gin_utils "dqix/pkg/gin"
	"dqix/pkg/htmx"
	"dqix/web/templ/pages"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *app) MonsterRoutes(engine *gin.Engine) {
	router := engine.Group("/monsters")

	router.GET("/")
	router.GET("/:family", a.MonsterFamilyWrapper(MonsterFamilyRenderer))
	router.GET("/:family/:id", a.MonsterWrapper(MonsterRenderer))
}

func (a *app) MonsterFamilyWrapper(handler func(*gin.Context, params.MonsterFamily)) func(*gin.Context) {
	return func(ctx *gin.Context) {
		familyId := ctx.Param("family")
		monsters := a.data.monsterMap.GetFamilySlice(familyId, ctx.Query("sort"))

		params := params.MonsterFamily{
			Family:          familyId,
			FamilyTitle:     types.ToFamilyTitle(familyId),
			Monsters:        monsters,
			DisplayMode:     ctx.Query("display"),
			SortPathGetter:  types.PrepareSimpleSortPath(*ctx.Request.URL),
			SortOrderGetter: types.GetSortOrder(ctx.Request.URL),
			LayoutParams: params.Layout{
				PageTitle:  "DQIX | " + types.ToFamilyTitle(familyId) + " Family",
				Page:       familyId,
				IsDarkMode: gin_utils.IsDarkMode(ctx),
				CSSVersion: a.cssVersion,
			},
		}

		handler(ctx, params)
	}
}

func MonsterFamilyRenderer(ctx *gin.Context, params params.MonsterFamily) {
	htmx.SetTitle(ctx, params.LayoutParams.PageTitle)
	htmx.SetIcon(ctx, params.LayoutParams.GetIconPath())

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

func (a *app) MonsterWrapper(handler func(*gin.Context, params.Monster)) func(*gin.Context) {
	return func(ctx *gin.Context) {
		familyId := ctx.Param("family")
		id := ctx.Param("id")
		monster := a.data.monsterMap.GetMonster(familyId, id)

		params := params.Monster{
			Monster: monster,
			Getter:  a.data.GetQuickThing,
			LayoutParams: params.Layout{
				PageTitle:  "DQIX | " + monster.Title,
				Page:       monster.GetFamilyID(),
				IsDarkMode: gin_utils.IsDarkMode(ctx),
				CSSVersion: a.cssVersion,
			},
		}

		handler(ctx, params)
	}
}

func MonsterRenderer(ctx *gin.Context, params params.Monster) {
	htmx.SetTitle(ctx, params.LayoutParams.PageTitle)
	htmx.SetIcon(ctx, params.LayoutParams.GetIconPath())

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
