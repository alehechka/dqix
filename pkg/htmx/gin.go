package htmx

import (
	"net/url"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

func IsHxRequest(ctx *gin.Context) bool {
	return ctx.GetHeader(RequestHeaderRequest) == "true"
}

func IsHxBoosted(ctx *gin.Context) bool {
	return ctx.GetHeader(RequestHeaderBoosted) == "true"
}

func GetHxSwapTarget(ctx *gin.Context) string {
	return ctx.GetHeader(RequestHeaderTarget)
}

func TryGetHxSwapTarget(ctx *gin.Context) (target string, hasTarget bool) {
	target = ctx.GetHeader(RequestHeaderTarget)

	return target, target != ""
}

func HasHxSwapTarget(ctx *gin.Context, expected string) bool {
	return ctx.GetHeader(RequestHeaderTarget) == expected
}

func PushUrl(ctx *gin.Context, redirect string) {
	ctx.Header(ResponseHeaderPushUrl, redirect)
}

func ReplaceUrl(ctx *gin.Context, replace string) {
	ctx.Header(ResponseHeaderReplaceUrl, replace)
}

func Retarget(ctx *gin.Context, target string) {
	ctx.Header(ResponseHeaderRetarget, target)
}

func Reswap(ctx *gin.Context, swapStrategy string) {
	ctx.Header(ResponseHeaderReswap, swapStrategy)
}

func GetHxCurrentUrl(ctx *gin.Context) string {
	return ctx.GetHeader(RequestHeaderCurrentURL)
}

func GetHxCurrentPath(ctx *gin.Context) string {
	currentUrl := ctx.GetHeader(RequestHeaderCurrentURL)

	uri, _ := url.Parse(currentUrl)
	return uri.Path
}

func HasMatchingParentPath(ctx *gin.Context) bool {
	currentPath := GetHxCurrentPath(ctx)
	currentPathParts := strings.Split(currentPath, "/")

	requestPath := ctx.Request.URL.Path
	requestPathParts := strings.Split(requestPath, "/")
	requestPathParent := path.Join(requestPathParts[0 : len(requestPathParts)-1]...)

	// If the currentPath is the parent, then check that against the requestedPath's parent
	if len(currentPathParts) == len(requestPathParts)-1 {
		return strings.TrimPrefix(currentPath, "/") == requestPathParent
	}

	currentPathParent := path.Join(currentPathParts[0 : len(currentPathParts)-1]...)

	// If the requestPath is the parent, then check that against the currentPath's parent
	if len(currentPathParts)-1 == len(requestPathParts) {
		return currentPathParent == strings.TrimPrefix(requestPath, "/")
	}

	// Else, check the equality of parent paths
	return currentPathParent == requestPathParent
}

func HasMatchingPath(ctx *gin.Context) bool {
	return GetHxCurrentPath(ctx) == ctx.Request.URL.Path
}

func SetTitle(ctx *gin.Context, title string) {
	ctx.Header(ResponseHeaderTitle, title)
}

func SetIcon(ctx *gin.Context, iconHref string) {
	ctx.Header(ResponseHeaderIcon, iconHref)
}
