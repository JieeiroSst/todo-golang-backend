package main

import (
	"fmt"
	"github.com/JIeeiroSst/togo/middleware"
	"github.com/JIeeiroSst/togo/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router *gin.Engine
)

func init(){
	router = gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	router.Use(cors.New(corsConfig))
	router.POST("/login",routers.Login)
	router.POST("/signup",routers.SingUp)
	resource := router.Group("/api")
	resource.Use(middleware.Authenticate())
	{
		resource.GET("/tasks",middleware.Authorize(),routers.TaskList)
		resource.GET("/tasks/:id",middleware.Authorize(),routers.TagById)
		resource.POST("/tasks",middleware.Authorize(),routers.CreateTag)
		resource.POST("/tasks/delete/:id",middleware.Authorize(),routers.DeleteTag)
		resource.POST("/tasks/update/:id",middleware.Authorize(),routers.UpdateTags)
	}

}

func main (){
	err := router.Run()
	if err != nil {
		log.Fatalln(fmt.Errorf("faild to start Gin application: %w", err))
	}
}