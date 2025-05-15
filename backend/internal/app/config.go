package app

import (
	"backend/internal/domain/repositories"
	"backend/internal/handlers"
	"backend/internal/middleware"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func SetupMiddleware(router *gin.Engine) {
	router.Use(middleware.CORSMiddleware())
}

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Backend is up and running"})
	})

	auth := handlers.NewAuthHandler(db)
	router.POST("/register", auth.Register)
	router.POST("/login", auth.Login)
	router.GET("/me", middleware.AuthMiddleware(), auth.Me)

	profileRepo := repositories.NewProfileRepository(db)
	profileHandler := handlers.NewProfileHandler(profileRepo)

	authGroup := router.Group("/profile")
	authGroup.Use(middleware.AuthMiddleware())
	{
		authGroup.GET("", profileHandler.GetProfile)
		authGroup.POST("", profileHandler.SaveProfile)
	}

	trainingRepo := repositories.NewTrainingPlanRepository(db)
	trainingHandler := handlers.NewTrainingPlanHandler(trainingRepo)
	plansGroup := router.Group("/plans")
	plansGroup.Use(middleware.AuthMiddleware())
	{
		plansGroup.POST("", trainingHandler.Create)
		plansGroup.GET("", trainingHandler.GetMyPlans)
	}

	workoutRepo := repositories.NewWorkoutRepository(db)
	workoutHandler := handlers.NewWorkoutHandler(workoutRepo)
	router.Group("/plans/:planID/workouts").
		Use(middleware.AuthMiddleware()).
		POST("", workoutHandler.Create).
		GET("", workoutHandler.List)

	exerciseRepo := repositories.NewExerciseRepository(db)
	exerciseHandler := handlers.NewExerciseHandler(exerciseRepo)
	router.Group("/workouts/:workoutID/exercises").
		Use(middleware.AuthMiddleware()).
		POST("", exerciseHandler.Create).
		GET("", exerciseHandler.List)

	calculator := handlers.NewCalculatorHandler(db)

	router.POST("/calculate", middleware.AuthMiddleware(), calculator.Calculate)

	productHandler := handlers.NewProductHandler(db)

	router.GET("/products", middleware.AuthMiddleware(), productHandler.List)

}
