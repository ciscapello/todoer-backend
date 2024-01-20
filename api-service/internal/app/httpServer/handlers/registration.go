package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ciscapello/api-service/internal/app/utils"

	"github.com/gin-gonic/gin"
)

type registrationRequestBody struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (h *Handler) Registration(c *gin.Context) {
	var body registrationRequestBody
	err := json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil {
		message := "Bad request"
		c.JSON(http.StatusBadRequest, gin.H{"message": message})
		return
	}
	err = validateRegistrationBody(c.Writer, body)
	if err != nil {
		return
	}

	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		message := "Cannot hash password"
		c.JSON(http.StatusBadRequest, gin.H{"message": message})
		return
	}

	id, errorMessage := h.services.UserService.CreateUser(body.Email, hashedPassword)
	if errorMessage != "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": errorMessage})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "ok",
		"id":      id,
	})
}

func validateRegistrationBody(writer http.ResponseWriter, body registrationRequestBody) error {
	if len(body.Email) < 5 {
		message := "email must be at least 11 symbols"
		utils.ReturnBadRequestError(writer, http.StatusBadRequest, "message", message)
		return errors.New(message)
	}

	if len(body.Password) < 8 {
		message := "password number must be at least 8 symbols"
		utils.ReturnBadRequestError(writer, http.StatusBadRequest, "message", message)
		return errors.New(message)
	}

	if body.Password != body.ConfirmPassword {
		message := "password should be equal to confirm password"
		utils.ReturnBadRequestError(writer, http.StatusBadRequest, "message", message)
		return errors.New(message)
	}
	return nil
}
