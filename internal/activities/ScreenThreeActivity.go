package activities

import (
	"OnBoardingPOC/internal/contracts"
	"context"
	"encoding/json"
	"go.uber.org/cadence/activity"
	"log"
)

func init() {
	activity.Register(ScreenThreeActivity)
}

func ScreenThreeActivity(ctx context.Context, inputData string) (string, error) {
	activity.GetLogger(ctx).Info("process for condition 1")
	var requestData contracts.RequestData
	json.Unmarshal([]byte(inputData), &requestData)
	data := requestData.TemplateData
	log.Println("Hey its in Screen Three", data.ScreenName)
	templateData := contracts.TemplateData{
		ScreenName: "welcome.html",
		Fields:     nil,
	}
	responseData := contracts.ResponseData{
		TemplateData: templateData,
		Message:      "WelCome User to Onboarding Service!",
	}
	response, err := json.Marshal(responseData)
	return string(response), err
}
