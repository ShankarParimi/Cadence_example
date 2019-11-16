package activities

import (
	"OnBoardingPOC/internal/contracts"
	"context"
	"encoding/json"
	"go.uber.org/cadence/activity"
	"log"
)

func init() {
	activity.Register(ScreenOneActivity)
}

func ScreenOneActivity(ctx context.Context, inputData string) (string, error) {
	activity.GetLogger(ctx).Info("process for condition 1")
	var requestData contracts.RequestData
	json.Unmarshal([]byte(inputData), &requestData)
	fieldsMap := requestData.TemplateData.Fields
	fieldMap := fieldsMap.(map[string]interface{})
	username := fieldMap["username"]
	log.Println(username)
	log.Println("Hey its in SCREEN ONE")
	fields := []string{"mobileNumber", "Address"}
	templateData := contracts.TemplateData{
		ScreenName: "details.html",
		Fields:     fields,
	}
	responseData := contracts.ResponseData{
		TemplateData: templateData,
		Message:      "Saved Successfully!",
	}
	response, err := json.Marshal(responseData)
	return string(response), err
}
