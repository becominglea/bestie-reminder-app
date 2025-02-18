package services

import (
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
)

type NotificationService struct {
}

func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

func (n *NotificationService) Notify(title string, body string, image string) {
	fmt.Print(title, body)

	err := beeep.Notify(title, body, image)
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
}
func (n *NotificationService) StartNotifications() {
	go n.startWaterReminder()
	go n.startPillReminder()
	go n.startOutdoorReminder()
}

func (n *NotificationService) startWaterReminder() {
	for {
		n.Notify("Bois de l'eau", "Il est temps de boire un verre d'eau !", "image.png")
		time.Sleep(30 * time.Second)
	}
}

func (n *NotificationService) startPillReminder() {
	for {
		now := time.Now()
		currentTime := now.Format("15:04")

		if currentTime == "10:00" || currentTime == "12:00" || currentTime == "19:00" {
			n.Notify("Prends tes médicaments", "N'oublie pas de prendre tes médicaments.", "image.png")
		}

		time.Sleep(time.Minute)
	}
}

func (n *NotificationService) startOutdoorReminder() {
	for {
		now := time.Now()
		currentTime := now.Format("15:04")

		if currentTime == "16:00" {
			n.Notify("Sors dehors", "Fais une pause et prends l'air frais !", "image.png")
		}

		time.Sleep(time.Minute)
	}
}
