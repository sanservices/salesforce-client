package salesforce

type TokenResponse struct {
	AccessToken     string `json:"access_token"`
	TokenType       string `json:"token_type"`
	ExpiresIn       int    `json:"expires_in"`
	Scope           string `json:"scope"`
	SOAPInstanceURL string `json:"soap_instance_url"`
	RESTInstanceURL string `json:"rest_instance_url"`
}

type InsertDataRowsResponse struct {
	RequestID      string   `json:"requestId"`
	ResultMessages []string `json:"resultMessages"`
}

type CheckInsertRowsStatus struct {
	Page           int                         `json:"page"`
	PageSize       int                         `json:"pageSize"`
	Count          int                         `json:"count"`
	Items          []CheckInsertRowsStatusItem `json:"items"`
	RequestID      string                      `json:"requestId"`
	ResultMessages []string                    `json:"resultMessages"`
}

type SendEmailResponse struct {
	EventInstanceId *string               `json:"eventInstanceId,omitempty" validate:"omitempty" `
	Response        *DeleteSchemaResponse `json:"response,omitempty" validate:"omitempty"`
}

type DeleteSchemaResponse struct {
	Documentation string
	Errorcode     int
	Message       string
}
type CheckInsertRowsStatusItem struct {
	Message   string `json:"message"`
	Status    string `json:"status"`
	ErrorCode *int   `json:"errorCode,omitempty"`
}
