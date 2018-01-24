package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"flag"
	"strconv"
	"github.com/firerainos/firerain-web-go/api"
	"github.com/firerainos/firerain-web-go/core"
)

var port = flag.Int("p",8080,"port")

func main() {
	err:=core.ParseConf("config.json")
	if err != nil {
		panic(err)
	}

	router:= gin.Default()

	apiRouter:= router.Group("/api")

	apiRouter.POST("/list/add", api.AddList)
	apiRouter.GET("/list/list", api.GetList)
	apiRouter.DELETE("/list/delete", api.DelList)
	apiRouter.GET("/list/pass")

	router.Run(":"+strconv.Itoa(*port))
}
