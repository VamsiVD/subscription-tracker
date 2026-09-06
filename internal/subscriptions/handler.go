package subscriptions

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo *Respository
}

func NewHandler(repo *Respository) *Handler {
	return &Handler{repo: repo}

}

func (h *Handler) CreateSubscription(c *gin.Context) {

	var req SubscriptionsCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sub, err := h.repo.CreateSubscription(c.Request.Context(), req)
	if err != nil {
		log.Printf("create subscription: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create subscription"})
		return
	}

	c.JSON(http.StatusCreated, sub)

}
