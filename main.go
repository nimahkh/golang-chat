package main

import (
	"gin/chat/controller"
	m "gin/chat/manager"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var roomManager *m.Manager

func main() {
	roomManager = m.NewRoomManager()
	router := gin.Default()
	router.LoadHTMLGlob("template/*")
	router.Use(static.Serve("/assets", static.LocalFile("./assets", true)))

	router.GET("/room/:roomid/:user", func(context *gin.Context) {
		controller.RoomGET(context)
	})
	router.POST("/room/:roomid/:user", func(context *gin.Context) {
		controller.RoomPOST(context, roomManager)
	})
	router.DELETE("/room/:roomid/:user", func(context *gin.Context) {
		controller.RoomDELETE(context, roomManager)
	})
	router.GET("/stream/:roomid/:user", func(context *gin.Context) {
		controller.Stream(context, roomManager)
	})

	router.Run()
}
