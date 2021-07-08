package model

type CreateUserDTO struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type UpdateUserDTO struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type UserDTO struct {
	Id       string `json:"id"`
	FullName string `json:"full_name"`
	Username string `json:"username"`
}
