package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"server/models"
	"server/services"
	"server/utils"
	"time"
)

type ReportController struct {
	Service *services.ReportService
}

func NewReportController(service *services.ReportService) *ReportController {
	return &ReportController{
		Service: service,
	}
}

func (this *ReportController) Handle(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()

	appName := request.Header.Get("App-Name")
	if len(appName) == 0 {
		utils.BadRequest(writer)
		return
	}

	if err != nil {
		utils.BadRequest(writer)
		return
	}

	report := models.Report{}
	err = json.Unmarshal(body, &report)
	if err != nil {
		fmt.Println(err)
		utils.BadRequest(writer)
		return
	}

	report.ServedTime = time.Now()
	report.AppName = appName
	this.Service.Process(report)
	utils.Ok(writer)
}
