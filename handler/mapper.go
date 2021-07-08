package handler

import "fasthttp_crud/model"

type IUserMapper interface {
	ToListUserDTO(list []model.User) (listDTO []model.UserDTO)
	ToUserDTO(user model.User) model.UserDTO
	ToUserOnCreateUser(dto model.CreateUserDTO) (user model.User)
	ToUserOnUpdateUser(dto model.UpdateUserDTO, oldUser model.User) (user model.User)
}
