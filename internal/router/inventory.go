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

func (a *app) InventoryRoutes(engine *gin.Engine) {
	router := engine.Group("/inventory")

	router.GET("/")
	// router.GET("/:type")
	// router.GET("/:type/:category")
	router.GET("/:type/:category/:classification", a.ClassificationHandler)

	router.GET("/:type/:category/:classification/:id", a.InventoryHandler)
}

func (a *app) ClassificationHandler(ctx *gin.Context) {
	typeId := ctx.Param("type")
	category := ctx.Param("category")
	classification := ctx.Param("classification")
	inventories := a.data.inventoryMap.GetClassificationSlice(typeId, category, classification, ctx.Query("sort"))

	// TODO add some utility function that checks if the request Accept's JSON (and/or weighted with others)
	if ctx.GetHeader("Accept") == "application/json" {
		ctx.JSON(http.StatusOK, inventories)
		return
	}

	pageTitle := "DQIX | " + strings.Title(classification)
	htmx.SetTitle(ctx, pageTitle)
	htmx.SetIcon(ctx, "/static/favicon.ico")
	params := pages.InventoryClassificationParams{
		Classification:  classification,
		Inventories:     inventories,
		Stats:           inventories.GetHasInventoryStats(),
		DisplayMode:     ctx.Query("display"),
		SortPathGetter:  types.PrepareSimpleSortPath(*ctx.Request.URL),
		SortOrderGetter: types.GetSortOrder(ctx.Request.URL),
		LayoutParams: base.LayoutParams{
			PageTitle:  pageTitle,
			Page:       classification,
			IsDarkMode: gin_utils.IsDarkMode(ctx),
			CSSVersion: a.cssVersion,
		},
	}

	switch htmx.GetHxSwapTarget(ctx) {
	case "inventory-table":
		ctx.HTML(http.StatusOK, "", pages.InventoryClassificationTable(params))
	case "page-content":
		ctx.HTML(http.StatusOK, "", pages.InventoryClassificationContent(params))
	case "sidenav-page-wrapper":
		ctx.HTML(http.StatusOK, "", pages.InventoryClassificationContentWithSideNav(params))
	default:
		ctx.HTML(http.StatusOK, "", pages.InventoryClassificationPage(params))
	}
}

func (a *app) InventoryHandler(ctx *gin.Context) {
	typeId := ctx.Param("type")
	category := ctx.Param("category")
	classification := ctx.Param("classification")
	id := ctx.Param("id")
	inventory := a.data.inventoryMap.GetInventory(typeId, category, classification, id)

	// TODO add some utility function that checks if the request Accept's JSON (and/or weighted with others)
	if ctx.GetHeader("Accept") == "application/json" {
		ctx.JSON(http.StatusOK, inventory)
		return
	}

	pageTitle := "DQIX | " + inventory.Title
	htmx.SetTitle(ctx, pageTitle)
	htmx.SetIcon(ctx, inventory.ImageSrc())
	params := pages.InventoryParams{
		Inventory: inventory,
		Getter:    a.data.GetQuickThing,
		LayoutParams: base.LayoutParams{
			PageTitle:  pageTitle,
			Page:       inventory.Classification,
			IsDarkMode: gin_utils.IsDarkMode(ctx),
			CSSVersion: a.cssVersion,
		},
	}

	switch htmx.GetHxSwapTarget(ctx) {
	case "page-content":
		if htmx.HasMatchingParentPath(ctx) {
			ctx.HTML(http.StatusOK, "", pages.InventoryContent(params))
		} else {
			htmx.Retarget(ctx, "#sidenav-page-wrapper")
			ctx.HTML(http.StatusOK, "", pages.InventoryContentWithSideNav(params))
		}
	case "sidenav-page-wrapper":
		ctx.HTML(http.StatusOK, "", pages.InventoryContentWithSideNav(params))
	default:
		ctx.HTML(http.StatusOK, "", pages.InventoryPage(params))
	}
}
