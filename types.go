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
	Metadata() ActionMetadata
	Execute() error
}

type ActionMetadata struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Category    string            `json:"category"`
	Scope       ActionScope       `json:"scope"`
	Links       map[string]string `json:"links,omitempty"`
	Rules       []ActionRule      `json:"rules,omitempty"`  // Rules define conditions that must be met for the action to be executed
	Access      ActionAccess      `json:"access,omitempty"` // Access defines resources that the action may access
	Input       ActionInput       `json:"input,omitempty"`  // Input defines the inputs that the action may consume
	Output      ActionOutput      `json:"output,omitempty"` // Output defines the outputs that the action may produce
}

type ActionScope string

const (
	ActionScopeProject ActionScope = "project"
	ActionScopeModule  ActionScope = "module"
)

type ActionRule struct {
	Type       string `json:"type"`
	Expression string `json:"expression"`
}

type ActionAccess struct {
	Environment []ActionAccessEnv        `json:"env,omitempty"`         // Environment variables that the action may access during execution
	Executables []ActionAccessExecutable `json:"executables,omitempty"` // Executables that the action may invoke during execution
	Network     []ActionAccessNetwork    `json:"network,omitempty"`     // Network access that the action may use during execution
}

type ActionAccessEnv struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Pattern     bool   `json:"pattern,omitempty"`
	Required    bool   `json:"required,omitempty"`
}

type ActionAccessExecutable struct {
	Name       string `json:"name"`
	Constraint string `json:"constraint,omitempty"`
}

type ActionAccessNetwork struct {
	Host string `json:"host"`
}

type ActionInput struct {
	Artifacts []ActionArtifactType `json:"artifacts,omitempty"`
}

type ActionOutput struct {
	Artifacts []ActionArtifactType `json:"artifacts,omitempty"`
}

type ActionArtifactType struct {
	Type          string `json:"type"`             // Type, e.g. "report", "binary"
	Format        string `json:"format,omitempty"` // Format, e.g. "sarif"
	FormatVersion string `json:"format_version,omitempty"`
}

// HealthcheckResponse defines model for HealthcheckResponse.
type HealthcheckResponse struct {
	Status string `json:"status"`
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
	ProjectDir   string            `json:"project_dir,omitempty"`
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
	ProjectDir        string                   `json:"project_dir,omitempty"`         // ProjectDir project root directory
	ModuleDir         string                   `json:"module_dir,omitempty"`          // ModuleDir module root directory
	Discovery         []ProjectModuleDiscovery `json:"discovery,omitempty"`           // Discovery module detected based on
	Name              string                   `json:"name,omitempty"`                // Name module name
	Slug              string                   `json:"slug,omitempty"`                // Slug module name
	Type              string                   `json:"type,omitempty"`                // Type is the module type
	BuildSystem       string                   `json:"build_system,omitempty"`        // BuildSystem module name
	BuildSystemSyntax string                   `json:"build_system_syntax,omitempty"` // BuildSystemSyntax module name
	SpecificationType string                   `json:"specification_type,omitempty"`  // SpecificationType is the type of the specification
	ConfigType        string                   `json:"config_type,omitempty"`         // ConfigType is the type of the configuration
	DeploymentSpec    string                   `json:"deployment_spec,omitempty"`     // DeploymentSpec is the kind of deployment specification
	DeploymentType    string                   `json:"deployment_type,omitempty"`     // DeploymentType is the type of the deployment
	Language          *map[string]string       `json:"language,omitempty"`            // Language module name
	Dependencies      *[]ProjectDependency     `json:"dependencies,omitempty"`        // Dependencies module name
	Files             []string                 `json:"files,omitempty"`               // Files all files in the project directory
	Submodules        *[]ProjectModule         `json:"submodules,omitempty"`          // Submodules submodules
}

// ProjectModuleDiscovery contains info on the files used to discover the module
type ProjectModuleDiscovery struct {
	File string `json:"file"`
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

type VCSDiff struct {
	FileFrom VCSFile       `json:"file_from"`
	FileTo   VCSFile       `json:"file_to"`
	Lines    []VCSDiffLine `json:"lines,omitempty"`
}

type VCSDiffLine struct {
	Operation int    `json:"operation"`
	Content   string `json:"content"`
}

// ActionArtifact contains information about generated artifacts
type ActionArtifact struct {
	BuildID       string `json:"build_id,omitempty"`
	JobID         string `json:"job_id,omitempty"`
	ID            string `json:"id,omitempty"`
	Module        string `json:"module,omitempty"`
	Type          string `json:"type,omitempty"`
	Name          string `json:"name,omitempty"`
	Format        string `json:"format,omitempty"`
	FormatVersion string `json:"format_version,omitempty"`
}
