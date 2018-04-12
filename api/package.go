package api

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"github.com/firerainos/firerain-web-go/core"
)

type Package struct {
	gorm.Model
	ItemID      uint `json:"itemID" form:"itemID" binding:"required"`
	Name        string `json:"name" form:"name" binding:"required" gorm:"type:varchar(100);unique"`
	Description string `json:"description" form:"description"`
}

func AddPackage(ctx *gin.Context) {
	pkg := Package{}

	if err := ctx.Bind(&pkg); err != nil {
		ctx.JSON(400, gin.H{
			"code":    106,
			"message": "name and itemId not null",
		})
		return
	}

	db, err := core.GetSqlConn()
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    106,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	if err := db.Create(&pkg).Error; err != nil {
		ctx.JSON(200, gin.H{
			"code":    106,
			"message": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"code": 0,
		})
	}
}

func EditPackage(ctx *gin.Context) {
	id := ctx.Param("id")
	pkg := Package{}
	if err := ctx.Bind(&pkg); err != nil {
		ctx.JSON(400, gin.H{
			"code":    106,
			"message": "name and itemId not null",
		})
		return
	}

	db, err := core.GetSqlConn()
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    106,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	if err := db.Model(&pkg).Where("id=?",id).Update(&pkg).Error;err!=nil{
		ctx.JSON(200, gin.H{
			"code": 106,
			"message":err.Error(),
		})
	}else {
		ctx.JSON(200, gin.H{
			"code": 0,
		})
	}
}

func DeletePackage(ctx *gin.Context) {
	id := ctx.Param("id")

	db, err := core.GetSqlConn()
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    106,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	if err := db.Unscoped().Delete(&Package{}, id).Error; err != nil {
		ctx.JSON(200, gin.H{
			"code":    106,
			"message": "package not found",
		})
	} else {
		ctx.JSON(200, gin.H{
			"code": 0,
		})
	}
}

func GetPackages(ctx *gin.Context) {
	db, err := core.GetSqlConn()
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    106,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	var packages []Package
	db.Find(&packages)

	ctx.JSON(200,gin.H{
		"code": 0,
		"packages":packages,
	})
}
