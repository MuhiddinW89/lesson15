package models

type Post struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	UserId string
}
type CreatePost struct {
	Title string `json:"title"`
	Description string `json:"description"`
	UserId string
}
	
type GetPostByID struct{
	Id string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	User  Usr
}