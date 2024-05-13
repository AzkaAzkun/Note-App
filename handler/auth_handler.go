package handler

import (
	"net/http"

	"github.com/AzkaAzkun/Note-App/config"
	"github.com/AzkaAzkun/Note-App/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx) error {
	var userInput models.User
	if err := c.BodyParser(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	var user models.User
	if err := config.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "Username atau password salah"}
			return c.Status(http.StatusUnauthorized).JSON(response)
		default:
			response := map[string]string{"message": err.Error()}
			return c.Status(http.StatusInternalServerError).JSON(response)
		}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": "Username atau password salah"}
		return c.Status(http.StatusUnauthorized).JSON(response)
	}

	// expTime := time.Now().Add(time.Minute*1)
	// claims
	return nil
}
