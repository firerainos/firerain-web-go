package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"flag"
	"strconv"
	"github.com/firerainos/firerain-web-go/api"
	"github.com/firerainos/firerain-web-go/core"
	"os"
	"fmt"
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

	db,err:= core.GetSqlConn()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&api.List{})
	db.Close()

	router := gin.Default()

	apiRouter := router.Group("/api")

	apiRouter.POST("/list/add", api.AddList)
	apiRouter.GET("/list/list", api.GetList)
	apiRouter.DELETE("/list/delete", api.DelList,gin.BasicAuth(gin.Accounts{
		core.Conf.Username:core.Conf.Password,
	}),api.DelList)
	apiRouter.GET("/list/pass",gin.BasicAuth(gin.Accounts{
		core.Conf.Username:core.Conf.Password,
	}),api.PassList)

	router.Run(":" + strconv.Itoa(*port))
}