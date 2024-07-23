package cidsdk

type LogMessageRequest struct {
	Level   string                 `json:"level"`
	Message string                 `json:"message"`
	Context map[string]interface{} `json:"context"`
}

// Log request
func (sdk SDK) Log(req LogMessageRequest) error {
	resp, err := sdk.client.R().
		SetHeader("Accept", "application/json").
		SetBody(req).
		SetError(&APIError{}).
		Post("/v1/log")

	if err != nil {
		return err
	} else if resp.IsSuccess() {
		return nil
	} else {
		return resp.Error().(*APIError)
	}
}
