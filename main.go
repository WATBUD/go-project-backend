package main

import (

	//_ "sample/internal/domain/user/service"
	// "sample/internal/domain/user/controller"

	// "sample/internal/domain/user/database"
	// "sample/internal/domain/user/repository"
	// "sample/internal/domain/user/service"
	// "sample/internal/util"

	//"github.com/jmoiron/sqlx"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// fx.New(
	// 	fx.Provide(func() *gin.Engine {
	// 		gin.SetMode(gin.ReleaseMode)
	// 		return gin.Default()
	// 	}),
	// 	fx.Provide(func() (*gorm.DB, error) {
	// 		return database.NewGormDB()
	// 	}),
	// 	fx.Provide(func(_database *gorm.DB) *repository.StockRepositoryGorm {
	// 		return repository.NewRepositoryGorm(_database)
	// 	}),
	// 	fx.Provide(func(repository *repository.StockRepositoryGorm) *service.StocksService {
	// 		return service.NewStocksService(repository)
	// 	}),
	// 	fx.Invoke(controller.InitializeStockController),

	// 	fx.Provide(func(_database *gorm.DB) *repository.UserRepositoryGorm {
	// 		return repository.NewUserRepositoryGorm(_database)
	// 	}),
	// 	fx.Provide(func(repository *repository.UserRepositoryGorm) *service.UserService {
	// 		return service.NewUserService(repository)
	// 	}),
	// 	fx.Invoke(controller.InitializeUserController),

	// 	fx.Invoke(func(router *gin.Engine) {
	// 		go func() {
	// 			util.PrintLogWithColor("Enter router.Run", "#0000ff")
	// 			if err := router.Run(":8083"); err != nil {
	// 				panic(err)
	// 			}
	// 		}()
	// 	}),
	// ).Run()
}
