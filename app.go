package main

import (
	"bestie-reminder-app/services"
	"context"
	"fmt"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	notifications := services.NewNotificationService()
	notifications.StartNotifications()
	a.ctx = ctx
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
