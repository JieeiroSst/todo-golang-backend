package models

import (
	"errors"
	entities "github.com/JIeeiroSst/togo/internal/storages"
	db "github.com/JIeeiroSst/togo/internal/storages/postgresql"
)

func CreateTask(task entities.Tasks) error {
	err := db.GetConn().Create(&task)
	if err != nil {
		errors.New("create failure")
	}
	return nil
}

func UpdateTask(id string,task entities.Tasks) error{
	err := db.GetConn().Model(entities.Tasks{}).Where("id = ?",id).Updates(&task)
	if err != nil {
		errors.New("update failure")
	}
	return nil
}

func DeleteTags(id string) error {
	err := db.GetConn().Delete(entities.Tasks{},"id = ?",id)
	if err != nil {
		errors.New("delete failure")
	}
	return nil
}

func TagsList() ([]entities.Tasks,error){
	var tags []entities.Tasks
	err := db.GetConn().Find(&tags).Error
	if err!=nil{
		return nil, errors.New("query failure")
	}

	return tags,nil
}

func TagById(id string) (entities.Tasks,error){
	var tag entities.Tasks
	err := db.GetConn().Where("id = ?",id).Find(&tag).Error
	if err!=nil{
		return entities.Tasks{}, errors.New("query failure")
	}
	return tag,nil
}