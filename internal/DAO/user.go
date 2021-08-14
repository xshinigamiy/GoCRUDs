package DAO

import (
	"GoCRUDs/internal/repository"
	"GoCRUDs/pkg/pojos"
	"errors"
	"math/rand"
)

var (
	UDAO *UserDAO
)

type UserDAO struct {
	Repository *repository.MySQLRepository
}

func init() {
	NewUserDAO()
}

func NewUserDAO() *UserDAO {
	UDAO = &UserDAO{
		Repository: repository.GetMySQLRepo(),
	}
	return UDAO
}

func GetUserDAO() *UserDAO {
	if UDAO != nil {
		return UDAO
	} else {
		NewUserDAO()
		return UDAO
	}
}

func (uDAO *UserDAO) CreateUser(user pojos.User) (error, int) {
	call := uDAO.Repository.CreateUser(user)
	if call != nil {
		return errors.New("error while inserting data into DB"), -1
	}
	return nil, rand.Int()
}
