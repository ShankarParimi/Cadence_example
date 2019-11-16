package contracts

type Request struct {
	Data       RequestData `json:"data"`
	WorkflowId string      `json:"workflow_id"`
}
