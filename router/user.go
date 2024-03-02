package routes

import (
	"ginapp/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {

	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.UserLoginWithPassword)
	r.POST("/send-otp", handlers.SendOtp)
	r.POST("/verify-otp", handlers.VerifyOTP)

	r.Use(middleware.AuthMiddleware)

	return r
}
