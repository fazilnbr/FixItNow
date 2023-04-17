package api

import (
	"fmt"
	"log"

	"github.com/fazilnbr/project-workey/pkg/api/handler"
	"github.com/fazilnbr/project-workey/pkg/api/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(authHandler handler.AuthHandler, adminHandler handler.AdminHandler, UserHandler handler.UserHandler, WorkerHandler handler.WorkerHandler, middleware middleware.Middleware) *ServerHTTP {
	engine := gin.New()
	authHandler.InitializeOAuthGoogle()

	// Use logger from Gin
	engine.Use(gin.Logger())

	// Swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Group users
	user := engine.Group("user")
	{
		// Phone number authentication
		user.POST("/sent-otp", authHandler.UserSendOTP)
		user.POST("/signup-and-login", authHandler.UserRegisterAndLogin)

		// Google authentication
		user.GET("/login-gl", authHandler.GoogleAuth)
		user.GET("/callback-gl", authHandler.CallBackFromGoogle)

		// Use Middileware
		user.Use(middleware.AthoriseJWT)

		user.POST("/add-profile", UserHandler.AddProfileAndUpdateMail)
	}

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	fmt.Print("\n\nddddddddd\n\n")
	err := sh.engine.Run(":9090")
	if err != nil {
		log.Fatalln(err)
	}
}
