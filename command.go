package cidsdk

// ExecuteCommandRequest defines model for ExecuteCommandRequest.
type ExecuteCommandRequest struct {
	// CaptureOutput capture and return the output (stdout and stderr will be passed through if not set)
	CaptureOutput bool `json:"capture_output,omitempty"`

	// Command command
	Command string `json:"command,omitempty"`

	// WorkDir directory to execute the command in (default = project root)
	WorkDir string `json:"work_dir,omitempty"`

	// Env contains additional env properties
	Env map[string]string `json:"env,omitempty"`

	// Ports that will be exposed
	Ports []int `json:"ports,omitempty"`
}

// ExecuteCommandResponse defines model for ExecuteCommandResponse.
type ExecuteCommandResponse struct {
	// Code command exit code
	Code int `json:"code,omitempty"`

	// Command the command being executed
	Command string `json:"command,omitempty"`

	// Dir directory the command is executed in
	Dir string `json:"dir,omitempty"`

	// Error error message
	Error string `json:"error,omitempty"`

	// Stderr error output (if capture-output was request, empty otherwise)
	Stderr string `json:"stderr,omitempty"`

	// Stdout standard output (if capture-output was request, empty otherwise)
	Stdout string `json:"stdout,omitempty"`
}

// ExecuteCommand command
func (sdk SDK) ExecuteCommand(req ExecuteCommandRequest) (*ExecuteCommandResponse, error) {
	resp, err := sdk.client.R().
		SetHeader("Accept", "application/json").
		SetBody(req).
		SetResult(&ExecuteCommandResponse{}).
		SetError(&APIError{}).
		Post("/command")

	if err != nil {
		return nil, err
	} else if resp.IsSuccess() {
		return resp.Result().(*ExecuteCommandResponse), nil
	} else {
		return nil, resp.Error().(*APIError)
	}
}
