package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPackage(ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{
		"deapplicant":"",
		"office":"",
	})
}