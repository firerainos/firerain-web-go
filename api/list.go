package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/firerainos/firerain-web-go/core"
	"net/http"
	"fmt"
	"net/smtp"
)

type List struct {
	gorm.Model
	Region       string `form:"region" json:"region" binding:"required"`
	Email        string `form:"email" json:"email" binding:"required"`
	Qq           string `form:"qq" json:"qq" binding:"required"`
	Introduction string `form:"introduction" json:"introduction" binding:"required"`
	Suggest      string `form:"suggest" json:"suggest" binding:"required"`
	State		 string
}

func GetList(context *gin.Context) {
	db, err := core.GetSqlConn()
	if err != nil {
		panic(err)
	}
	var lists []List
	db.Find(&lists)

	context.JSON(http.StatusOK, gin.H{
		"code":0,
		"list":lists,
	})
}

func AddList(context *gin.Context) {
	var list List
	err := context.Bind(&list)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":103,
			"message":err.Error(),
		})
		return
	}
	db, err := core.GetSqlConn()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":103,
			"message":err.Error(),
		})
		return
	}
	db.Create(&list)
	fmt.Println(list)

	context.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
}

func DelList(context *gin.Context) {
	id := context.Query("id")
	db, err := core.GetSqlConn()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":103,
			"message":err.Error(),
		})
		db.Close()
		return
	}
	var list List
	db.First(&list,id)
	db.Delete(&list)
	db.Close()
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
}

func PassList(context *gin.Context) {
	id := context.Query("id")
	db, err := core.GetSqlConn()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":103,
			"message":err.Error(),
		})
		return
	}
	var list List
	db.First(&list,id)
	username := core.Conf.Smtp.Username
	auth := smtp.PlainAuth("", username, core.Conf.Smtp.Password, core.Conf.Smtp.Host)
	subject := "FireRainOS内测申请审核通过"
	body := "您已通过FireRainOS内测申请审核，请及时(过时将关闭进群审核)加入qq群:615676312 (入群请填写申请时用的邮箱)来进一步获取内部内测消息及问题建议反馈"

	msg := []byte("To: " + list.Email + "\nFrom: " + username + "\nSubject: " + subject + "\n\n" + body)
	err = smtp.SendMail(core.Conf.Smtp.Host+":25", auth, core.Conf.Smtp.Username, []string{list.Email}, msg)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":103,
			"message":err.Error(),
		})
		db.Close()
		return
	}
	db.Model(&list).Update("state", "pass")
	db.Close()
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
}
