package main

import (
	"deercoder-gin/routers"
	_ "deercoder-gin/util/db"
	"github.com/gin-gonic/gin"
	"deercoder-gin/conf"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := routers.SetRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":"+conf.GetConfigValue("http_port"))
}
