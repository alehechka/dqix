package router

import (
	gin_utils "dqix/pkg/gin"
	"dqix/pkg/htmx"
	"dqix/web/templ/pages"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *app) InventoryRoutes(engine *gin.Engine) {
	router := engine.Group("/inventory")

	router.GET("/")
	router.GET("/:type")
	router.GET("/:type/:category")
	router.GET("/:type/:category/:classification", func(ctx *gin.Context) {
		typeId := ctx.Param("type")
		category := ctx.Param("category")
		classification := ctx.Param("classification")
		inventories := a.data.inventoryMap.GetClassificationSlice(typeId, category, classification)
		hasStats := inventories.GetHasInventoryStats()
		displayMode := ctx.Query("display")

		// TODO add some utility function that checks if the request Accept's JSON (and/or weighted with others)
		if ctx.GetHeader("Accept") == "application/json" {
			ctx.JSON(http.StatusOK, inventories)
			return
		}

		switch htmx.GetHxSwapTarget(ctx) {
		case "page-content":
			ctx.HTML(http.StatusOK, "", pages.InventoryClassificationContent(classification, inventories, hasStats, displayMode))
		case "sidenav-page-wrapper":
			ctx.HTML(http.StatusOK, "", pages.InventoryClassificationContentWithSideNav(classification, inventories, hasStats, displayMode))
		default:
			ctx.HTML(http.StatusOK, "", pages.InventoryClassificationPage(classification, inventories, hasStats, displayMode, gin_utils.IsDarkMode(ctx)))
		}
	})

	router.GET("/:type/:category/:classification/:id", func(ctx *gin.Context) {
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

		switch htmx.GetHxSwapTarget(ctx) {
		case "page-content":
			if htmx.HasMatchingParentPath(ctx) {
				ctx.HTML(http.StatusOK, "", pages.InventoryContent(inventory, a.data.GetQuickThing))
			} else {
				htmx.Retarget(ctx, "#sidenav-page-wrapper")
				ctx.HTML(http.StatusOK, "", pages.InventoryContentWithSideNav(inventory, a.data.GetQuickThing))
			}
		case "sidenav-page-wrapper":
			ctx.HTML(http.StatusOK, "", pages.InventoryContentWithSideNav(inventory, a.data.GetQuickThing))
		default:
			ctx.HTML(http.StatusOK, "", pages.InventoryPage(inventory, a.data.GetQuickThing, gin_utils.IsDarkMode(ctx)))
		}
	})
}
