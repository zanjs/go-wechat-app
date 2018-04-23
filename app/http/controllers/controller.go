package controllers

import (
	"github.com/georgehao/wechat/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Message struct {
	Errno  int         `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Debug  interface{} `json:"debug"`
	Data   interface{} `json:"data"`
}

type Controller struct {
}

func (c *Controller) Success(context *gin.Context, data interface{}) {
	r := success(data)
	context.IndentedJSON(http.StatusOK, r)
}

func (c *Controller) Fail(context *gin.Context, errno int) {
	errmsg := "data failure"
	if _, ok := config.ErrMsg[errno]; ok {
		errmsg = config.ErrMsg[errno]
	}
	r := fail(errmsg, errno)
	context.IndentedJSON(http.StatusOK, r)
}

func Unauthorized(msg string) Message {
	return fail(msg, config.ErrorAuth)
}

/*
	Success return success data
	@param data interface{} can be anyone
*/
func success(data interface{}) Message {
	return Message{Errno: 0, Errmsg: "success", Debug: "success", Data: data}
}

// Fail return failure data
func fail(errmsg string, errno int) Message {
	return Message{Errno: errno, Errmsg: errmsg, Debug: nil}
}
