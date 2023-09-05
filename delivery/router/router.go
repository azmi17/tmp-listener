package router

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kpango/glg"
)

func Start() error {
	gin.SetMode(gin.ReleaseMode)

	//Discard semua output yang dicatat oleh gin karena print out akan dicetak sesuai kebutuhan programmer
	gin.DefaultWriter = io.Discard

	//create router engine by default
	router := gin.Default()
	router.Use(gin.Recovery())

	RegisterHandler(router)

	listenerPort := os.Getenv("app.listener_port")
	_ = glg.Logf("[HTTP] Listening at : %s", listenerPort)
	return router.Run(":" + listenerPort)

}
