package api

import (
	"gitlab.com/firerainos/firerain-web-go/database"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Item struct {
	gorm.Model
	Name     string `json:"name" form:"name" gorm:"type:varchar(100);unique" binding:"required"`
	Title    string `json:"title" form:"title" gorm:"type:varchar(100);unique" binding:"required"`
	Packages []Package
}

func AddItem(ctx *gin.Context) {
	item := Item{}

	if err := ctx.Bind(&item); err != nil {
		ctx.JSON(400, gin.H{
			"code":    107,
			"message": "name not null",
		})
		return
	}

	db := database.Instance()

	if err := db.Create(&item).Error; err != nil {
		ctx.JSON(200, gin.H{
			"code":    107,
			"message": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"code": 0,
		})
	}

}

func DeleteItem(ctx *gin.Context) {
	id := ctx.Param("id")

	db := database.Instance()

	if err := db.Unscoped().Delete(&Item{}, id).Error; err != nil {
		ctx.JSON(200, gin.H{
			"code":    107,
			"message": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"code": 0,
		})
	}
}

func EditItem(ctx *gin.Context) {
	id := ctx.Param("id")

	item := Item{}

	if err := ctx.Bind(&item); err != nil {
		ctx.JSON(400, gin.H{
			"code":    107,
			"message": "name not null",
		})
		return
	}

	tmp, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    107,
			"message": "id must int",
		})
		return
	}
	item.ID = uint(tmp)

	db := database.Instance()

	if err := db.Model(&item).Update(&item).Error; err != nil {
		ctx.JSON(200, gin.H{
			"code":    107,
			"message": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"code": 0,
		})
	}
}

func GetItems(ctx *gin.Context) {
	var items []Item

	db := database.Instance()

	db.Preload("Packages").Find(&items)

	ctx.JSON(200, gin.H{
		"code":  0,
		"items": items,
	})
}
