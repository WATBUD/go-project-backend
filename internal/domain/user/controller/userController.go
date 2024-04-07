package controller

import (
	"net/http"
	userRequest "sample/internal/domain/user/controller/userRequest"

	"sample/internal/domain/user/usecase"

	"sample/internal/util"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	_usecase usecase.CreateUserUseCase
}

func InitializeUserController(userUseCase usecase.CreateUserUseCase, gin_Instance *gin.Engine) *UserController {
	_target := &UserController{
		_usecase: userUseCase,
	}

	_target.setAssignRoutes(gin_Instance)
	return _target
}

func (_target *UserController) setAssignRoutes(gin_Instance *gin.Engine) {

	util.PrintLogWithColor("Enter setAssignRoutes")
	//userService := _target.UserService
	gin_Instance.POST("/CreateNewUser", func(_ginCTX *gin.Context) {
		var CreateUserRequest userRequest.CreateUserRequest
		if err := _ginCTX.ShouldBindJSON(&CreateUserRequest); err != nil {
			_ginCTX.String(http.StatusBadRequest, "Invalid request data")
			return
		}

		newUserID, err := _target._usecase.CreateNewUser(CreateUserRequest)

		if err != nil {
			_ginCTX.String(http.StatusInternalServerError, "Error creating new user")
			return
		}

		_ginCTX.String(http.StatusOK, "New user created with ID: %v", newUserID)
	})

	gin_Instance.GET("/UserController", func(_ginCTX *gin.Context) {
		_ginCTX.String(http.StatusOK, "Hello, World! ----UserController 20240401")
	})

}
