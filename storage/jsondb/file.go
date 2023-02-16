package jsondb

import (
	"lesson15/storage"
	"os"
)

type Store struct{
	user *userRepo
	post *postRepo
}

func NewJson() (*Store, error) {
	userfile, err := os.Open("/home/muhiddin/Documents/goprogram/lesson15/users.json")
	if err != nil {
		return nil, err
	}

	postfile ,err := os.Open("/home/muhiddin/Documents/goprogram/lesson15/post.json")
	if err != nil {
		return nil, err
	}

	return &Store{
		post: NewPostRepo("/home/muhiddin/Documents/goprogram/lesson15/post.json", postfile),
		user: NewUserRepo("/home/muhiddin/Documents/goprogram/lesson15/users.json", userfile),
	}, nil
}

func (s *Store) CloseDB() {
	s.user.file.Close()
}

func (s *Store) User() storage.UserRepoInerface {
	return s.user
}

func (s *Store) Post() storage.PostRepoInerface {
	return s.post
}
