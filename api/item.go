package api

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"github.com/firerainos/firerain-web-go/core"
	"strconv"
)

type Item struct {
	gorm.Model
	Name string `json:"name" form:"name" gorm:"type:varchar(100);unique" binding:"required"`
	packages []Package
}

func AddItem(ctx *gin.Context) {
	item := Item{}

	if err := ctx.Bind(&item);err!=nil{
		ctx.JSON(400,gin.H{
			"code":107,
			"message":"name not null",
		})
		return
	}

	db,err := core.GetSqlConn()
	if err != nil {
		ctx.JSON(200,gin.H{
			"code":107,
			"message":err.Error(),
		})
		return
	}
	defer db.Close()

	if err :=db.Create(&item).Error;err!=nil{
		ctx.JSON(200,gin.H{
			"code":107,
			"message":err.Error(),
		})
	}else{
		ctx.JSON(200,gin.H{
			"code":0,
		})
	}

}

func DeleteItem(ctx *gin.Context) {
	id := ctx.Param("id")

	db,err := core.GetSqlConn()
	if err != nil {
		ctx.JSON(200,gin.H{
			"code":107,
			"message":err.Error(),
		})
		return
	}
	defer db.Close()

	if err := db.Delete(&Item{},id).Error;err!=nil{
		ctx.JSON(200,gin.H{
			"code":107,
			"message":err.Error(),
		})
	}else{
		ctx.JSON(200,gin.H{
			"code":0,
		})
	}
}

func EditItem(ctx *gin.Context) {
	id := ctx.Param("id")

	item := Item{}

	if err := ctx.Bind(&item);err!=nil{
		ctx.JSON(400,gin.H{
			"code":107,
			"message":"name not null",
		})
		return
	}

	tmp,err:=strconv.ParseUint(id,10,32)
	if err != nil {
		ctx.JSON(200,gin.H{
			"code":107,
			"message":"id must int",
		})
		return
	}
	item.ID=uint(tmp)

	db,err := core.GetSqlConn()
	if err != nil {
		ctx.JSON(200,gin.H{
			"code":107,
			"message":err.Error(),
		})
		return
	}
	defer db.Close()

	if err:= db.Model(&item).Update("name",item.Name).Error;err!= nil{
		ctx.JSON(200,gin.H{
			"code":107,
			"message":err.Error(),
		})
	}else {
		ctx.JSON(200, gin.H{
			"code": 0,
		})
	}
}

func GetItems(ctx *gin.Context) {
	var items []Item

	db,err := core.GetSqlConn()
	if err != nil {
		ctx.JSON(200,gin.H{
			"code":107,
			"message":err.Error(),
		})
		return
	}
	defer db.Close()

	db.Find(&items)

	ctx.JSON(200,gin.H{
		"code":0,
		"items":items,
	})
}


