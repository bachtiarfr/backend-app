package handlers

import (
	"backend-app/application/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type SwipeHandler struct {
	swipeService services.SwipeService
}

func NewSwipeHandler(swipeService services.SwipeService) *SwipeHandler {
	return &SwipeHandler{swipeService}
}

// Swipe godoc
// @Summary Swipe left or right on a profile
// @Description Swipe left (pass) or right (like) on a profile.
// @Tags swipe
// @Accept json
// @Produce json
// @Param profileID path int true "Profile ID to swipe"
// @Param action path string true "Swipe action: 'like' or 'pass'"
// @Param Authorization header string true "Authorization token"
// @Security ApiKeyAuth
// @Router /auth/swipe/{profileID}/{action} [post]
func (h *SwipeHandler) Swipe(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in context"})
		return
	}

	profileID, err := strconv.Atoi(c.Param("profileID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid profileID"})
		return
	}

	action := c.Param("action")
	if action != "like" && action != "pass" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid action"})
		return
	}

	err = h.swipeService.Swipe(userID.(uint), uint(profileID), action)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "swipe recorded"})
}
