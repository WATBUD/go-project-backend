package usecase

import (
	//_ "sample/internal/domain/user/service"
	"sample/internal/domain/user/controller/userRequest"
)

type CreateUserUseCase interface {
	CreateNewUser(data userRequest.CreateUserRequest) (int32, error)
}
