package main

import (
	"flag"
	"github.com/firerainos/firerain-web-go/api"
	"github.com/firerainos/firerain-web-go/core"
	"os"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/firerainos/firerain-web-go/userCenter"
	"strconv"
	"log"
	"strings"
)

var port = flag.Int("p", 8080, "port")

func main() {
	flag.Parse()

	err := core.ParseConf("config.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("请配置config.json")
			os.Exit(0)
		}
		panic(err)
	}

	initDB()

	router := gin.Default()

	store := sessions.NewCookieStore([]byte("firerain"))
	router.Use(sessions.Sessions("firerain-session", store))

	apiRouter := router.Group("/api", checkPermissionMiddleware)

	apiRouter.POST("/login", api.Login)
	apiRouter.POST("/logout", api.Logout)
	apiRouter.POST("/signup", api.Signup)

	apiRouter.GET("/avatar/:username", api.GetAvatar)
	apiRouter.POST("/avatar", api.UploadAvatar)

	apiRouter.GET("/list", api.GetList)
	apiRouter.POST("/list", api.AddList)
	apiRouter.DELETE("/list/:id", api.DelList)
	apiRouter.PATCH("/list/:id", api.PassList)

	packageRouter := apiRouter.Group("/package")

	packageRouter.GET("", api.GetPackages)
	packageRouter.POST("", api.AddPackage)
	packageRouter.DELETE("/:id", api.DeletePackage)
	packageRouter.PUT("/:id", api.EditPackage)

	itemRouter := apiRouter.Group("/item")

	itemRouter.GET("", api.GetItems)
	itemRouter.POST("", api.AddItem)
	itemRouter.DELETE("/:id", api.DeleteItem)
	itemRouter.PUT("/:id", api.EditItem)

	uCenterRouter := apiRouter.Group("/userCenter")

	uCenterRouter.GET("/user", api.GetUser)
	uCenterRouter.POST("/user", api.AddUser)
	uCenterRouter.DELETE("/user/:id", api.DeleteUser)
	uCenterRouter.PATCH("/user/:id", api.EditUser)

	uCenterRouter.POST("/group", api.AddGroup)
	uCenterRouter.DELETE("/group/:name", api.DeleteGroup)
	uCenterRouter.GET("/group", api.GetGroup)

	router.Run(":" + strconv.Itoa(*port))
}

func checkPermissionMiddleware(ctx *gin.Context) {
	if strings.Contains(ctx.Request.RequestURI, "/api/list") &&
		ctx.Request.Method != "POST" {
			checkPermission(ctx, "admin")
	} else if strings.Contains(ctx.Request.RequestURI, "/api/package") ||
		strings.Contains(ctx.Request.RequestURI, "/api/item") {
		if ctx.Request.Method == "GET" {
			checkPermission(ctx, "insider")
		} else {
			checkPermission(ctx, "admin")
		}
	} else if strings.Contains(ctx.Request.RequestURI, "/api/userCenter") {
		if ctx.Request.Method == "PATCH" &&
			strings.Contains(ctx.Request.RequestURI, "/api/userCenter/user") {
			checkPermission(ctx, "insider")
		} else {
			checkPermission(ctx, "admin")
		}
	} else {
		ctx.Next()
	}
}

func checkPermission(ctx *gin.Context, group string) {
	session := sessions.Default(ctx)

	tmp := session.Get("username")
	if tmp == nil {
		ctx.JSON(200, gin.H{
			"code":    101,
			"message": "unauthorized",
		})

		ctx.Abort()
	}

	username := tmp.(string)

	user, err := userCenter.GetUserByName(username)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    101,
			"message": "user no found",
		})

		ctx.Abort()
	}

	if group != "admin" {
		if user.HasGroup("admin") {
			ctx.Next()
			return
		}
	}

	if !user.HasGroup(group) {
		ctx.JSON(200, gin.H{
			"code":    101,
			"message": "permission denied",
		})
		ctx.Abort()
	}

	ctx.Next()

}

func initDB() {
	db, err := core.GetSqlConn()
	if err != nil {
		log.Panic(err)
	}

	var createUser, createGroup bool

	createUser = !db.HasTable(&userCenter.Group{})
	createGroup = !db.HasTable(&userCenter.User{})

	db.AutoMigrate(&api.List{}, &userCenter.User{}, &userCenter.Group{}, &api.Package{}, &api.Item{})

	if createUser {
		userCenter.AddGroup("users", "user")
		userCenter.AddGroup("admin", "administration")
		userCenter.AddGroup("insider", "insider")
	}

	if createGroup {
		userCenter.AddUser("admin", "admin", "admin", "", []string{"users", "admin"})
	}

	db.Close()
}
