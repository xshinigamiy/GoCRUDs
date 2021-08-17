package handler

import (
	"GoCRUDs/internal/DAO"
	"GoCRUDs/pkg/pojos"
	"GoCRUDs/pkg/request"
	"sync"
)

var (
	singleInstanceOs sync.Once
	CS               *UserHandler
)

func init() {
	NewUserHandler()
}

type UserHandler struct {
	userDAO *DAO.UserDAO
}

func NewUserHandler() *UserHandler {
	CS = &UserHandler{
		userDAO: DAO.GetUserDAO(),
	}
	return CS
}

func GetUserHandler() *UserHandler {
	singleInstanceOs.Do(func() {
		NewUserHandler()
	})
	return CS
}

func (uHandler *UserHandler) CreateUser(user request.User) (error, int) {
	err, Id := uHandler.userDAO.CreateUser(UserObjectMapper(user))
	return err, Id
}

func UserObjectMapper(user request.User) pojos.User {
	return pojos.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Age:       user.Age,
		// CreatedAt: user.CreatedAt,
		// UpdatedAt: user.UpdatedAt,
		UpdatedBy: user.UpdatedBy,
		CreatedBy: user.CreatedBy,
	}
}
