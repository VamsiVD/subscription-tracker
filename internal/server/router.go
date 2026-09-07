package server

import (
	"subscriptionTracker/internal/subscriptions"

	"github.com/gin-gonic/gin"
)

func NewRouter(subscriptionHandler *subscriptions.Handler) *gin.Engine {
	r := gin.Default()
	r.POST("/subscriptions", subscriptionHandler.CreateSubscription)
	r.PATCH("/subscriptions/:uuid", subscriptionHandler.UpdateSubscription)
	r.DELETE("/subscriptions/:uuid", subscriptionHandler.DeleteSubscription)
	r.GET("/subscriptions/:owner_id", subscriptionHandler.Getsubscriptions)
	return r
}
