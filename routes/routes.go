package routes

import (
	// "net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/lingesh-sk/Employee_dashboard_application/controller"
	"github.com/lingesh-sk/Employee_dashboard_application/dal"
	"github.com/lingesh-sk/Employee_dashboard_application/service"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB, redisClient *redis.Client, recentDelClient *redis.Client) {
	todoDAL := dal.NewTodoDAL(db, redisClient, recentDelClient, "recent_del:", time.Minute*5)

	todoService := service.NewTodoService(todoDAL)

	todoController := controller.NewTodoController(todoService)

	router.GET("/login", todoController.Login)
	router.GET("/callback", todoController.Callback)
	router.GET("/user", todoController.GetUserHandler)

	router.POST("/register", todoController.RegisterUser)
	router.GET("/getregister", todoController.GetRegisterUser)
	router.POST("/login", todoController.VerifyLogin)

	router.POST("/todos", todoController.CreateTodo)
	router.GET("/todos/:id", todoController.GetTodo)
	router.GET("/todos", todoController.GetTodos)
	router.PUT("/todos/:id", todoController.UpdateTodo)
	router.DELETE("/todos/:id", todoController.DeleteTodo)

	router.GET("/recent_del", todoController.GetRecentDeletedTodos)
	router.POST("/clear", todoController.ClearData)
}
