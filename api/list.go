package api

import (
	"fmt"
	"github.com/firerainos/firerain-web-go/core"
	"github.com/firerainos/firerain-web-go/database"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type List struct {
	gorm.Model
	Region       string `form:"region" json:"region" binding:"required"`
	Email        string `form:"email" json:"email" binding:"required"`
	Qq           string `form:"qq" json:"qq" binding:"required"`
	Introduction string `form:"introduction" json:"introduction" binding:"required"`
	Suggest      string `form:"suggest" json:"suggest" binding:"required"`
	State        string
}

func GetList(context *gin.Context) {
	db := database.Instance()
	var lists []List
	db.Find(&lists)

	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"list": lists,
	})
}

func AddList(context *gin.Context) {
	var list List
	err := context.Bind(&list)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    103,
			"message": err.Error(),
		})
		return
	}
	db := database.Instance()
	db.Create(&list)
	fmt.Println(list)

	context.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
}

func DelList(context *gin.Context) {
	id := context.Param("id")
	db := database.Instance()
	var list List
	db.First(&list, id)
	db.Delete(&list)
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
}

func PassList(context *gin.Context) {
	id := context.Param("id")
	db := database.Instance()
	var list List
	db.First(&list, id)
	err := core.SendMail(list.Email)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":    103,
			"message": err.Error(),
		})
		return
	}
	db.Model(&list).Update("state", "pass")
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
}
