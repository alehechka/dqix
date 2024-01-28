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
	router.GET("/:type/:category/:classification/:id", a.InventoryWrapper(InventoryRenderer))
}

func (a *app) InventoryClassificationWrapper(handler func(*gin.Context, params.InventoryClassification)) func(*gin.Context) {
	return func(ctx *gin.Context) {
		typeId := ctx.Param("type")
		category := ctx.Param("category")
		classification := ctx.Param("classification")
		inventories := a.data.inventoryMap.GetClassificationSlice(typeId, category, classification, ctx.Query("sort"))

		params := params.InventoryClassification{
			Classification:  classification,
			Inventories:     inventories,
			Stats:           inventories.GetHasInventoryStats(),
			DisplayMode:     ctx.Query("display"),
			SortPathGetter:  types.PrepareSimpleSortPath(*ctx.Request.URL),
			SortOrderGetter: types.GetSortOrder(ctx.Request.URL),
			LayoutParams: params.Layout{
				PageTitle:  strings.Title(classification),
				Page:       classification,
				IsDarkMode: gin_utils.IsDarkMode(ctx),
				CSSVersion: a.cssVersion,
			},
		}

		handler(ctx, params)
	}
}

func InventoryClassificationRenderer(ctx *gin.Context, params params.InventoryClassification) {
	htmx.SetTitle(ctx, params.LayoutParams.GetPageTitle())
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

func (a *app) InventoryWrapper(handler func(*gin.Context, params.Inventory)) func(*gin.Context) {
	return func(ctx *gin.Context) {
		typeId := ctx.Param("type")
		category := ctx.Param("category")
		classification := ctx.Param("classification")
		id := ctx.Param("id")
		inventory := a.data.inventoryMap.GetInventory(typeId, category, classification, id)

		params := params.Inventory{
			Inventory: inventory,
			Getter:    a.data.GetQuickThing,
			LayoutParams: params.Layout{
				PageTitle:  inventory.Title,
				Page:       inventory.Classification,
				IsDarkMode: gin_utils.IsDarkMode(ctx),
				CSSVersion: a.cssVersion,
				IconPath:   inventory.ImageSrc(),
			},
		}

		handler(ctx, params)
	}
}

func InventoryRenderer(ctx *gin.Context, params params.Inventory) {
	htmx.SetTitle(ctx, params.LayoutParams.GetPageTitle())
	htmx.SetIcon(ctx, params.LayoutParams.GetIconPath())

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
