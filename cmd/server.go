package main

import (
	"ctraderapi/config"
	"ctraderapi/messagehandler"
	"ctraderapi/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	router := initServer(cfg)
	router.Run(":" + "8070")

}

func initServer(cfg *config.Config) *gin.Engine {
	hub := middlewares.NewHub()
	go hub.Run()
	router := gin.Default()

	//[Websocket] Echo Endpoint ------
	router.GET("/ws", func(c *gin.Context) {
		messagehandler.ConnectToOpenApi(cfg.Server.Endpoint, cfg.Server.Port, hub, c.Writer, c.Request)
	})
	return router
}
