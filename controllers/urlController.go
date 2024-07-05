package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
	"url-shortener/config"
	"url-shortener/models"
)

// CreateShortURL godoc
// @Summary Create a short URL
// @Description Creates a new short URL for the authenticated user
// @Tags URL Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param url body models.URL true "URL Data"
// @Success 200 {object} models.URL "Successfully created short URL"
// @Failure 400 {object} map[string]interface{} "error: error message"
// @Failure 401 {object} map[string]interface{} "error: Unauthorized"
// @Failure 500 {object} map[string]interface{} "error: error message"
// @Router /shorten [post]
func CreateShortURL(c *gin.Context) {
	var input models.URL
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Générer une URL courte unique
	shortURL, err := shortid.Generate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate short URL"})
		return
	}
	input.ShortURL = shortURL
	input.UserID = userID.(uint)

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func RedirectURL(c *gin.Context) {
	var url models.URL
	shortURL := c.Param("short_url")
	if err := config.DB.Where("short_url = ?", shortURL).First(&url).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	// Mettre à jour LastAccessedAt et incrémenter Clicks
	now := time.Now()
	if err := config.DB.Model(&url).Updates(map[string]interface{}{
		"last_accessed_at": &now,
		"clicks":           gorm.Expr("clicks + ?", 1),
	}).Error; err != nil {
		// Log l'erreur mais continuer la redirection
		log.Printf("Error updating URL stats: %v", err)
	}

	c.Redirect(http.StatusMovedPermanently, url.LongURL)
}

func GetUserURLs(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var urls []models.URL
	if err := config.DB.Where("user_id = ?", userID).Find(&urls).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving URLs"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": urls})
}

// DeleteUrlById godoc
// @Summary Delete a URL by ID
// @Description Deletes a URL by ID
// @Tags URL Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "URL ID"
// @Success 200 {object} map[string]interface{} "Successfully deleted URL"
// @Failure 400 {object} map[string]interface{} "error: error message"
// @Failure 401 {object} map[string]interface{} "error: Unauthorized"
// @Failure 404 {object} map[string]interface{} "error: URL not found"
// @Failure 500 {object} map[string]interface{} "error: error message"
// @Router /del/{id} [delete]
func DeleteURL(c *gin.Context) {
	var url models.URL
	urlID := c.Param("id")
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := config.DB.Where("id = ? AND user_id = ?", urlID, userID).First(&url).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	if err := config.DB.Delete(&url).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "URL deleted"})
}

// GetURLStats godoc
// @Summary Get URL statistics
// @Description Get statistics for a URL by ID
// @Tags URL Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "URL ID"
// @Success 200 {object} models.URL "URL statistics"
// @Failure 400 {object} map[string]interface{} "error: error message"
// @Failure 401 {object} map[string]interface{} "error: Unauthorized"
// @Failure 404 {object} map[string]interface{} "error: URL not found"
// @Failure 500 {object} map[string]interface{} "error: error message"
// @Router /stats/{id} [get]
func GetURLStats(c *gin.Context) {
	var url models.URL
	urlID := c.Param("id")
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := config.DB.Where("id = ? AND user_id = ?", urlID, userID).First(&url).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	lastAccessed := "Never"
	if url.LastAccessedAt != nil {
		lastAccessed = url.LastAccessedAt.Format(time.RFC3339)
	}

	stats := gin.H{
		"id":               url.ID,
		"long_url":         url.LongURL,
		"short_url":        url.ShortURL,
		"clicks":           url.Clicks,
		"created_at":       url.CreatedAt.Format(time.RFC3339),
		"last_accessed_at": lastAccessed,
	}

	c.JSON(http.StatusOK, gin.H{"data": stats})
}
