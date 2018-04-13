package api

import (
	"github.com/firerainos/firerain-web-go/userCenter"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddUser(ctx *gin.Context) {
	type Data struct {
		Nickname string   `json:"nickname" form:"nickname" binding:"required"`
		Username string   `json:"username" form:"username" binding:"required"`
		Password string   `json:"password" form:"username" binding:"required"`
		Email    string   `json:"email" form:"email" binding:"required"`
		Group    []string `json:"group" form:"group" binding:"required"`
	}

	data := Data{}

	if err := ctx.Bind(&data); err != nil {
		ctx.JSON(400, gin.H{
			"code":    104,
			"message": err.Error(),
		})
		return
	}

	if err := userCenter.AddUser(data.Nickname, data.Username, data.Password, data.Email, data.Group); err != nil {
		ctx.JSON(200, gin.H{
			"code":    104,
			"message": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"code": 0,
		})
	}
}

func DeleteUser(ctx *gin.Context) {
	type Data struct {
		Id int `json:"id" form:"id" binding:"required"`
	}

	data := Data{}

	if err := ctx.Bind(&data); err != nil {
		ctx.JSON(400, gin.H{
			"code":    104,
			"message": err.Error(),
		})
		return
	}

	user, err := userCenter.GetUserById(data.Id)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    104,
			"message": err.Error(),
		})
		return
	}

	if err = user.Delete(); err != nil {
		ctx.JSON(200, gin.H{
			"code":    104,
			"message": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"code": 0,
		})
	}
}

func GetUser(ctx *gin.Context) {
	users, err := userCenter.GetUser()
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    104,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code":  0,
		"users": users,
	})
}

func EditUser(ctx *gin.Context) {
	id := ctx.Param("id")

	i, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    104,
			"message": "id must int",
		})
		return
	}

	user, err := userCenter.GetUserById(i)

	session := sessions.Default(ctx)

	tmp := session.Get("username")
	if tmp == nil {
		ctx.JSON(200, gin.H{
			"code":    101,
			"message": "unauthorized",
		})

		return
	}

	if user.Username != tmp.(string) {
		ctx.JSON(200, gin.H{
			"code":    101,
			"message": "unauthorized",
		})

		return
	}

	type Data struct {
		Nickname string `json:"nickname" form:"nickname"`
		Password string `json:"password" form:"password"`
	}

	data := Data{}

	if err := ctx.Bind(&data); err != nil {
		ctx.JSON(400, gin.H{
			"code":    104,
			"message": err.Error(),
		})
		return
	}

	if err = user.Edit(data.Nickname, data.Password); err != nil {
		ctx.JSON(200, gin.H{
			"code":    104,
			"message": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"code": 0,
		})
	}
}

func AddGroup(ctx *gin.Context) {
	type Data struct {
		Group       string `json:"group" form:"group" binding:"required"`
		Description string `json:"description" form:"description" binding:"required"`
	}

	data := Data{}

	if err := ctx.Bind(&data); err != nil {
		ctx.JSON(400, gin.H{
			"code":    104,
			"message": err.Error(),
		})
		return
	}

	if err := userCenter.AddGroup(data.Group, data.Description); err != nil {
		ctx.JSON(200, gin.H{
			"code":    104,
			"message": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"code": 0,
		})
	}
}

func DeleteGroup(ctx *gin.Context) {
	type Data struct {
		Group string `json:"group" form:"group" binding:"required"`
	}

	data := Data{}

	if err := ctx.Bind(&data); err != nil {
		ctx.JSON(400, gin.H{
			"code":    104,
			"message": err.Error(),
		})
		return
	}

	if err := userCenter.DeleteGroup(data.Group); err != nil {
		ctx.JSON(200, gin.H{
			"code":    104,
			"message": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"code": 0,
		})
	}
}

func GetGroup(ctx *gin.Context) {
	groups, err := userCenter.GetGroup()
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    104,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code":   0,
		"groups": groups,
	})
}
