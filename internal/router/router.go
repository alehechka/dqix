package router

import (
	gin_utils "dqix/pkg/gin"
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
		ctx.HTML(http.StatusOK, "", pages.IndexPage(gin_utils.IsDarkMode(ctx)))
	})

	a.InventoryRoutes(engine)

	return engine
}

func (r *RouterHandler) Run(port int) error {
	if err := r.app.loadData(); err != nil {
		return err
	}

	addr := fmt.Sprintf(":%d", port)

	return r.router.Run(addr)
}
