package gin_utils

import "github.com/gin-gonic/gin"

func IsDarkMode(ctx *gin.Context) bool {
	theme, err := ctx.Cookie("theme")
	return err == nil && theme == "dark"
}
