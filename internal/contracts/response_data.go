package contracts

type ResponseData struct {
	TemplateData TemplateData `json:"template_data"`
	Message      string       `json:"message"`
}
