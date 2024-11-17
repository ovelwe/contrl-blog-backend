package handlers

import (
	"contrl-blog/internal/db"
	"contrl-blog/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var input struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
		UserID  uint   `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
		UserID:  input.UserID,
	}

	if result := db.DB.Create(&post); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании поста"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}

func GetAllPosts(c *gin.Context) {
	var posts []models.Post
	if result := db.DB.Preload("User").Find(&posts); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении постов"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func GetPostByID(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if result := db.DB.Preload("User").Preload("Comments").First(&post, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пост не найден"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")

	var input struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var post models.Post
	if result := db.DB.First(&post, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пост не найден"})
		return
	}

	post.Title = input.Title
	post.Content = input.Content

	if result := db.DB.Save(&post); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка изменения поста"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	if result := db.DB.First(&post, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пост не найден"})
		return
	}

	if result := db.DB.Delete(&post); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить пост"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пост успешно удален"})
}

func CommentOnPost(c *gin.Context) {
	id := c.Param("id")

	var input struct {
		Content string `json:"content" binding:"required"`
		UserID  uint   `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var post models.Post
	if result := db.DB.First(&post, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пост не найден"})
		return
	}

	comment := models.Comment{
		Content: input.Content,
		UserID:  input.UserID,
		PostID:  post.ID,
	}

	if result := db.DB.Create(&comment); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при добавлении комментария"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Комментарий добавлен", "comment": comment})
}

func LikePost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	if result := db.DB.First(&post, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пост не найден"})
		return
	}

	post.Likes++

	if result := db.DB.Save(&post); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось лайкнуть пост"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пост лайкнут", "likes": post.Likes})
}
