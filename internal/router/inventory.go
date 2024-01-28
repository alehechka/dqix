package router

import (
	"dqix/internal/types"
	"dqix/internal/types/params"
	gin_utils "dqix/pkg/gin"
	"dqix/pkg/htmx"
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
	router.GET("/:type/:category/:classification", a.InventoryClassificationWrapper(InventoryClassificationRenderer))

	router.GET("/:type/:category/:classification/:id", a.InventoryHandler)
}

func (a *app) InventoryClassificationWrapper(handler func(*gin.Context, params.InventoryClassification)) func(*gin.Context) {
	return func(ctx *gin.Context) {
		typeId := ctx.Param("type")
		category := ctx.Param("category")
		classification := ctx.Param("classification")
		inventories := a.data.inventoryMap.GetClassificationSlice(typeId, category, classification, ctx.Query("sort"))

		pageTitle := "DQIX | " + strings.Title(classification)
		params := params.InventoryClassification{
			Classification:  classification,
			Inventories:     inventories,
			Stats:           inventories.GetHasInventoryStats(),
			DisplayMode:     ctx.Query("display"),
			SortPathGetter:  types.PrepareSimpleSortPath(*ctx.Request.URL),
			SortOrderGetter: types.GetSortOrder(ctx.Request.URL),
			LayoutParams: params.Layout{
				PageTitle:  pageTitle,
				Page:       classification,
				IsDarkMode: gin_utils.IsDarkMode(ctx),
				CSSVersion: a.cssVersion,
			},
		}

		handler(ctx, params)
	}
}

func InventoryClassificationRenderer(ctx *gin.Context, params params.InventoryClassification) {
	htmx.SetTitle(ctx, params.LayoutParams.PageTitle)
	htmx.SetIcon(ctx, params.LayoutParams.GetIconPath())

	switch htmx.GetHxSwapTarget(ctx) {
	case "inventory-table":
		ctx.HTML(http.StatusOK, "", pages.InventoryClassificationTable(params))
	case "page-content":
		ctx.HTML(http.StatusOK, "", pages.InventoryClassificationContent(params))
	case "sidenav-page-wrapper":
		if htmx.HasMatchingPath(ctx) || htmx.IsRequestingParentPath(ctx) {
			htmx.Retarget(ctx, "#page-content")
			htmx.Reswap(ctx, "innerHTML")
			ctx.HTML(http.StatusOK, "", pages.InventoryClassificationContent(params))
		} else {
			ctx.HTML(http.StatusOK, "", pages.InventoryClassificationContentWithSideNav(params))
		}
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
	params := params.Inventory{
		Inventory: inventory,
		Getter:    a.data.GetQuickThing,
		LayoutParams: params.Layout{
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
