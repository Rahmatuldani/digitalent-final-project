package middlewares

import (
	"github.com/Rahmatuldani/digitalent-project/data/response"
	"github.com/Rahmatuldani/digitalent-project/helper"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helper.VerifyToken(ctx)

		if err != nil {
			ctx.JSON(401, response.ErrorResponse{
				Message: "Bearer token error",
				Error: err.Error(),
			})
			return
		}
		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}
