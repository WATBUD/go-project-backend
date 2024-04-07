package service

import (
	"sample/internal/domain/user/repository/model"
	"sample/internal/util"
)

type UserService struct {
	structRepository i_UserRepository
}

type i_UserRepository interface {
	CreateNewUser(data *model.User) (int32, error)
}

func NewUserService(userRepo i_UserRepository) *UserService {
	util.PrintLogWithColor("Enter stockService", "#00ffff")

	callBack := UserService{
		structRepository: userRepo,
	}
	return &callBack
}

func (_target *UserService) CreateNewUser(importData *model.User) (int32, error) {
	// randomNum := rand.Intn(10000)
	// randomNumStr := strconv.Itoa(randomNum)

	// // 創建一個新的用戶結構
	// user := &model.User{
	// 	Account:   "example_account" + randomNumStr,
	// 	Username:  "example_username",
	// 	Password:  "example_password",
	// 	Email:     "example@example.com",
	// 	CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	// 	UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	// }

	// 調用 structRepository 的 CreateNewUser 方法創建新用戶
	createdUser, err := _target.structRepository.CreateNewUser(importData)
	if err != nil {
		return 0, err
	}

	return createdUser, nil
}
