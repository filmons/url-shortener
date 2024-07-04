package controllers

import (
    "net/http"
    "url-shortener/config"
    "url-shortener/models"
    "github.com/gin-gonic/gin"
    "github.com/teris-io/shortid"
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
