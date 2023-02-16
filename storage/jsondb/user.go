package jsondb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"lesson15/models"
	"os"
	"strings"

	"github.com/google/uuid"
)

type userRepo struct{
	fileName string
	file *os.File
}

func NewUserRepo(a string, b *os.File) *userRepo{
	return &userRepo{
		fileName: a,
		file: b,
	}
}

//  create


func (u *userRepo) GetList () []models.Usr{
	user := []models.Usr{}

	data, err := ioutil.ReadFile("/home/muhiddin/Documents/goprogram/lesson15/users.json")
	if err != nil {
		fmt.Println(err)
	}

	erro := json.Unmarshal(data, &user)
	if erro != nil {
		fmt.Println(erro)
	}
	return user
}



func (u *userRepo)FindUsr (a string) []models.Usr{
	user := []models.Usr{}

	data, err := ioutil.ReadFile("/home/muhiddin/Documents/goprogram/lesson15/users.json")
	if err != nil {
		fmt.Println(err)
	}

	erro := json.Unmarshal(data, &user)
	if erro != nil {
		fmt.Println(erro)
	}

	retUsrs := []models.Usr{}
	for _,v:=range user{
		fullName := strings.ToLower( v.FirstName + " " + v.LastName)
		if strings.Contains(fullName, a) {
			retUsrs = append(retUsrs, v)
		}
	}
	return retUsrs
}


func (u *userRepo)GetById (a string) (models.Usr, error){
	user := []models.Usr{}

	data, err := ioutil.ReadFile("/home/muhiddin/Documents/goprogram/lesson15/users.json")
	if err != nil {
		fmt.Println(err)
	}

	erro := json.Unmarshal(data, &user)
	if erro != nil {
		fmt.Println(erro)
	}

	for _,v:=range user{
		if v.Id == a {
			return v,nil
		}
	}
	return models.Usr{},errors.New("User not found")
}


func (u *userRepo)DleteUsr (a string) ([]models.Usr, error){
	user := []models.Usr{}

	data, err := ioutil.ReadFile("/home/muhiddin/Documents/goprogram/lesson15/users.json")
	if err != nil {
		fmt.Println(err)
	}

	erro := json.Unmarshal(data, &user)
	if erro != nil {
		fmt.Println(erro)
	}

	retUsrs := []models.Usr{}
	for _,v:=range user{
		if v.Id != a {
			retUsrs = append(retUsrs, v)
		}
	}

	body, err := json.MarshalIndent(retUsrs, "", "   ")

	err = ioutil.WriteFile("/home/muhiddin/Documents/goprogram/lesson15/users.json", body, os.ModePerm)
	if err != nil {
		return []models.Usr{}, err
	}

	return retUsrs, nil
}


func (u *userRepo)UpdateUsr (a models.Usr) (models.Usr, error){
	user := []models.Usr{}

	data, err := ioutil.ReadFile("/home/muhiddin/Documents/goprogram/lesson15/users.json")
	if err != nil {
		fmt.Println(err)
	}

	erro := json.Unmarshal(data, &user)
	if erro != nil {
		fmt.Println(erro)
	}

	updatedUsr := models.Usr{}
	for i,v:=range user{
		if a.Id == v.Id {   
		user[i].FirstName = a.FirstName
		user[i].LastName = a.LastName
		user[i].Gender = a.Gender
		user[i].Card_number = a.Card_number
		user[i].Birthday = a.Birthday
		user[i].Profession = a.Profession
		updatedUsr = user[i]
		}
	}


	body, err := json.MarshalIndent(user, "", "   ")

	err = ioutil.WriteFile("/home/muhiddin/Documents/goprogram/lesson15/users.json", body, os.ModePerm)
	if err != nil {
		return models.Usr{}, err
	}

	return updatedUsr, nil
}


func (u *userRepo) CreateUsr (a models.Usr) ([]models.Usr, error){
	user := []models.Usr{}

	data, err := ioutil.ReadFile("/home/muhiddin/Documents/goprogram/lesson15/users.json")
	if err != nil {
		fmt.Println(err)
	}

	erro := json.Unmarshal(data, &user)
	if erro != nil {
		fmt.Println(erro)
	}

	d := uuid.New().String()
	createdUsr := []models.Usr{}

	b := models.Usr{
		Id: d,
		FirstName: a.FirstName,
		LastName: a.LastName,
		Gender: a.Gender,
		Card_number: a.Card_number,
		Birthday: a.Birthday,
		Profession: a.Profession,
	} 
	createdUsr = append(createdUsr, b)

	for i := 0; i < len(user); i++ {
		createdUsr = append(createdUsr, user[i])
	}


	body, err := json.MarshalIndent(createdUsr, "", "   ")

	err = ioutil.WriteFile("/home/muhiddin/Documents/goprogram/lesson15/users.json", body, os.ModePerm)
	if err != nil {
		return []models.Usr{}, err
	}

	return createdUsr, nil
}