package router

import (
	"dqix/pkg/htmx"
	"dqix/web/templ/pages"
	"fmt"
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type RouterHandler struct {
	app    *app
	router *gin.Engine
}

func NewRouter(options ...RouterOption) *RouterHandler {
	a := new(app)

	for _, option := range options {
		option.apply(a)
	}

	return &RouterHandler{
		app:    a,
		router: a.SetupRouter(),
	}

}

func (a *app) SetupRouter() *gin.Engine {
	engine := gin.Default()
	engine.HTMLRender = &TemplRender{}
	engine.Use(gzip.Gzip(gzip.DefaultCompression))
	engine.SetTrustedProxies(nil)

	a.StaticFiles(engine)

	engine.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "", pages.IndexPage())
	})

	engine.GET("/inventory")
	engine.GET("/inventory/:type")
	engine.GET("/inventory/:type/:category")
	engine.GET("/inventory/:type/:category/:classification", func(ctx *gin.Context) {
		typeId := ctx.Param("type")
		category := ctx.Param("category")
		classification := ctx.Param("classification")

		switch htmx.GetHxSwapTarget(ctx) {
		case "page-content":
			ctx.HTML(http.StatusOK, "", pages.InventoryClassificationContent(classification, a.data.inventoryMap.GetClassificationSlice(typeId, category, classification)))
		case "sidenav-page-wrapper":
			ctx.HTML(http.StatusOK, "", pages.InventoryClassificationContentWithSideNav(classification, a.data.inventoryMap.GetClassificationSlice(typeId, category, classification)))
		default:
			ctx.HTML(http.StatusOK, "", pages.InventoryClassificationPage(classification, a.data.inventoryMap.GetClassificationSlice(typeId, category, classification)))
		}

	})
	engine.GET("/inventory/:type/:category/:classification/:id")

	return engine
}

func (r *RouterHandler) Run(port int) error {
	if err := r.app.loadData(); err != nil {
		return err
	}

	fmt.Println(len(r.app.data.inventoryMap))

	addr := fmt.Sprintf(":%d", port)

	return r.router.Run(addr)
}
