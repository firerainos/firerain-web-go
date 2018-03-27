package api

import (
	"github.com/gin-gonic/gin"
	"github.com/firerainos/firerain-web-go/userCenter"
)

func AddUser(ctx *gin.Context) {

}

func DeleteUser(ctx *gin.Context) {
	type Data struct {
		Id int `json:"id" form:"id" binding:"required"`
	}

	data := Data{}

	ctx.Bind(&data)

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