package routers

import (
	"fmt"
	entities "github.com/JIeeiroSst/togo/internal/storages"
	"github.com/JIeeiroSst/togo/internal/storages/postgresql"
	"github.com/JIeeiroSst/togo/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
)

func Login(c *gin.Context) {
	id ,password := c.Query("id"), c.Query("password")
	for iter := postgresql.GlobalCache.Iterator(); iter.SetNext(); {
		info, err := iter.Value()
		if err != nil {
			continue
		}
		if string(info.Value()) == id {
			postgresql.GlobalCache.Delete(info.Key())
			log.Printf("forced %s to log out\n", id)
			break
		}
	}
	message,token,err:=models.Login(id,password)
	if err != nil  {
		c.JSON(200, entities.RestResponse{Code: 200,Message: message,Data: token})
		return
	}
	u, err := uuid.NewRandom()
	if err != nil {
		log.Println(fmt.Errorf("failed to generate UUID: %w", err))
	}
	sessionId := fmt.Sprintf("%s-%s", u.String(), id)
	err = postgresql.GlobalCache.Set(sessionId, []byte(id))
	if err != nil {
		c.JSON(200, entities.RestResponse{Code: 200,Message: "failed to store current subject in cache"})
		return
	}
	c.SetCookie("current_subject", sessionId, 30*60, "/api", "", false, true)
	c.JSON(200, entities.RestResponse{Code: 200, Message: message,Data: token})
}

func SingUp(c *gin.Context) {
	username ,password := c.Query("username"), c.Query("password")
	info:= models.SignUp(username,password)
	c.JSON(200,entities.RestResponse{Code:200,Message: "signup in successfully",Data:info})
}