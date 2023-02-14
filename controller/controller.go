package controller

import (
	"lesson15/models"
	"lesson15/storage"
)

type Controller struct {
	store storage.Store
}

func NewController(a storage.Store) *Controller {
	return &Controller{store : a}
}

func (c *Controller) GetAll() []models.Usr {
  getlist := c.store.User.GetList()
  return getlist
}
