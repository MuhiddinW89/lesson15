package controller

import (
	"lesson15/storage"
)

type Controller struct {
	store storage.Storage
}

func NewController(a storage.Storage) *Controller {
	return &Controller{store : a}
}

