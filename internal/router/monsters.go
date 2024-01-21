package router

import (
	"dqix/internal/types"
	gin_utils "dqix/pkg/gin"
	"dqix/pkg/htmx"
	"dqix/web/templ/components/base"
	"dqix/web/templ/pages"
	"net/http"
	"strings"

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

	pageTitle := "DQIX | " + strings.Title(familyId) + " Family"
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
		ctx.HTML(http.StatusOK, "", pages.MonsterFamilyContentWithSideNav(params))
	default:
		ctx.HTML(http.StatusOK, "", pages.MonsterFamilyPage(params))
	}
}

func (a *app) MonsterHandler(ctx *gin.Context) {
	ctx.Status(http.StatusNotImplemented)
}
