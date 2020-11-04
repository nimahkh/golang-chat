package controller

import (
	"github.com/gin-gonic/gin"
	m "github.com/nimahkh/golang-chat/manager"
	"io"
	"net/http"
)

func Stream(c *gin.Context, rm *m.Manager) {
	roomid := c.Param("roomid")
	listener := rm.OpenListener(roomid)

	defer rm.CloseListener(roomid, listener)

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

func RoomGET(c *gin.Context) {
	roomid := c.Param("roomid")
	userid := c.Param("user")
	isme := c.Param("user")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"roomid": roomid,
		"userid": userid,
		"isme":   isme,
	})
}

func RoomPOST(c *gin.Context, rm *m.Manager) {
	roomid := c.Param("roomid")
	userid := c.PostForm("user")
	message := c.PostForm("message")
	rm.Submit(userid, roomid, message)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": message,
	})
}

func RoomDELETE(c *gin.Context, rm *m.Manager) {
	roomid := c.Param("roomid")
	rm.DeleteBroadcast(roomid)
}
