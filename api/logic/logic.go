package logic

import (
	"github.com/gin-gonic/gin"
	"goim/pkg/proto/logic"
	"net/http"
)

func sendMsg(ctx *gin.Context) {
	arg := new(pb_logic.SendMessageReq)
	if err := ctx.Bind(arg); err != nil {
		// todo 参数绑定错误
		return
	}
	srv.SendMessage(ctx, arg)
	ctx.JSON(http.StatusOK, gin.H{})
}
