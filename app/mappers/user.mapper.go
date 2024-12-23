package mappers

import (
	"github.com/IKHINtech/go-fiber-rest-boilerplate/app/dto"
	"github.com/IKHINtech/go-fiber-rest-boilerplate/app/models"
)

func UserModel_ToUserDTO(user *models.User) *dto.UserDTO {
	return &dto.UserDTO{
		Name:     user.Name,
		Phone:    user.Phone,
		Email:    user.Email,
		IsActive: user.IsActive,
	}
}

func UserDTO_ToUserModel(userDTO *dto.UserDTO) *models.User {
	return &models.User{
		Name: userDTO.Name,
	}
}

func UpdateUserDTO_ToUserModel(updateUserDTO *dto.UpdateUserDTO) *models.User {
	return &models.User{
		Name: updateUserDTO.Name,
	}
}

func UsersModel_ToUsersDTOs(users []*models.User) []*dto.UserDTO {
	dtos := make([]*dto.UserDTO, len(users))
	for i, user := range users {
		dtos[i] = &dto.UserDTO{
			Name: user.Name,
		}
	}
	return dtos
}

func CreateUserDTO_ToUserModel(createUserDTO *dto.CreateUserDTO) *models.User {
	return &models.User{
		Name: createUserDTO.Name,
	}
}
