package middleware

import (
	entities "github.com/JIeeiroSst/togo/internal/storages"
	"github.com/JIeeiroSst/togo/internal/storages/postgresql"
	"github.com/JIeeiroSst/togo/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken:=c.Request.Header.Get("Authorization")
		if bearToken == ""{
			c.AbortWithStatusJSON(401,entities.RestResponse{Message: "Authentication failure: Token not provided"})
			return
		}
		strArr := strings.Split(bearToken, " ")
		message,err:=utils.ParseToken(strArr[1])
		if err!=nil{
			c.AbortWithStatusJSON(400,entities.RestResponse{Message: message})
			return
		}
		sessionId, _ := c.Cookie("current_subject")
		sub, err := postgresql.GlobalCache.Get(sessionId)
		if err != nil {
			c.AbortWithStatusJSON(401, entities.RestResponse{Message: "user hasn't logged in yet"})
			return
		}
		c.Set("current_subject", string(sub))
		c.Next()
	}
}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, existed := c.Get("current_subject")
		if !existed {
			c.AbortWithStatusJSON(401, entities.RestResponse{Message: "user hasn't logged in yet"})
			return
		}
		token:=c.Request.Header.Get("Authorization")
		strArr := strings.Split(token, " ")
		_ ,err :=utils.ParseToken(strArr[1])
		if err != nil {
			c.AbortWithStatusJSON(401, entities.RestResponse{Message: "the user does not have access to access "})
			return
		}
		c.Next()
	}
}
