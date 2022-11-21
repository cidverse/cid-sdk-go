package sdk

import (
	"strconv"
)

// APIError defines model for Error.
type APIError struct {
	Detail *string `json:"detail,omitempty"`
	Status *int    `json:"status,omitempty"`
	Title  *string `json:"title,omitempty"`
}

func (e *APIError) Error() string {
	return strconv.Itoa(*e.Status) + ": " + *e.Title
}

// HealthcheckResponse defines model for HealthcheckResponse.
type HealthcheckResponse struct {
	Status string `json:"status"`
}

// ProjectDependency defines model for ProjectDependency.
type ProjectDependency struct {
	Id      *string `json:"id,omitempty"`
	Type    *string `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

// ProjectEnvResponse defines model for ProjectEnvResponse.
type ProjectEnvResponse map[string]string

// ProjectInfoResponse defines model for ProjectInfoResponse.
type ProjectInfoResponse struct {
	// ProjectDir project root directory
	ProjectDir *string `json:"project_dir,omitempty"`

	// WorkDir current working directory
	WorkDir *string `json:"work_dir,omitempty"`

	// UserDisplayName display name of the current user
	UserDisplayName *string `json:"user_display_name,omitempty"`

	// UserGroupId group id
	UserGroupId *string `json:"user_group-id,omitempty"`

	// UserId user id
	UserId *string `json:"user_id,omitempty"`

	// UserLoginName login name of the current user
	UserLoginName *string `json:"user_login_name,omitempty"`
}

// ProjectModule defines model for ProjectModule.
type ProjectModule struct {
	// ProjectDir project root directory
	ProjectDir *string `json:"project_dir,omitempty"`

	// ModuleDir module root directory
	ModuleDir *string `json:"module_dir,omitempty"`

	// Discovery module detected based on
	Discovery *[]string `json:"discovery,omitempty"`

	// Name module name
	Name *string `json:"name,omitempty"`

	// Slug module name
	Slug *string `json:"slug,omitempty"`

	// BuildSystem module name
	BuildSystem *string `json:"build_system,omitempty"`

	// BuildSystemSyntax module name
	BuildSystemSyntax *string `json:"build_system_syntax,omitempty"`

	// Language module name
	Language *map[string]string `json:"language,omitempty"`

	// Dependencies module name
	Dependencies *[]ProjectDependency `json:"dependencies,omitempty"`

	// Files all files in the project directory
	Files *[]string `json:"files,omitempty"`

	// Submodules submodules
	Submodules *[]ProjectModule `json:"submodules,omitempty"`
}

type ModuleListResponse []ProjectModule

type ModuleCurrentResponse ProjectModule

// ExecuteCommandRequest defines model for ExecuteCommandRequest.
type ExecuteCommandRequest struct {
	// CaptureOutput capture and return the output (stdout and stderr will be passed thru if not set)
	CaptureOutput *bool `json:"capture-output,omitempty"`

	// Command command
	Command *string `json:"command,omitempty"`

	// WorkDir directory to execute the command in (default = project root)
	WorkDir *string `json:"work-dir,omitempty"`
}

// ExecuteCommandResponse defines model for ExecuteCommandResponse.
type ExecuteCommandResponse struct {
	// Code command exit code
	Code *float32 `json:"code,omitempty"`

	// Command the command being executed
	Command *string `json:"command,omitempty"`

	// Dir directory the command is executed in
	Dir *string `json:"dir,omitempty"`

	// Error error message
	Error *string `json:"error,omitempty"`

	// Stderr error output (if capture-output was request, empty otherwise)
	Stderr *string `json:"stderr,omitempty"`

	// Stdout standard output (if capture-output was request, empty otherwise)
	Stdout *string `json:"stdout,omitempty"`
}
