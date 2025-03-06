package controllers

import (
	"go-echo/db"
	"go-echo/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateUser creates a new user
func CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.String(http.StatusBadRequest, "Invalid input")
	}

	// Insert user into the database
	if err := db.GetDB().Create(&user).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Failed to create user")
	}

	return c.JSON(http.StatusCreated, user)
}

// GetAllUsers fetches all users
func GetAllUsers(c echo.Context) error {
	var users []models.User
	if err := db.GetDB().Find(&users).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get users")
	}

	return c.JSON(http.StatusOK, users)
}

// GetUserByID fetches user by ID
func GetUserByID(c echo.Context) error {
	id := c.Param("id")
	var user models.User
	if err := db.GetDB().First(&user, id).Error; err != nil {
		return c.String(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, user)
}

// UpdateUser updates user data by ID
func UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	var user models.User
	if err := db.GetDB().First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "User not found",
		})
	}
	
	updatedUser := new(models.User)
	if err := c.Bind(updatedUser); err != nil {
		return c.String(http.StatusBadRequest, "Invalid input")
	}
	
	user.Name = updatedUser.Name
	user.Age = updatedUser.Age

	if err := db.GetDB().Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to update user",
		})
	}

	return c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a user by ID
func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	var user models.User
	result := db.GetDB().Where("id = ?", id).First(&user, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "User not found",
		})
	}

	if err := result.Delete(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to delete user",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "User deleted successfully",
	})
}
