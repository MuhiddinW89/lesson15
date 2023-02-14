package storage

import "os"

type Store struct{
	User *userRepo
}

func NewJson() (*Store, error) {
	userfile, err := os.Open("/home/muhiddin/Documents/goprogram/lesson15/users.json")
	if err != nil {
		return nil, err
	}

	return &Store{
		User: NewUserRepo("/home/muhiddin/Documents/goprogram/lesson15/users.json", userfile),
	}, nil
}


