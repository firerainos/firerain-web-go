package api

import (
	"github.com/gin-gonic/gin"
	"github.com/firerainos/firerain-web-go/search"
)

func Search(ctx *gin.Context) {
	arch := ctx.Query("arch")
	repo := ctx.Query("repo")
	query := ctx.Query("query")
	maintainer := ctx.Query("maintainer")
	flagged := ctx.Query("flagged")
	page := ctx.Query("page")

	pkgs,num,pages := search.GetPackages(page,arch,repo,query,maintainer,flagged)

	ctx.JSON(200,gin.H{
		"code":0,
		"packages":pkgs,
		"num":num,
		"pages":pages,
	})
}

func GetArch(ctx *gin.Context) {
	ctx.JSON(200,gin.H{
		"code": 0,
		"arch": search.GetArch(),
	})
}

func GetRepo(ctx *gin.Context) {
	ctx.JSON(200,gin.H{
		"code": 0,
		"repo": search.GetRepository(),
	})
}

func GetMaintainer(ctx *gin.Context) {
	ctx.JSON(200,gin.H{
		"code":       0,
		"maintainer": search.GetMaintainer(),
	})
}

func GetFlagged(ctx *gin.Context) {
	ctx.JSON(200,gin.H{
		"code":    0,
		"flagged": search.GetFlagged(),
	})
}