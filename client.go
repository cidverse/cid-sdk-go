package cidsdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net"
	"net/http"
	"os"
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

// ProjectInfo request
func (sdk SDK) ProjectInfo() (*ProjectInfoResponse, error) {
	resp, err := sdk.client.R().
		SetHeader("Accept", "application/json").
		SetResult(&ProjectInfoResponse{}).
		SetError(&APIError{}).
		Get("/project")

	if err != nil {
		return nil, err
	} else if resp.IsSuccess() {
		return resp.Result().(*ProjectInfoResponse), nil
	} else {
		return nil, resp.Error().(*APIError)
	}
}

// ProjectEnv request
func (sdk SDK) ProjectEnv() (*ProjectEnvResponse, error) {
	resp, err := sdk.client.R().
		SetHeader("Accept", "application/json").
		SetResult(&ProjectEnvResponse{}).
		SetError(&APIError{}).
		Get("/project")

	if err != nil {
		return nil, err
	} else if resp.IsSuccess() {
		return resp.Result().(*ProjectEnvResponse), nil
	} else {
		return nil, resp.Error().(*APIError)
	}
}

// Modules request
func (sdk SDK) Modules() (*ModuleListResponse, error) {
	resp, err := sdk.client.R().
		SetHeader("Accept", "application/json").
		SetResult(&ModuleListResponse{}).
		SetError(&APIError{}).
		Get("/module")

	if err != nil {
		return nil, err
	} else if resp.IsSuccess() {
		return resp.Result().(*ModuleListResponse), nil
	} else {
		return nil, resp.Error().(*APIError)
	}
}

// CurrentModule request
func (sdk SDK) CurrentModule() (*ModuleCurrentResponse, error) {
	resp, err := sdk.client.R().
		SetHeader("Accept", "application/json").
		SetResult(&ModuleCurrentResponse{}).
		SetError(&APIError{}).
		Get("/module/current")

	if err != nil {
		return nil, err
	} else if resp.IsSuccess() {
		return resp.Result().(*ModuleCurrentResponse), nil
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
