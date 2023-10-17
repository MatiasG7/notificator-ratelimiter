package notificator

import (
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

type Notificator struct {
	limiters map[string]*rate.Limiter
}

func NewNotificator() *Notificator {
	return &Notificator{
		limiters: make(map[string]*rate.Limiter)}
}

func (n *Notificator) Send(notifType, userID, message string) error {
	// Here should be parameters validation.
	limiter, ok := n.limiters[notifType+userID]
	if !ok {
		limiter = rate.NewLimiter(getLimitAndBurstByNotifType(notifType))
		n.limiters[notifType+userID] = limiter
	}
	if !limiter.Allow() {
		fmt.Println("not allowed notification to ", userID, ", with message", message, ", type:", notifType)
		return fmt.Errorf("rejected notification to userID: %s with message: %s and type: %s", userID, message, notifType)
	}
	fmt.Println("sending notification to ", userID, ", with message", message, ", type:", notifType)
	return nil
}

func getLimitAndBurstByNotifType(notifType string) (rate.Limit, int) {
	// This could be better, notification types could be constants, and they could even be passed as constructor parameters.
	// All the notification type config could be in a structure inside Notificator in order to set it in main.go and not anywhere else.
	// I use this configuration in order to facilitate the test.
	switch notifType {
	case "news":
		return rate.Every(24 * time.Hour), 1
	case "status":
		return rate.Every(2 * time.Minute), 2
	case "project_invitation":
		return rate.Every(10 * time.Second), 2
	case "marketing":
		return rate.Every(1 * time.Minute), 2
	default:
		return rate.Every(1 * time.Hour), 1
	}
}
