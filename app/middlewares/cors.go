package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			origin = ctx.Request.Header.Get("Origin")
		)

		if origin != "" {
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			ctx.Header("Access-Control-Allow-Origin", "*")                                        // 允许访问所有域
			ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE") //服务器支持的所有跨域请求的方法

			//  header的类型
			ctx.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, "+
				"Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, "+
				"Accept-Language,DNT, X-CustomHeader, Keep-Alive, Operator-Agent, X-Requested-With, "+
				"If-Modified-Since, Cache-Control, Content-Type, Pragma")

			//  允许跨域设置
			ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, "+
				"Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,"+
				"Last-Modified,Pragma,FooBar") // 让浏览器可以解析
			ctx.Header("Access-Control-Max-Age", "600")            // 缓存请求信息 单位为秒
			ctx.Header("Access-Control-Allow-Credentials", "true") //  跨域请求是否需要带cookie信息 默认设置为true
			ctx.Set("content-type", "application/json")            // 设置返回格式是json
		}

		ctx.Next() //  处理请求
	}
}

// Cors 跨域处理
func CorsV2() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		origin := ctx.Request.Header.Get("Origin")
		ctx.Header("Access-Control-Allow-Origin", origin)
		ctx.Header("Access-Control-Max-Age", "600")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		ctx.Header("Access-Control-Allow-Headers", "x-token,token,DNT,X-Mx-ReqToken,Keep-Alive,Operator-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization")
		ctx.Next()
	}
}
