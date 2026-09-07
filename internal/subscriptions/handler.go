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

func (h *Handler) UpdateSubscription(c *gin.Context) {
	ownerId := c.Param("uuid")
	var req SubscriptionsUpdateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sub, err := h.repo.UpdateSubscription(c.Request.Context(), req, ownerId)
	if err != nil {
		log.Printf("create subscription: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update subscription"})
		return
	}

	c.JSON(http.StatusCreated, sub)

}

func (h *Handler) DeleteSubscription(c *gin.Context) {
	uuid := c.Param("uuid")

	err := h.repo.DeleteSubscription(c.Request.Context(), uuid)

	if err != nil {
		log.Printf("create subscription: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete subscription"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "subscription deleted successfully"})

}

func (h *Handler) Getsubscriptions(c *gin.Context) {

	owner_id := c.Param("owner_id")

	sub, err := h.repo.Getsubscriptions(c.Request.Context(), owner_id)

	if err != nil {
		log.Printf("create subscription: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrive subscription"})
		return
	}

	c.JSON(http.StatusOK, sub)

}
