package cidsdk

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

func NewSDK() (*SDK, error) {
	client := resty.New()
	client.SetDisableWarn(true)
	client.SetRetryCount(0)
	client.SetTimeout(60 * time.Minute)

	// socket endpoint
	unixSocket := os.Getenv("CID_API_SOCKET")
	if len(unixSocket) > 0 {
		transport := http.Transport{
			DialContext: func(_ context.Context, _ string, _ string) (net.Conn, error) {
				return net.Dial("unix", unixSocket)
			},
		}

		// base url
		client.SetTransport(&transport).SetBaseURL("http://localhost")
	}

	// http endpoint
	addr := os.Getenv("CID_API_ADDR")
	if len(addr) > 0 {
		client.SetBaseURL(addr)
	}

	// auth
	secret := os.Getenv("CID_API_SECRET")
	if len(secret) > 0 {
		client.SetAuthScheme("Bearer")
		client.SetAuthToken(secret)
	}

	if client == nil {
		return nil, errors.New("failed to initialize api client")
	}
	return &SDK{client: client}, nil
}

type SDKClient interface {
	Health() (*HealthcheckResponse, error)
	Log(req LogMessageRequest) error
	ModuleAction(cfg any) (ModuleActionData, error)
	ProjectAction(cfg any) (ProjectActionData, error)
	Env() (map[string]string, error)
	Modules() (*[]ProjectModule, error)
	CurrentModule() (*ProjectModule, error)
	CurrentConfig() (*CurrentConfig, error)
	VCSCommits(request VCSCommitsRequest) (*[]VCSCommit, error)
	VCSCommitByHash(request VCSCommitByHashRequest) (*VCSCommit, error)
	VCSTags() (*[]VCSTag, error)
	VCSReleases() (*[]VCSRelease, error)
	ExecuteCommand(req ExecuteCommandRequest) (*ExecuteCommandResponse, error)
	FileRead(file string) (string, error)
	FileList(req FileRequest) (files []File, err error)
	FileRename(old string, new string) error
	FileRemove(file string) error
	FileWrite(file string, content []byte) error
	ArtifactList(request ArtifactListRequest) (*[]ActionArtifact, error)
	ArtifactUpload(request ArtifactUploadRequest) error
	ArtifactDownload(request ArtifactDownloadRequest) error
	UUID() string
}

type SDK struct {
	client *resty.Client
}

// Health request
func (sdk SDK) Health() (*HealthcheckResponse, error) {
	resp, err := sdk.client.R().
		SetHeader("Accept", "application/json").
		SetResult(&HealthcheckResponse{}).
		SetError(&APIError{}).
		Get("/health")

	if err != nil {
		return nil, err
	} else if resp.IsSuccess() {
		fmt.Printf("RESULT: %+v\n", resp.Result())
		return resp.Result().(*HealthcheckResponse), nil
	} else {
		return nil, resp.Error().(*APIError)
	}
}

// Env request
func (sdk SDK) Env() (map[string]string, error) {
	resp, err := sdk.client.R().
		SetHeader("Accept", "application/json").
		SetResult(&map[string]string{}).
		SetError(&APIError{}).
		Get("/env")

	if err != nil {
		return nil, err
	} else if resp.IsSuccess() {
		return *resp.Result().(*map[string]string), nil
	} else {
		return nil, resp.Error().(*APIError)
	}
}

// CurrentConfig request
func (sdk SDK) CurrentConfig() (*CurrentConfig, error) {
	resp, err := sdk.client.R().
		SetHeader("Accept", "application/json").
		SetResult(&CurrentConfig{}).
		SetError(&APIError{}).
		Get("/config/current")

	if err != nil {
		return nil, err
	} else if resp.IsSuccess() {
		return resp.Result().(*CurrentConfig), nil
	} else {
		return nil, resp.Error().(*APIError)
	}
}

// Modules request
func (sdk SDK) Modules() (*[]ProjectModule, error) {
	resp, err := sdk.client.R().
		SetHeader("Accept", "application/json").
		SetResult(&[]ProjectModule{}).
		SetError(&APIError{}).
		Get("/module")

	if err != nil {
		return nil, err
	} else if resp.IsSuccess() {
		return resp.Result().(*[]ProjectModule), nil
	} else {
		return nil, resp.Error().(*APIError)
	}
}

// CurrentModule request
func (sdk SDK) CurrentModule() (*ProjectModule, error) {
	resp, err := sdk.client.R().
		SetHeader("Accept", "application/json").
		SetResult(&ProjectModule{}).
		SetError(&APIError{}).
		Get("/module/current")

	if err != nil {
		return nil, err
	} else if resp.IsSuccess() {
		return resp.Result().(*ProjectModule), nil
	} else {
		return nil, resp.Error().(*APIError)
	}
}

