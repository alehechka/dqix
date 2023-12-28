package htmx

import (
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
