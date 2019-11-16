package admin

import (
	"OnBoardingPOC/internal/activities"
	"OnBoardingPOC/internal/common"
	"OnBoardingPOC/internal/contracts"
	"context"
	"encoding/json"
	"github.com/pborman/uuid"
	"go.uber.org/cadence"
	"go.uber.org/cadence/client"
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
	"log"
	"time"
)

func init() {
	workflow.Register(SignUpWorkflow)
}

const ApplicationName = "adminGroup"

func SignUpWorkflow(ctx workflow.Context) error {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Second,
		StartToCloseTimeout:    time.Minute * 10,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	//Waiting For Signal(User Input to be given)
	ch := workflow.GetSignalChannel(ctx, SignalName)
	e := workflow.SetQueryHandler(ctx, "status", func(input []byte) (string, error) {
		return "WAITING", nil
	})

	if e != nil {
		logger.Info("SetQueryHandler failed: " + e.Error())
		return e
	}
	var inputData string
	//Screen 1 Task Waiting
	if more := ch.Receive(ctx, &inputData); !more {
		logger.Info("Signal channel closed")
		return cadence.NewCustomError("signal_channel_closed")
	}
	logger.Info("Signal received.", zap.String("signal", inputData))
	var processResult string
	err := workflow.ExecuteActivity(ctx, activities.ScreenOneActivity, inputData).Get(ctx, &processResult)
	if err != nil {
		return err
	}
	e1 := workflow.SetQueryHandler(ctx, "status", func(input []byte) (string, error) {
		return processResult, nil
	})
	if e1 != nil {
		logger.Info("SetQueryHandler failed: " + e1.Error())
		return e1
	}

	// Screen 2 Task Waiting..
	if more := ch.Receive(ctx, &inputData); !more {
		logger.Info("Signal channel closed")
		return cadence.NewCustomError("signal_channel_closed")
	}
	e = workflow.SetQueryHandler(ctx, "status", func(input []byte) (string, error) {
		return "WAITING", nil
	})

	if e != nil {
		logger.Info("SetQueryHandler failed: " + e.Error())
		return e
	}
	var activity2Result string
	err = workflow.ExecuteActivity(ctx, activities.ScreenTwoActivity, inputData).Get(ctx, &activity2Result)
	if err != nil {
		return err
	}
	e = workflow.SetQueryHandler(ctx, "status", func(input []byte) (string, error) {
		return activity2Result, nil
	})
	if e != nil {
		logger.Info("SetQueryHandler failed: " + e.Error())
		return e
	}

	return nil
}

func StartWorkflow(ctx context.Context, h *common.SampleHelper) ([]byte, error) {
	workflowOptions := client.StartWorkflowOptions{
		ID:                              "onboarding_" + uuid.New(),
		TaskList:                        ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute * 10,
		DecisionTaskStartToCloseTimeout: time.Minute,
		WorkflowIDReusePolicy:           client.WorkflowIDReusePolicyAllowDuplicate,
	}
	workflowID, runID := h.StartWorkflow(workflowOptions, SignUpWorkflow)
	log.Print(runID)
	fields := []string{"username", "password", "email"}
	templateData := contracts.TemplateData{
		ScreenName: "signUp.html",
		Fields:     fields,
	}
	data := contracts.ResponseData{
		TemplateData: templateData,
		Message:      "",
	}
	response := contracts.Response{
		Data:       data,
		WorkflowId: workflowID,
	}

	return json.Marshal(response)
}

func Process(workflowId string, inputJson string, h *common.SampleHelper) ([]byte, error) {
	h.SignalWorkflow(workflowId, SignalName, inputJson)
	var respData contracts.ResponseData
	data := h.QueryWorkflow(workflowId, "", "status")
	s := data.(string)
	json.Unmarshal([]byte(s), &respData)
	response := contracts.Response{
		Data:       respData,
		WorkflowId: workflowId,
	}
	return json.Marshal(response)

}
