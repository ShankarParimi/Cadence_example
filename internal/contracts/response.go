package contracts

type Response struct {
	Data       interface{} `json:"data"`
	WorkflowId string      `json:"workflow_id"`
}
