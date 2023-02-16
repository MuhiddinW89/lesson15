package storage

import "lesson15/models"


type Storage interface{
	CloseDB()
	User() UserRepoInerface
	Post() PostRepoInerface
}
type UserRepoInerface interface{
	GetList () []models.Usr
	FindUsr (a string) []models.Usr
	GetById (a string) (models.Usr, error)
	DleteUsr (a string) ([]models.Usr, error)
	UpdateUsr (a models.Usr) (models.Usr, error)
	CreateUsr (a models.Usr) ([]models.Usr, error)
}

type PostRepoInerface interface{
	CreatePost (req models.CreatePost) (string)
	GetByIdPost (req string) (models.Post, error)
	GetAllPost () []models.Post
	GetUserPosts (req string) ([]models.Post)
}
