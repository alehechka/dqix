package router

import (
	"github.com/gin-gonic/gin"
)

func (a *app) StaticFiles(engine *gin.Engine) {
	static := engine.Group("/static", func(ctx *gin.Context) {
		ctx.Header("Cache-Control", "private, max-age=86400, s-maxage=86400, no-transform")
	})

	static.Static("", "./web/static")
}
