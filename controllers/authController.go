package controllers

import (
    "log"
    "net/http"
    "time"
    "url-shortener/config"
    "url-shortener/models"
    "github.com/dgrijalva/jwt-go"
    "golang.org/x/crypto/bcrypt"
    "github.com/gin-gonic/gin"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
    UserID uint `json:"user_id"`
    jwt.StandardClaims
}

// Register godoc
// @Summary Register new user
// @Description Register a new user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "User registration data"
// @Success 200 {object} map[string]interface{} "message: User registered successfully"
// @Failure 400 {object} map[string]interface{} "error: error message"
// @Failure 500 {object} map[string]interface{} "error: error message"
// @Router /register [post]
func Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
        return
    }

    user.Password = string(hashedPassword)
    if err := config.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

type loginInput struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

// Login godoc
// @Summary User login
// @Description Logs in a user and returns a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body loginInput true "Login Credentials"
// @Success 200 {object} map[string]interface{} "token: JWT token"
// @Failure 400 {object} map[string]interface{} "error: error message"
// @Failure 401 {object} map[string]interface{} "error: Invalid email or password"
// @Failure 500 {object} map[string]interface{} "error: error message"
// @Router /login [post]
func Login(c *gin.Context) {
    var input loginInput
    if err := c.ShouldBindJSON(&input); err != nil {
        log.Println("Error binding JSON:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    log.Println("Login Attempt - Email:", input.Email, "Password:", input.Password)

    var user models.User
    if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        log.Println("Invalid email or password:", err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        log.Println("Invalid email or password:", err)
        log.Println("Error Details:", err)
        log.Printf("Stored: %s, Provided: %s", user.Password, input.Password)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    log.Println("Password Match Successful")

    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        UserID: user.ID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        log.Println("Error generating token:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
        return
    }

    log.Println("User logged in successfully:", user.Email)
    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
