package dto

import "TTMS/model"

type ManagerDto struct {
	Username string `json:"username"`
	Token    string `json:"token"`
	Email    string `json:"email"`
}

func BuildManager(manager *model.Manager, token string) *ManagerDto {
	return &ManagerDto{
		Email:    manager.Email,
		Username: manager.Username,
		Token:    token,
	}
}

func BuildManagerList(manager []model.Manager) []*ManagerDto {
	var managers []*ManagerDto
	for _, manager := range manager {
		userDto := &ManagerDto{
			Email:    manager.Email,
			Username: manager.Username,
		}
		managers = append(managers, userDto)
	}
	return managers
}
