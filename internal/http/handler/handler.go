package handler

import (
	"github.com/nozimamax/fitness-tracking-app/storage"
	"github.com/nozimamax/fitness-tracking-app/internal/email"
)

type HandlerST struct {
	Queries *storage.Queries
	Notification email.NotificationRepo
}