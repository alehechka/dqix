package htmx

import (
	"net/url"
	"path"
	"strings"

	htmx_go "github.com/alehechka/htmx-go"
	htmx_ext "github.com/alehechka/htmx-go/ext"
	"github.com/gin-gonic/gin"
)

func IsHxRequest(ctx *gin.Context) bool {
	return ctx.GetHeader(htmx_go.HxRequest) == "true"
}

func IsHxBoosted(ctx *gin.Context) bool {
	return ctx.GetHeader(htmx_go.HxBoosted) == "true"
}

func GetHxSwapTarget(ctx *gin.Context) string {
	return ctx.GetHeader(htmx_go.HxTarget)
}

func TryGetHxSwapTarget(ctx *gin.Context) (target string, hasTarget bool) {
	target = ctx.GetHeader(htmx_go.HxTarget)

	return target, target != ""
}

func HasHxSwapTarget(ctx *gin.Context, expected string) bool {
	return ctx.GetHeader(htmx_go.HxTarget) == expected
}

func PushUrl(ctx *gin.Context, redirect string) {
	ctx.Header(htmx_go.HxPushUrl, redirect)
}

func ReplaceUrl(ctx *gin.Context, replace string) {
	ctx.Header(htmx_go.HxReplaceUrl, replace)
}

func Retarget(ctx *gin.Context, target string) {
	ctx.Header(htmx_go.HxRetarget, target)
}

func Reswap(ctx *gin.Context, swapStrategy string) {
	ctx.Header(htmx_go.HxReswap, swapStrategy)
}

func GetHxCurrentUrl(ctx *gin.Context) string {
	return ctx.GetHeader(htmx_go.HxCurrentURL)
}

func GetHxCurrentPath(ctx *gin.Context) string {
	currentUrl := ctx.GetHeader(htmx_go.HxCurrentURL)

	uri, _ := url.Parse(currentUrl)
	return uri.Path
}

func HasMatchingParentPath(ctx *gin.Context) bool {
	currentPath := GetHxCurrentPath(ctx)
	currentPathParts := strings.Split(currentPath, "/")

	requestPath := ctx.Request.URL.Path
	requestPathParts := strings.Split(requestPath, "/")
	requestPathParent := path.Join(requestPathParts[0 : len(requestPathParts)-1]...)

	// If the currentPath is a parent, then check that against the requestedPath's parent
	if len(currentPathParts) == len(requestPathParts)-1 {
		return strings.TrimPrefix(currentPath, "/") == requestPathParent
	}

	// Else, check the equality of parent paths
	currentPathParent := path.Join(currentPathParts[0 : len(currentPathParts)-1]...)
	return currentPathParent == requestPathParent
}

func IsRequestingParentPath(ctx *gin.Context) bool {
	currentPath := GetHxCurrentPath(ctx)
	currentPathParts := strings.Split(currentPath, "/")
	currentPathParent := path.Join(currentPathParts[0 : len(currentPathParts)-1]...)

	requestPath := ctx.Request.URL.Path
	requestPathParts := strings.Split(requestPath, "/")

	// If the requestPath is a parent, then check that against the currentPath's parent
	return len(currentPathParts)-1 == len(requestPathParts) && currentPathParent == strings.TrimPrefix(requestPath, "/")
}

func HasMatchingPath(ctx *gin.Context) bool {
	return GetHxCurrentPath(ctx) == ctx.Request.URL.Path
}

func SetTitle(ctx *gin.Context, title string) {
	ctx.Header(htmx_ext.HxTitle, title)
}

func SetIcon(ctx *gin.Context, iconHref string) {
	ctx.Header(htmx_ext.HxIcon, iconHref)
}
