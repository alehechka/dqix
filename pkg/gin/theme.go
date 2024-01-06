package gin_utils

import (
	"github.com/gin-gonic/gin"
)

const (
	AttributeDataMode string = "data-mode"
	CookieTheme       string = "theme"
	ThemeDark         string = "dark"
	ThemeLight        string = "light"
)

func IsDarkMode(ctx *gin.Context) bool {
	theme, err := ctx.Cookie(CookieTheme)
	return err == nil && theme == ThemeDark
}
