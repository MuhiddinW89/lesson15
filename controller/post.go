package controller

import (
	"errors"
	"lesson15/models"
)


func (c *Controller) CreatePost (req models.CreatePost) (string,error){
	_,err:=c.store.User().GetById(req.UserId)
	if err!=nil{
		return "", errors.New("bunday id li user yoq")
	}
	createpost := c.store.Post().CreatePost(req)
	return createpost,nil
}

func (c *Controller) GetByIdPost (req string) (models.Post, error){
	getbyidpost, erro := c.store.Post().GetByIdPost(req)
	return getbyidpost, erro
}

func (c *Controller) GetAllPost () []models.Post{
	getallpost := c.store.Post().GetAllPost()
	return getallpost
}

func (c *Controller) GetUserPosts (req string) ([]models.Post){
	getusersposts := c.store.Post().GetUserPosts(req)
	return getusersposts
}