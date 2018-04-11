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

	apiRouter := router.Group("/api")

	apiRouter.POST("/login", api.Login)
	apiRouter.POST("/signup", api.Signup)

	apiRouter.GET("/avatar/:username",api.GetAvatar)
	apiRouter.POST("/avatar",api.UploadAvatar)

	apiRouter.GET("/list", checkAdminMiddleware, api.GetList)
	apiRouter.POST("/list", api.AddList)
	apiRouter.DELETE("/list/:id", checkAdminMiddleware, api.DelList)
	apiRouter.PATCH("/list/:id", checkAdminMiddleware, api.PassList)

	packageRouter := apiRouter.Group("/package")

	packageRouter.GET("",checkPermissionMiddleware,api.GetPackages)
	packageRouter.POST("",checkAdminMiddleware, api.AddPackage)
	packageRouter.DELETE("/:id",checkAdminMiddleware,api.DeletePackage)
	packageRouter.PUT("/:id",checkAdminMiddleware,api.EditPackage)

	itemRouter := apiRouter.Group("/item")

	itemRouter.GET("",checkPermissionMiddleware,api.GetItems)
	itemRouter.POST("",checkAdminMiddleware,api.AddItem)
	itemRouter.DELETE("/:id",checkAdminMiddleware,api.DeleteItem)
	itemRouter.PUT("/:id",checkAdminMiddleware,api.EditItem)

	uCenterRouter := apiRouter.Group("/userCenter", checkAdminMiddleware)

	uCenterRouter.GET("/user", api.GetUser)
	uCenterRouter.POST("/user", api.AddUser)
	uCenterRouter.DELETE("/user/:id", api.DeleteUser)
	uCenterRouter.PATCH("/user/:id", api.EditUser)

	uCenterRouter.POST("/group", api.AddGroup)
	uCenterRouter.DELETE("/group/:name", api.DeleteGroup)
	uCenterRouter.GET("/group", api.GetGroup)

	router.Run(":" + strconv.Itoa(*port))
}

func checkAdminMiddleware(ctx *gin.Context) {
	if !strings.Contains(ctx.Request.RequestURI, "/api/userCenter/user") && ctx.Request.Method != "PATCH" {
		checkPermission(ctx,"admin")
	}
}

func checkPermissionMiddleware(ctx *gin.Context) {
	checkPermission(ctx,"insider")
}

func checkPermission(ctx *gin.Context,group string){
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

	if !user.HasGroup(group) {
		ctx.JSON(200, gin.H{
			"code":    101,
			"message": "permission denied",
		})
	}

	ctx.Next()

}

func initDB() {
	db, err := core.GetSqlConn()
	if err != nil {
		log.Panic(err)
	}

	var createUser,createGroup bool

	createUser = !db.HasTable(&userCenter.Group{})
	createGroup = !db.HasTable(&userCenter.User{})

	db.AutoMigrate(&api.List{},&userCenter.User{},&userCenter.Group{},&api.Package{},&api.Item{})

	if createUser {
		userCenter.AddGroup("users", "user")
		userCenter.AddGroup("admin", "administration")
		userCenter.AddGroup("insider", "insider")
	}

	if createGroup {
		userCenter.AddUser("admin","admin", "admin", "", []string{"users", "admin"})
	}

	db.Close()
}
