package handler

import (
	"fasthttp_crud/model"
)

type IUserRepo interface {
	GetAll(traceId string, limit, offset int) (list []model.User, totalData int, err error)
	GetUserById(traceId, userId string) (user model.User, err error)
	GetUserByUsername(traceId, username string) (user model.User, err error)
	AddNewUser(traceId string, user model.User) (inserted bool, err error)
	UpdateUserById(traceId string, user model.User) (err error)
	DeleteUserById(traceId, userId string) (err error)
}
