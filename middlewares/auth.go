package middlewares

import (
	"strconv"
	"time"

	"github.com/Rahmatuldani/digitalent-project/data/response"
	"github.com/Rahmatuldani/digitalent-project/helper"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication(ctx *gin.Context) {
	verifyToken, err := helper.VerifyToken(ctx)
	_ = verifyToken

	if err != nil {
		ctx.JSON(401, response.ErrorResponse{
			Message: "Unauthorized",
			Error: err.Error(),
		})
		ctx.Abort()
		return
	}
	
	token := verifyToken.(jwt.MapClaims)
	id := token["id"].(string)
	aid, _ := strconv.ParseUint(id, 10, 64)
	
	exp := token["exp"].(float64)
	if float64(time.Now().Unix()) > exp {
		ctx.JSON(500, response.ErrorResponse{
			Message: "Unauthorized",
			Error: "token expired",
		})
		ctx.Abort()
		return
	}

	ctx.Set("userId", aid)
	CheckUser(ctx)
}