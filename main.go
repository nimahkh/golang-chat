package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

var roomManager *Manager

func main() {
	roomManager = NewRoomManager()
	router := gin.Default()
	router.SetHTMLTemplate(Html)
	router.Use(static.Serve("/assets", static.LocalFile("./assets", true)))

	router.GET("/room/:roomid/:user", roomGET)
	router.POST("/room/:roomid/:user", roomPOST)
	router.DELETE("/room/:roomid/:user", roomDELETE)
	router.GET("/stream/:roomid/:user", stream)

	router.Run(":8082")
}

func stream(c *gin.Context) {
	roomid := c.Param("roomid")
	listener := roomManager.OpenListener(roomid)
	defer roomManager.CloseListener(roomid, listener)

	clientGone := c.Writer.CloseNotify()
	c.Stream(func(w io.Writer) bool {
		select {
		case <-clientGone:
			return false
		case message := <-listener:
			c.SSEvent("message", message)
			return true
		}
	})
}

func roomGET(c *gin.Context) {
	roomid := c.Param("roomid")
	userid := c.Param("user")
	isme := c.Param("user")
	c.HTML(http.StatusOK, "chat_room", gin.H{
		"roomid": roomid,
		"userid": userid,
		"isme": isme,
	})
}

func roomPOST(c *gin.Context) {
	roomid := c.Param("roomid")
	userid := c.PostForm("user")
	message := c.PostForm("message")
	roomManager.Submit(userid, roomid, message)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": message,
	})
}

func roomDELETE(c *gin.Context) {
	roomid := c.Param("roomid")
	roomManager.DeleteBroadcast(roomid)
}