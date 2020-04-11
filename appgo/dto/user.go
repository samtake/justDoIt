package dto

import (
	"samtake/justdoit/model"
)

//UserDto .
type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

//ToUserDto .
func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
