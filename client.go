package cidsdk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
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
		client.SetTransport(&transport).SetScheme("http").SetBaseURL(unixSocket)
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
	ProjectEnv() (map[string]string, error)
	Modules() ([]ProjectModule, error)
	CurrentModule() (*ProjectModule, error)
	CurrentConfig() (*CurrentConfig, error)
	VCSCommits(changes bool, limit int) ([]VCSCommit, error)
	VCSCommitByHash(hash string, changes bool) (*VCSCommit, error)
	VCSTags() ([]VCSTag, error)
	VCSReleases() ([]VCSRelease, error)
	ExecuteCommand(req ExecuteCommandRequest) (*ExecuteCommandResponse, error)
	PrepareAction(cfg any) (ActionEnv, error)
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

// Log request
func (sdk SDK) Log(req LogMessageRequest) error {
	resp, err := sdk.client.R().
		SetHeader("Accept", "application/json").
		SetBody(req).
		SetError(&APIError{}).
		Post("/log")

	if err != nil {
		return err
	} else if resp.IsSuccess() {
		return nil
	} else {
		return resp.Error().(*APIError)
	}
}

// PrepareAction is a utility function that prepares some common data for actions
func (sdk SDK) PrepareAction(cfg any) (ActionEnv, error) {
	config, err := sdk.CurrentConfig()
	if err != nil {
		return ActionEnv{}, err
	}

	module, err := sdk.CurrentModule()
	if err != nil {
		return ActionEnv{}, err
	}

	if config.Config != "" && cfg != nil {
		err := json.Unmarshal([]byte(config.Config), cfg)
		if err != nil {
			return ActionEnv{}, err
		}
	}

	return ActionEnv{Config: *config, Module: *module}, nil
}

// ProjectEnv request
func (sdk SDK) ProjectEnv() (map[string]string, error) {
	resp, err := sdk.client.R().
		SetHeader("Accept", "application/json").
		SetResult(&map[string]string{}).
		SetError(&APIError{}).
		Get("/project")

	if err != nil {
		return nil, err
	} else if resp.IsSuccess() {
		return resp.Result().(map[string]string), nil
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
func (sdk SDK) Modules() ([]ProjectModule, error) {
	resp, err := sdk.client.R().
		SetHeader("Accept", "application/json").
		SetResult(&[]ProjectModule{}).
		SetError(&APIError{}).
		Get("/module")

	if err != nil {
		return nil, err
	} else if resp.IsSuccess() {
		return resp.Result().([]ProjectModule), nil
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

// VCSCommits request
func (sdk SDK) VCSCommits(changes bool, limit int) ([]VCSCommit, error) {
	resp, err := sdk.client.R().
		SetHeader("Accept", "application/json").
		SetResult(&[]VCSCommit{}).
		SetError(&APIError{}).
		Get("/vcs/commit?changes=" + strconv.FormatBool(changes) + "&limit=" + strconv.Itoa(limit))

	if err != nil {
		return nil, err
	} else if resp.IsSuccess() {
		return resp.Result().([]VCSCommit), nil
	} else {
		return nil, resp.Error().(*APIError)
	}
}

// VCSCommitByHash request
func (sdk SDK) VCSCommitByHash(hash string, changes bool) (*VCSCommit, error) {
	resp, err := sdk.client.R().
		SetHeader("Accept", "application/json").
		SetResult(&VCSCommit{}).
		SetError(&APIError{}).
		Get("/vcs/commit/" + hash + "?changes=" + strconv.FormatBool(changes))

	if err != nil {
		return nil, err
	} else if resp.IsSuccess() {
		return resp.Result().(*VCSCommit), nil
	} else {
		return nil, resp.Error().(*APIError)
	}
}

// VCSTags request
func (sdk SDK) VCSTags() ([]VCSTag, error) {
	resp, err := sdk.client.R().
		SetHeader("Accept", "application/json").
		SetResult(&[]VCSTag{}).
		SetError(&APIError{}).
		Get("/vcs/tag")

	if err != nil {
		return nil, err
	} else if resp.IsSuccess() {
		return resp.Result().([]VCSTag), nil
	} else {
		return nil, resp.Error().(*APIError)
	}
}

// VCSReleases request
func (sdk SDK) VCSReleases() ([]VCSRelease, error) {
	resp, err := sdk.client.R().
		SetHeader("Accept", "application/json").
		SetResult(&[]VCSRelease{}).
		SetError(&APIError{}).
		Get("/vcs/release")

	if err != nil {
		return nil, err
	} else if resp.IsSuccess() {
		return resp.Result().([]VCSRelease), nil
	} else {
		return nil, resp.Error().(*APIError)
	}
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
