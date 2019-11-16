package contracts

type TemplateData struct {
	ScreenName string      `json:"screen_name"`
	Fields     interface{} `json:"fields"`
}
