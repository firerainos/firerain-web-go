package api

import (
	"github.com/gin-gonic/gin"
	"github.com/firerainos/firerain-web-go/userCenter"
)

func AddUser(ctx *gin.Context) {
	type Data struct {
		Username string `json:"username" form:"username" binding:"required"`
		Password string `json:"password" form:"username" binding:"required"`
		Email string `json:"email" form:"email" binding:"required"`
		Group []string `json:"group" form:"group" binding:"required"`
	}

	data := Data{}

	if err := ctx.Bind(&data);err != nil {
		ctx.JSON(400, gin.H{
			"code":104,
			"message":err.Error(),
		})
		return
	}

	if err := userCenter.AddUser(data.Username,data.Password,data.Email,data.Group); err != nil {
		ctx.JSON(200,gin.H{
			"code":104,
			"message":err.Error(),
		})
	}else {
		ctx.JSON(200,gin.H{
			"code":0,
		})
	}
}

func DeleteUser(ctx *gin.Context) {
	type Data struct {
		Id int `json:"id" form:"id" binding:"required"`
	}

	data := Data{}

	if err := ctx.Bind(&data);err != nil {
		ctx.JSON(400, gin.H{
			"code":104,
			"message":err.Error(),
		})
		return
	}

	user,err := userCenter.GetUserById(data.Id)
	if err != nil {
		ctx.JSON(200,gin.H{
			"code":104,
			"message":err.Error(),
		})
		return
	}

	if err =user.Delete(); err != nil {
		ctx.JSON(200,gin.H{
			"code":104,
			"message":err.Error(),
		})
	}else {
		ctx.JSON(200,gin.H{
			"code":0,
		})
	}
}

func GetUser(ctx *gin.Context) {
	users,err := userCenter.GetUser()
	if err != nil {
		ctx.JSON(200,gin.H{
			"code":104,
			"message":err.Error(),
		})
		return
	}

	ctx.JSON(200,gin.H{
		"code":0,
		"users": users,
	})
}


func AddGroup(ctx *gin.Context) {
	type Data struct {
		Group string `json:"group" form:"group" binding:"required"`
	}

	data := Data{}

	if err := ctx.Bind(&data);err != nil {
		ctx.JSON(400, gin.H{
			"code":104,
			"message":err.Error(),
		})
		return
	}

	if err := userCenter.AddGroup(data.Group); err != nil {
		ctx.JSON(200,gin.H{
			"code":104,
			"message":err.Error(),
		})
	}else {
		ctx.JSON(200,gin.H{
			"code":0,
		})
	}
}

func DeleteGroup(ctx *gin.Context) {

}

func GetGroup(ctx *gin.Context) {
	groups,err := userCenter.GetGroup()
	if err != nil {
		ctx.JSON(200,gin.H{
			"code":104,
			"message":err.Error(),
		})
		return
	}

	ctx.JSON(200,gin.H{
		"code":0,
		"groups": groups,
	})
}