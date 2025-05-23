package cidsdk

// ExecuteCommandRequest defines model for ExecuteCommandRequest.
type ExecuteCommandRequest struct {
	CaptureOutput bool              `json:"capture_output,omitempty"` // CaptureOutput capture and return both stdout and stderr
	HideStdout    bool              `json:"hide_stdout,omitempty"`    // HideStdout hide the stdout output
	HideStderr    bool              `json:"hide_stderr,omitempty"`    // HideStderr hide the stderr output
	Command       string            `json:"command,omitempty"`        // Command
	WorkDir       string            `json:"work_dir,omitempty"`       // WorkDir directory to execute the command in (default = project root)
	Env           map[string]string `json:"env,omitempty"`            // Env contains additional env properties
	Ports         []int             `json:"ports,omitempty"`          // Ports that will be exposed
	Constraint    string            `json:"constraint,omitempty"`     // A version Constraint for the binary used in command
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
		Post("/v1/command/execute")

	if err != nil {
		return nil, err
	} else if resp.IsSuccess() {
		return resp.Result().(*ExecuteCommandResponse), nil
	} else {
		return nil, resp.Error().(*APIError)
	}
}
