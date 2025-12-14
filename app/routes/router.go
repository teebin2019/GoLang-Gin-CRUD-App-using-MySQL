package routes

import (
	"database/sql"

	"GoCRUDApplicationMySQL/app/controllers"
	"GoCRUDApplicationMySQL/app/repositories"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	userRepo := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(*userRepo)

	vegetableRepo := repositories.NewVegetableRepository(db)
	vegetableController := controllers.NewVegetableController(*vegetableRepo)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", userController.GetAllUsers)
		v1.POST("/users", userController.CreateUser)

		v1.GET("/vegetables", vegetableController.GetAllVegetables)
		v1.POST("/vegetables", vegetableController.CreateVegetable)
		// Add other routes for GetUser, UpdateUser, and DeleteUser
	}

	return r
}
