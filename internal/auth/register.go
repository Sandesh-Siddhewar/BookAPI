package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUser(reg *gin.Context, db *gorm.DB) {
	var user Users
	if err := reg.ShouldBindJSON(&user); err != nil {
		reg.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		reg.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = hashedPassword
	if err := db.Create(&user).Error; err != nil {
		reg.JSON(500, gin.H{"error": "Failed to register user"})
		return
	}

	reg.JSON(201, gin.H{"message": "User registered successfully"})
}

func LoginUser(log *gin.Context, db *gorm.DB) {
	var user Users
	if err := log.ShouldBindJSON(&user); err != nil {
		log.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	var dbUser Users
	if err := db.Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
		log.JSON(401, gin.H{"error": "Invalid username or password"})
		return
	}
	if !CheckPasswordHash(user.Password, dbUser.Password) {
		log.JSON(401, gin.H{"error": "Invalid username or password"})
		return
	}
	token, err := GenerateToken(dbUser)
	if err != nil {
		log.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}
	log.JSON(200, gin.H{"token": token})
}
