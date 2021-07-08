package mapper

import (
	"fasthttp_crud/handler"
	"fasthttp_crud/model"
	"fasthttp_crud/util"
	"strings"
)

type UserMapper struct {
}

func NewUserMapper() handler.IUserMapper {
	return &UserMapper{}
}

func (m *UserMapper) ToListUserDTO(list []model.User) (listDTO []model.UserDTO) {
	for _, row := range list {
		listDTO = append(listDTO, m.ToUserDTO(row))
	}
	return listDTO
}

func (m *UserMapper) ToUserDTO(user model.User) model.UserDTO {
	fullName := user.FirstName
	if user.LastName != "" {
		fullName += " " + user.LastName
	}
	return model.UserDTO{
		Id:       user.Id,
		FullName: fullName,
		Username: user.Username,
	}
}

func (m *UserMapper) ToUserOnCreateUser(dto model.CreateUserDTO) (user model.User) {
	user = model.User{
		Id:        strings.ToLower(util.CreateTraceID()),
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Username:  dto.Username,
		Password:  dto.Password,
		Secret:    strings.ToLower(util.CreateRandomHex(10)),
	}
	return user
}

func (m *UserMapper) ToUserOnUpdateUser(dto model.UpdateUserDTO, oldUser model.User) (user model.User) {
	user = model.User{
		Id:        oldUser.Id,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Username:  dto.Username,
		Password:  dto.Password,
		Secret:    oldUser.Secret,
	}
	return user
}
