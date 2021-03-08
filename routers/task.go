package routers

import (
	entities "github.com/JIeeiroSst/togo/internal/storages"
	"github.com/JIeeiroSst/togo/models"
	"github.com/gin-gonic/gin"
	"time"
)

func TaskList(c *gin.Context){
	tags,err:=models.TagsList()
	if err != nil {
		c.JSON(200,entities.RestResponse{
			Code:    200,
			Message: "no data found",
			Data:    nil,
		})
		return
	}
	c.JSON(200,entities.RestResponse{
		Code:    200,
		Message: "found Data",
		Data:    tags,
	})
}

func TagById(c *gin.Context){
	id:=c.Param("id")
	tag,err:=models.TagById(id)
	if err!=nil {
		c.JSON(200,entities.RestResponse{
			Code:    200,
			Message: "no data found",
			Data:    nil,
		})
		return
	}
	c.JSON(200,entities.RestResponse{
		Code:    200,
		Message: "found Data",
		Data:    tag,
	})
}

func CreateTag(c *gin.Context){
	tag:=entities.Tasks{
		Id:        c.Query("id"),
		Content:   c.Query("content"),
		UserId:    c.Query("user_id"),
		CreatedDate: time.Now().String(),
	}
	err:=models.CreateTask(tag)
	if err!=nil{
		c.JSON(204,entities.RestResponse{
			Code:    204,
			Message: "create failure",
			Data:    nil,
		})
		return
	}
	c.JSON(201,entities.RestResponse{
		Code:    201,
		Message: "create successfully",
		Data:    tag,
	})
}

func DeleteTag(c *gin.Context){
	id:=c.Param("id")
	err:=models.DeleteTags(id)
	if err!=nil{
		c.JSON(204,entities.RestResponse{
			Code:    204,
			Message: "delete failure",
			Data:    nil,
		})
		return
	}
	c.JSON(200,entities.RestResponse{
		Code:    200,
		Message: "Delete successfully",
		Data:    id,
	})
}

func UpdateTags(c *gin.Context){
	id:=c.Param("id")
	tag:=entities.Tasks{
		Id:          id,
		Content:     c.Query("content"),
	}
	err:=models.UpdateTask(id,tag)
	if err!=nil{
		c.JSON(204,entities.RestResponse{
			Code:    204,
			Message: "update failure",
			Data:    nil,
		})
		return
	}
	c.JSON(200,entities.RestResponse{
		Code:    200,
		Message: "Update successfully",
		Data:    tag,
	})
}