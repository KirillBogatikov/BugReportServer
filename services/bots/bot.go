package bots

import "server/models"

type Bot interface {
	SendReport(report models.Report)
	ServeEvents()
	Quit()
}
