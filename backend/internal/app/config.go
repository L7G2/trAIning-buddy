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

	router.GET("/products", productHandler.List)

	//User:
	userRepo := repositories.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)

	router.GET("/users", userHandler.List)
	router.POST("/users", userHandler.Create)
	router.GET("/users/:userID", userHandler.GetByID)

	//ProgressReport:
	progressRepo := repositories.NewProgressReportRepository(db)
	progressHandler := handlers.NewProgressReportHandler(progressRepo)
	router.Group("/users/:userID/progress-reports").
		POST("", progressHandler.Create).
		GET("", progressHandler.List)

	//DietPlan:
	dietPlanRepo := repositories.NewDietPlanRepository(db)
	dietPlanHandler := handlers.NewDietPlanHandler(dietPlanRepo)
	router.Group("/users/:userID/diet-plans").
		POST("", dietPlanHandler.Create).
		GET("", dietPlanHandler.List)
	router.GET("/diet-plans/:dietPlanID/summary", dietPlanHandler.Summary)

	//Meal:
	mealRepo := repositories.NewMealRepository(db)
	mealHandler := handlers.NewMealHandler(mealRepo)
	router.Group("/diet-plans/:dietPlanID/meals").
		POST("", mealHandler.Create).
		GET("", mealHandler.List)

	//MediaFile:
	mediaFileRepo := repositories.NewMediaRepository(db)
	mediaFileHandler := handlers.NewMediaHandler(mediaFileRepo)

	mediaGroup := router.Group("/media")
	{
		mediaGroup.POST("", mediaFileHandler.Create)
		mediaGroup.GET("", mediaFileHandler.List)
	}
	router.GET("/exercises/:exerciseID/media", mediaFileHandler.GetByExerciseID)

	// Exercise-media-map:
	exerciseMediaRepo := repositories.NewExerciseMediaRepository(db)
	exerciseMediaHandler := handlers.NewExerciseMediaHandler(exerciseMediaRepo)

	router.POST("/exercise-media", exerciseMediaHandler.CreateLink)
	router.GET("/exercise-media/:exerciseID", exerciseMediaHandler.GetMediaByExerciseID)

}
