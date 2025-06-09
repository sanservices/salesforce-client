package salesforceclient

type SendEmailRequest struct {
	ContactKey         string      `json:"ContactKey,omitempty"`
	EventDefinitionKey string      `json:"EventDefinitionKey,omitempty"`
	Data               interface{} `json:"Data,omitempty"`
}
