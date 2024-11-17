package routes

import (
	"contrl-blog/internal/handlers"
	"contrl-blog/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/register", handlers.RegisterUser)
		api.POST("/login", handlers.LoginUser)

		api.Use(middlewares.AuthMiddleware())

		api.GET("/posts", handlers.GetAllPosts)
		api.GET("/posts/:id", handlers.GetPostByID)
		api.POST("/posts", handlers.CreatePost)
		api.PUT("/posts/:id", handlers.UpdatePost)
		api.DELETE("/posts/:id", handlers.DeletePost)

		// Лайки и комменты
		api.POST("/posts/:id/like", handlers.LikePost)
		api.POST("/posts/:id/comment", handlers.CommentOnPost)
	}
}