type VCSCommitsRequest struct {
	FromHash       string `json:"from"`
	ToHash         string `json:"to"`
	IncludeChanges bool   `json:"changes"`
	Limit          int    `json:"limit"`
}

// VCSCommits request
func (sdk SDK) VCSCommits(request VCSCommitsRequest) (*[]VCSCommit, error) {
	resp, err := sdk.client.R().
		SetHeader("Accept", "application/json").
		SetResult(&[]VCSCommit{}).
		SetError(&APIError{}).
		Get(fmt.Sprintf("/vcs/commit?from=%s&to=%schanges=%s&limit=%s", request.FromHash, request.ToHash, strconv.FormatBool(request.IncludeChanges), strconv.Itoa(request.Limit)))

	if err != nil {
		return nil, err
	} else if resp.IsSuccess() {
		return resp.Result().(*[]VCSCommit), nil
	} else {
		return nil, resp.Error().(*APIError)
	}
}

type VCSCommitByHashRequest struct {
	Hash           string `json:"hash"`
	IncludeChanges bool   `json:"changes"`
}

// VCSCommitByHash request
func (sdk SDK) VCSCommitByHash(request VCSCommitByHashRequest) (*VCSCommit, error) {
	resp, err := sdk.client.R().
		SetHeader("Accept", "application/json").
		SetResult(&VCSCommit{}).
		SetError(&APIError{}).
		Get("/vcs/commit/" + request.Hash + "?changes=" + strconv.FormatBool(request.IncludeChanges))

	if err != nil {
		return nil, err
	} else if resp.IsSuccess() {
		return resp.Result().(*VCSCommit), nil
	} else {
		return nil, resp.Error().(*APIError)
	}
}

// VCSTags request
func (sdk SDK) VCSTags() (*[]VCSTag, error) {
	resp, err := sdk.client.R().
		SetHeader("Accept", "application/json").
		SetResult(&[]VCSTag{}).
		SetError(&APIError{}).
		Get("/vcs/tag")

	if err != nil {
		return nil, err
	} else if resp.IsSuccess() {
		return resp.Result().(*[]VCSTag), nil
	} else {
		return nil, resp.Error().(*APIError)
	}
}

// VCSReleases request
func (sdk SDK) VCSReleases() (*[]VCSRelease, error) {
	resp, err := sdk.client.R().
		SetHeader("Accept", "application/json").
		SetResult(&[]VCSRelease{}).
		SetError(&APIError{}).
		Get("/vcs/release")

	if err != nil {
		return nil, err
	} else if resp.IsSuccess() {
		return resp.Result().(*[]VCSRelease), nil
	} else {
		return nil, resp.Error().(*APIError)
	}
}

type ArtifactListRequest struct {
}

// ArtifactList request
func (sdk SDK) ArtifactList(request ArtifactListRequest) (*[]ActionArtifact, error) {
	resp, err := sdk.client.R().
		SetHeader("Accept", "application/json").
		SetResult(&[]ActionArtifact{}).
		SetError(&APIError{}).
		Get("/artifact")

	if err != nil {
		return nil, err
	} else if resp.IsSuccess() {
		return resp.Result().(*[]ActionArtifact), nil
	} else {
		return nil, resp.Error().(*APIError)
	}
}

type ArtifactUploadRequest struct {
	File          string `json:"file"`
	Module        string `json:"module"`
	Type          string `json:"type"`
	Format        string `json:"format"`
	FormatVersion string `json:"format_version"`
}

// ArtifactUpload request
func (sdk SDK) ArtifactUpload(request ArtifactUploadRequest) error {
	f, err := os.Open(request.File)
	if err != nil {
		return err
	}
	defer f.Close()

	// upload
	resp, err := sdk.client.R().
		SetBody(f).
		SetFormData(map[string]string{
			"type":           request.Type,
			"module":         request.Module,
			"format":         request.Format,
			"format_version": request.FormatVersion,
		}).
		SetFile("file", request.File).
		SetContentLength(true).
		SetError(&APIError{}).
		Post("/artifact")
	if err != nil {
		return err
	} else if resp.IsError() {
		return resp.Error().(*APIError)
	}

	return nil
}

type ArtifactDownloadRequest struct {
	Module     string `json:"module"`
	Type       string `json:"type"`
	Name       string `json:"name"`
	TargetFile string `json:"target_file"`
}

// ArtifactDownload request
func (sdk SDK) ArtifactDownload(request ArtifactDownloadRequest) error {
	resp, err := sdk.client.R().
		SetQueryParam("module", request.Module).
		SetQueryParam("type", request.Type).
		SetQueryParam("name", request.Name).
		SetOutput(request.TargetFile).
		SetError(&APIError{}).
		Get("/artifact/download")
	if err != nil {
		return err
	} else if resp.IsError() {
		return resp.Error().(*APIError)
	}

	return nil
}
