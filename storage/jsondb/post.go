package jsondb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"lesson15/models"
	"os"

	"github.com/google/uuid"
)

type postRepo struct{
	fileName string
	file *os.File
}

func NewPostRepo(a string, b *os.File) *postRepo{
	return &postRepo{
		fileName: a,
		file: b,
	}
}

func (p *postRepo) CreatePost (req models.CreatePost) (string) {
	
	post := []models.Post{}
	data, err := ioutil.ReadFile("/home/muhiddin/Documents/goprogram/lesson15/post.json")
	if err != nil {
		fmt.Println(err)
	}

	erro := json.Unmarshal(data, &post)
	if erro != nil {
		fmt.Println(erro)
	}

	id := uuid.New().String()

	pst := models.Post{
		Id: id,
		Title: req.Title,
		Description: req.Description,
		UserId: req.UserId,
	}

	postedList := []models.Post{}
	postedList = append(postedList, pst)
	for i := 0; i < len(post); i++ {
		postedList = append(postedList, post[i])
	}

	body, err := json.MarshalIndent(postedList, "", "   ")

	err = ioutil.WriteFile("/home/muhiddin/Documents/goprogram/lesson15/post.json", body, os.ModePerm)
	if err != nil {
		return "Failed to write post to file"
	}
	return pst.Id
}


func (p *postRepo) GetByIdPost (req string) (models.Post, error){
	post := []models.Post{}
	data, err := ioutil.ReadFile("/home/muhiddin/Documents/goprogram/lesson15/post.json")
	if err != nil {
		fmt.Println(err)
	}

	erro := json.Unmarshal(data, &post)
	if erro != nil {
		fmt.Println(erro)
	}

	for _, v := range post{
		if req == v.Id {
			return v, nil
		}
	}
	return models.Post{}, errors.New("Post not found")
}

func (p *postRepo) GetAllPost () []models.Post{
	post := []models.Post{}
	data, err := ioutil.ReadFile("/home/muhiddin/Documents/goprogram/lesson15/post.json")
	if err != nil {
		fmt.Println(err)
	}

	erro := json.Unmarshal(data, &post)
	if erro != nil {
		fmt.Println(erro)
	}

	return post
}

func (p *postRepo) GetUserPosts (req string) ([]models.Post) {
	post := []models.Post{}
	data, err := ioutil.ReadFile("/home/muhiddin/Documents/goprogram/lesson15/post.json")
	if err != nil {
		fmt.Println(err)
	}

	erro := json.Unmarshal(data, &post)
	if erro != nil {
		fmt.Println(erro)
	}

	slice := []models.Post{}
	for _, v := range post{
		if req == v.UserId {
			slice = append(slice, v)
		}
	}
	return slice
}


