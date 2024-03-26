package dto

import "TTMS/model"

type CustomerDto struct {
	Username string `json:"username"`
	Token    string `json:"token"`
	Email    string `json:"email"`
}

func BuildUser(user *model.Customer, token string) *CustomerDto {
	return &CustomerDto{
		Email:    user.Email,
		Username: user.Username,
		Token:    token,
	}
}

func BuildUserList(user []model.Customer) []*CustomerDto {
	var users []*CustomerDto
	for _, user := range user {
		userDto := &CustomerDto{
			Email:    user.Email,
			Username: user.Username,
		}
		users = append(users, userDto)
	}
	return users
}
