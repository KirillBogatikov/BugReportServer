package main

import (
	"encoding/json"
	"fmt"
	"os"
	"server/database"
	"server/models"
	"server/services"
	"server/services/bots"
	"strconv"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		fmt.Println(err)
		return
	}

	repository, err := database.NewPgRepository(os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err)
		return
	}

	settings := models.Settings{}
	err = json.Unmarshal([]byte(os.Getenv("SETTINGS")), &settings)
	if err != nil {
		fmt.Println(err)
		return
	}

	activeBots := make([]bots.Bot, 0)
	for _, botSettings := range settings.Bots {
		switch botSettings.Type {
		case "telegram":
			telegramBot, err := bots.NewTelegramBot(repository, botSettings.Token, botSettings.Password)
			if err == nil {
				activeBots = append(activeBots, telegramBot)
			} else {
				fmt.Println(err)
			}
		}
	}

	doneSignal := make(chan int)

	service, err := services.NewReportService(activeBots)
	if err != nil {
		fmt.Println(err)
		return
	}
	service.Start()

	server, err := NewServer(settings.Clients, service)
	if err != nil {
		fmt.Println(err)
		return
	}

	server.Start(port)

	<-doneSignal
	repository.Close()
}
