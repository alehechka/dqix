package router

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func (a *app) StaticFiles(engine *gin.Engine) {
	static := engine.Group("/static", func(ctx *gin.Context) {
		visibility := "private"
		maxAge := 86_400 // 1 day

		if strings.HasPrefix(ctx.Request.URL.Path, "/static/gallery") {
			maxAge = 31_536_000 // 1 year
			visibility = "public"
		}

		ctx.Header("Cache-Control", fmt.Sprintf("%s, max-age=%d, s-maxage=%d, no-transform", visibility, maxAge, maxAge))
	})

	static.Static("", "./web/static")
}
