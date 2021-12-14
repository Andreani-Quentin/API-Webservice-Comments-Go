package controllers

import (
	"API-Webservice/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateCommentInput struct {
	Text  string `json:"comment" binding:"required"`
}

type UpdateCommentInput struct {
	Text  string `json:"comment"`
}

// GET /comments
// Find all comments
func FindComments(c *gin.Context) {
	var comments []models.Comment
	models.DB.Find(&comments)

	c.JSON(http.StatusOK, gin.H{"data": comments})
}

// GET /comments/:id
// Find a comment
func FindComment(c *gin.Context) {
	// Get model if exist
	var comment models.Comment
	if err := models.DB.Where("id = ?", c.Param("id")).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comment})
}

// POST /comments
// Create new comment
func CreateComment(c *gin.Context) {
	// Validate input
	var input CreateCommentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create comment
	comment := models.Comment{Text: input.Text}
	models.DB.Create(&comment)

	c.JSON(http.StatusOK, gin.H{"data": comment})
}

// PATCH /comments/:id
// Update a comment
func UpdateComment(c *gin.Context) {
	// Get model if exist
	var comment models.Comment
	if err := models.DB.Where("id = ?", c.Param("id")).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateCommentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&comment).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": comment})
}

// DELETE /comments/:id
// Delete a comment
func DeleteComment(c *gin.Context) {
	// Get model if exist
	var comment models.Comment
	if err := models.DB.Where("id = ?", c.Param("id")).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&comment)

	c.JSON(http.StatusOK, gin.H{"data": true})
}