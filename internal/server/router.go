package server

import (
	"subscriptionTracker/internal/subscriptions"

	"github.com/gin-gonic/gin"
)

func NewRouter(subscriptionHandler *subscriptions.Handler) *gin.Engine {
	r := gin.Default()
	r.POST("/subscriptions", subscriptionHandler.CreateSubscription)
	return r
}
