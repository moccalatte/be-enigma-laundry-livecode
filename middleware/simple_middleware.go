package middleware

import "github.com/gin-gonic/gin"

func SimpleMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.WriteString("Sebelum request\n")

		ctx.Next()

		ctx.Writer.WriteString("Sesudah request\n")
	}
}