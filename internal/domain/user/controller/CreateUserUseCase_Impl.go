package controller

import (
	userRequest "sample/internal/domain/user/controller/userRequest"
	"sample/internal/domain/user/repository/model"
	"sample/internal/domain/user/service"
	"sample/internal/domain/user/usecase"
	"time"

	"github.com/jinzhu/copier"
)

type UserUseCase struct {
	UserService *service.UserService
}

func InitializeUseCase(userService *service.UserService) usecase.CreateUserUseCase {
	return &UserUseCase{}
}

func (u *UserUseCase) CreateNewUser(_struct userRequest.CreateUserRequest) (userID int32, err error) {
	// newUserData := &model.User{
	// 	Account:   userInput.Account,
	// 	Username:  userInput.Username,
	// 	Password:  userInput.Password,
	// 	Email:     userInput.Email,
	// 	CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	// 	UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	// }
	UserData := &model.User{}
	if err = copier.Copy(&UserData, &_struct); err != nil {
		return
	}
	UserData.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	UserData.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	if userID, err = u.UserService.CreateNewUser(UserData); err != nil {
		return
	}
	return
}
