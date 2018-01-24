package api

import (
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/jinzhu/gorm"
	"github.com/firerainos/firerain-web-go/core"
	"net/http"
	"fmt"
)

type List struct {
	gorm.Model
	Region       string `form:"region" json:"region" binding:"required"`
	Email        string `form:"email" json:"email" binding:"required"`
	Qq           string `form:"qq" json:"qq" binding:"required"`
	Introduction string `form:"introduction" json:"introduction" binding:"required"`
	Suggest      string `form:"suggest" json:"suggest" binding:"required"`
}

func GetList(context *gin.Context) {
	db, err := core.GetSqlConn()
	if err != nil {
		panic(err)
	}
	var lists []List
	db.Find(&lists)

	context.JSON(http.StatusOK, lists)
}

func AddList(context *gin.Context) {
	var list List
	err := context.Bind(&list)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":"failure",
			"error":err.Error(),
		})
		return
	}
	db, err := core.GetSqlConn()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":"failure",
			"error":err.Error(),
		})
		return
	}
	db.Create(&list)
	fmt.Println(list)

	context.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func DelList(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func PassList(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
