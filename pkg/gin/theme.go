package gin_utils

import (
	"dqix/web/templ/components/buttons"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	CookieTheme string = "theme"
	ThemeDark   string = "dark"
	ThemeLight  string = "light"
)

func IsDarkMode(ctx *gin.Context) bool {
	theme, err := ctx.Cookie(CookieTheme)
	return err == nil && theme == ThemeDark
}

func ToggleThemeHandler(ctx *gin.Context) {
	isDarkMode := IsDarkMode(ctx)

	value := ThemeDark
	if isDarkMode {
		value = ThemeLight
	}
	ctx.SetCookie(CookieTheme, value, 0, "", "", false, true)
	ctx.HTML(http.StatusOK, "", buttons.ThemeToggleButton(!isDarkMode))
}
