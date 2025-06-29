package controllers

import (
	"fmt"
	"gin-fleamarket/services"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service}
}

func (c *UserController) ShowGreetings() {
	names := c.service.GetGreetingUserNames()
	for _, line := range names {
		fmt.Println(line)
	}
}
