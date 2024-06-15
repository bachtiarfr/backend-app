package handlers

import (
	"backend-app/application/dto"
	"backend-app/application/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService}
}

// SignUp godoc
// @Summary Sign up a new user
// @Description Create a new user with the provided details
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body dto.UserDTO true "User Data"
// @Success 201 {object} dto.UserDTO
// @Router /signup [post]
func (h *UserHandler) SignUp(c *gin.Context) {
	var userDTO dto.UserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userService.SignUp(userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, userDTO)
}

// Login godoc
// @Summary Log in a user
// @Description Authenticate a user and return a JWT token
// @Tags user
// @Accept  json
// @Produce  json
// @Param login body dto.UserDTO true "Login Data"
// @Router /login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var loginDTO dto.UserDTO
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.userService.Login(loginDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// PurchasePremium godoc
// @Summary Purchase a premium package
// @Description Allows a user to purchase a premium package
// @Tags user
// @Accept  json
// @Produce  json
// @Param purchase body dto.PurchaseDTO true "Purchase Data"
// @Param Authorization header string true "Authorization token"
// @Security ApiKeyAuth
// @Router /auth/purchase [post]
func (h *UserHandler) PurchasePremium(c *gin.Context) {
	var purchaseDTO dto.PurchaseDTO
	if err := c.ShouldBindJSON(&purchaseDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve user ID from the context (assumes user ID is stored in the context by the auth middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})
		return
	}
	fmt.Println("userID :", userID)
	err := h.userService.PurchasePremium(userID.(uint), purchaseDTO.Package)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Package purchased successfully"})
}
