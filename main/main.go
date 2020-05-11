package main

import (
	"encoding/json"
	. "github.com/KirillBogatikov/Cuba/log/go"
	. "github.com/KirillBogatikov/Cuba/log/go/common"
	. "github.com/KirillBogatikov/Cuba/log/go/fmt"
	"os"
	"server/database"
	"server/models"
	"server/services"
	"server/services/bots"
	"strconv"
)

func makeLog() Log {
	stream := NewStream(nil)
	cfg := Configuration{
		Debug: stream,
		Info:  stream,
		Warn:  stream,
		Error: stream,
	}
	return NewLog(cfg)
}

const (
	TAG = "BugReporter"
)

func main() {
	log := makeLog()

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.ErrorE(TAG, "Can not get PORT", err)
		return
	}

	repository, err := database.NewPgRepository(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.ErrorE(TAG, "Can not get DATABASE_URL", err)
		return
	}

	settings := models.Settings{}
	err = json.Unmarshal([]byte(os.Getenv("SETTINGS")), &settings)
	if err != nil {
		log.ErrorE(TAG, "Can not get SETTINGS json", err)
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
				log.ErrorE(TAG, "Can not create Telegram Bot", err)
			}
		}
	}

	doneSignal := make(chan int)

	service, err := services.NewReportService(activeBots)
	if err != nil {
		log.ErrorE(TAG, "Can not create Report Service", err)
		return
	}
	service.Start()

	server, err := NewServer(settings.Clients, service)
	if err != nil {
		log.ErrorE(TAG, "Can not create Server", err)
		return
	}

	server.Start(port)

	<-doneSignal
	err = repository.Close()
	if err != nil {
		log.WarnE(TAG, "Repository doesn't closed", err)
	}
}
