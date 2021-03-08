package models

import (
	"errors"
	entities "github.com/JIeeiroSst/togo/internal/storages"
	db "github.com/JIeeiroSst/togo/internal/storages/postgresql"
	"github.com/JIeeiroSst/togo/utils"
	"log"
)

func CheckAccount(Id string) (string,string,bool) {
	var accounts []entities.Users
	_ = db.GetConn().Find(&accounts)
	for _,account:=range accounts{
		if account.Id == Id{
			return account.Id, account.Password, true
		}
	}
	return  "", "", false
}

func CheckAccountExists(id string) bool {
	var account [] entities.Users
	_ = db.GetConn().Find(&account)
	for _,item:=range account{
		if item.Id == id {
			return false
		}
	}
	return true
}

func Login(id,password string) (string,string,error){
	Id, hashPassword,check := CheckAccount(id)
	if check == false {
		return "User does not exist","", errors.New("not exist")
	}
	if checkPass := utils.CheckPassowrd(password, hashPassword); checkPass != nil {
		return "password entered incorrectly","", errors.New("incorrectly")
	}
	token, _ := utils.GenerateToken(Id)
	return "logged in successfully",token ,nil
}

func SignUp(id,password string) string{
	check := CheckAccountExists(id)
	if check == false {
		return "user already exists"
	}
	hashPassword, err := utils.HashPassword(password)
	if err != nil {
		log.Println("error server", err)
	}
	account := entities.Users{
		Id:   id,
		Password:   hashPassword,
	}
	_ = db.GetConn().Create(&account)

	return id+":"+hashPassword
}

