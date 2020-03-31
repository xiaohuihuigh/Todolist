package utils
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
var (
	HandlerSuccess      = ErrorString{0, "success"}
	ErrorParamsEncode   = ErrorString{-1001, "data json encode error"}
	ErrorApiHandler     = ErrorString{-1002, "服务端异常"}
	NotNewUserOrAlerted = ErrorString{-1, "不是新用户或者已经弹过"}
)

type ErrorString struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// 统一返回响应接口格式
func Response(ctx *gin.Context, errString ErrorString, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":        errString.Code,
		"currentTime": time.Now().Unix(),
		"msg":         errString.Message,
		"data":        data,
	})
	return
}

