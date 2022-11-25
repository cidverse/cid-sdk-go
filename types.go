package cidsdk

import (
	"strconv"
	"time"
)

// APIError defines model for Error.
type APIError struct {
	Detail string `json:"detail,omitempty"`
	Status int    `json:"status,omitempty"`
	Title  string `json:"title,omitempty"`
}

func (e *APIError) Error() string {
	return strconv.Itoa(e.Status) + ": " + e.Title
}

// Action is the common interface for all actions
type Action interface {
	Execute() error
}

// HealthcheckResponse defines model for HealthcheckResponse.
type HealthcheckResponse struct {
	Status string `json:"status"`
}

type ActionEnv struct {
	Module ProjectModule
	Config CurrentConfig
}

// ProjectDependency defines model for ProjectDependency.
type ProjectDependency struct {
	Id      string `json:"id,omitempty"`
	Type    string `json:"type,omitempty"`
	Version string `json:"version,omitempty"`
}

// CurrentConfig defines model for CurrentConfig.
type CurrentConfig struct {
	Debug        bool              `json:"debug,omitempty"`
	Log          map[string]string `json:"log,omitempty"`
	TempDir      string            `json:"temp_dir,omitempty"`
	ArtifactDir  string            `json:"artifact_dir,omitempty"`
	HostName     string            `json:"host_name,omitempty"`
	HostUserId   string            `json:"host_user_id,omitempty"`
	HostUserName string            `json:"host_user_name,omitempty"`
	HostGroupId  string            `json:"host_group_id,omitempty"`
	Config       string            `json:"config,omitempty"`
}

func (c CurrentConfig) DebugFlag(id string, flag string) string {
	if c.Debug || c.Log[id] == "debug" {
		return flag
	}

	return ""
}

// ProjectModule defines model for ProjectModule.
type ProjectModule struct {
	// ProjectDir project root directory
	ProjectDir string `json:"project_dir,omitempty"`

	// ModuleDir module root directory
	ModuleDir string `json:"module_dir,omitempty"`

	// Discovery module detected based on
	Discovery []string `json:"discovery,omitempty"`

	// Name module name
	Name string `json:"name,omitempty"`

	// Slug module name
	Slug string `json:"slug,omitempty"`

	// BuildSystem module name
	BuildSystem string `json:"build_system,omitempty"`

	// BuildSystemSyntax module name
	BuildSystemSyntax string `json:"build_system_syntax,omitempty"`

	// Language module name
	Language *map[string]string `json:"language,omitempty"`

	// Dependencies module name
	Dependencies *[]ProjectDependency `json:"dependencies,omitempty"`

	// Files all files in the project directory
	Files []string `json:"files,omitempty"`

	// Submodules submodules
	Submodules *[]ProjectModule `json:"submodules,omitempty"`
}

type VCSCommit struct {
	HashShort   string            `json:"hash_short,omitempty"`
	Hash        string            `json:"hash,omitempty"`
	Message     string            `json:"message,omitempty"`
	Description string            `json:"description,omitempty"`
	Author      VCSAuthor         `json:"author,omitempty"`
	Committer   VCSAuthor         `json:"committer,omitempty"`
	Tags        *[]VCSTag         `json:"tags,omitempty"`
	AuthoredAt  time.Time         `json:"authored_at,omitempty"`
	CommittedAt time.Time         `json:"committed_at,omitempty"`
	Changes     *[]VCSChange      `json:"changes,omitempty"`
	Context     map[string]string `json:"context,omitempty"`
}

type VCSAuthor struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type VCSTag struct {
	RefType string `json:"type,omitempty"`
	Value   string `json:"value,omitempty"`
	Hash    string `json:"hash,omitempty"`
}

type VCSRelease struct {
	Version string `json:"version,omitempty"`
	Ref     VCSTag `json:"ref,omitempty"`
}

type VCSChange struct {
	ChangeType string  `json:"type,omitempty"`
	FileFrom   VCSFile `json:"file_from,omitempty"`
	FileTo     VCSFile `json:"file_to,omitempty"`
	Patch      string  `json:"patch,omitempty"`
}

type VCSFile struct {
	Name string `json:"name,omitempty"`
	Size int    `json:"size,omitempty"`
	Hash string `json:"hash,omitempty"`
}

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
}

type LogMessageRequest struct {
	Level   string                 `json:"level"`
	Message string                 `json:"message"`
	Context map[string]interface{} `json:"context"`
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
