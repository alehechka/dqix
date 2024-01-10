package router

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func (a *app) StaticFiles(engine *gin.Engine) {
	static := engine.Group("/static", func(ctx *gin.Context) {
		if _, err := os.Stat("./web" + ctx.Request.URL.Path); err != nil {
			return
		}

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
