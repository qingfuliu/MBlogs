package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goProject/generate"
	"strings"
)

func JWTAUTHMiddleWare() func(*gin.Context) {
	return func(c *gin.Context) {
		jwtHeader := c.GetHeader("aToken")
		jwtContent := strings.SplitN(jwtHeader, ",", 2)
		if len(jwtContent) != 2 || jwtContent[0] != "Bearer" {
			ResponseErrorWithData(c, CodeInvaildParams, errors.New("token parse failed"))
			c.Abort()
			return
		}
		mClaims, err := generate.ParseToken(jwtContent[1])
		if err != nil {
			//if generate.RefreshToken()
			ResponseErrorWithData(c, CodeInvaildParams, err)
			c.Abort()
			return
		}
		c.Set("username", mClaims.Username)
		c.Next()
	}
}
