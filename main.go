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

	apiRouter.GET("/installer/package/:de", api.GetPackage)

	apiRouter.POST("/login", api.Login)

	apiRouter.GET("/list", checkAdminMiddleware, api.GetList)
	apiRouter.POST("/list", api.AddList)
	apiRouter.DELETE("/list/:id", checkAdminMiddleware, api.DelList)
	apiRouter.PATCH("/list/:id", checkAdminMiddleware, api.PassList)

	packageRouter := apiRouter.Group("/package", checkAdminMiddleware)

	packageRouter.GET("/")
	packageRouter.POST("/")
	packageRouter.DELETE("/:package")
	packageRouter.PUT("/:package")

	itemRouter := apiRouter.Group("/item", checkAdminMiddleware)

	itemRouter.GET("/")
	itemRouter.POST("/")
	itemRouter.DELETE("/:item")
	itemRouter.PUT("/:item")

	uCenterRouter := apiRouter.Group("/userCenter", checkAdminMiddleware)

	uCenterRouter.POST("/user", api.AddUser)
	uCenterRouter.DELETE("/user/:id", api.DeleteUser)
	uCenterRouter.GET("/user", api.GetUser)

	uCenterRouter.POST("/group", api.AddGroup)
	uCenterRouter.DELETE("/group/:name", api.DeleteGroup)
	uCenterRouter.GET("/group", api.GetGroup)

	router.Run(":" + strconv.Itoa(*port))
}

func checkAdminMiddleware(ctx *gin.Context) {
	checkPermission(ctx,"admin")
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

	db.AutoMigrate(&api.List{})
	db.AutoMigrate(&userCenter.User{})
	db.AutoMigrate(&userCenter.Group{})

	if createUser {
		userCenter.AddGroup("users", "user")
		userCenter.AddGroup("admin", "administration")
		userCenter.AddGroup("insider", "insider")
	}

	if createGroup {
		userCenter.AddUser("admin", "admin", "", []string{"users", "admin"})
	}

	db.Close()
}
