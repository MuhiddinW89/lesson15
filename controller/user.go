package controller

import "lesson15/models"



func (c *Controller) GetList() []models.Usr {
	getlist := c.store.User().GetList()
	return getlist
  }
  
  func (c *Controller) FindUsr(a string) []models.Usr {
	  findusr := c.store.User().FindUsr(a)
	  return findusr
  }
  
  func (c *Controller) GetById (a string) (models.Usr, error){
	  getbyid, err := c.store.User().GetById(a)
	  return getbyid, err
  }
  
  func (c *Controller) DleteUsr (a string) ([]models.Usr, error){
	  dleteusr, err := c.store.User().DleteUsr(a)
	  return dleteusr, err
  }
  
  func (c *Controller) UpdateUsr (a models.Usr) (models.Usr, error){
	  updateusr, err := c.store.User().UpdateUsr(a)
	  return updateusr, err
  }
  
  func (c *Controller) CreateUsr (a models.Usr) ([]models.Usr, error){
	  createusr, err := c.store.User().CreateUsr(a)
	  return createusr, err
  }