package services

import (
	"server/models"
	bots "server/services/bots"
)

type ReportService struct {
	bots []bots.Bot
}

func NewReportService(bots []bots.Bot) (*ReportService, error) {
	return &ReportService{
		bots: bots,
	}, nil
}

func (this *ReportService) Start() {
	for _, bot := range this.bots {
		go bot.ServeEvents()
	}
}

func (this *ReportService) Stop() {
	for _, bot := range this.bots {
		bot.Quit()
	}
}

func (this *ReportService) Process(report models.Report) {
	for _, bot := range this.bots {
		bot.SendReport(report)
	}
}
