package cidsdk

type ModuleActionData struct {
	ProjectDir string              `json:"project-dir"`
	Config     CurrentConfig       `json:"config"`
	Env        map[string]string   `json:"env"`
	Module     ProjectModule       `json:"module"`
	Deployment *DeploymentResponse `json:"deployment"`
}

// ModuleActionDataV1 collects data for module-scoped actions
func (sdk SDK) ModuleActionDataV1() (*ModuleActionData, error) {
	resp, err := sdk.client.R().
		SetHeader("Accept", "application/json").
		SetResult(&ModuleActionData{}).
		SetError(&APIError{}).
		Get("/v1/job/module-action-data")

	if err != nil {
		return nil, err
	} else if resp.IsSuccess() {
		return resp.Result().(*ModuleActionData), nil
	} else {
		return nil, resp.Error().(*APIError)
	}
}

type ProjectActionData struct {
	ProjectDir string            `json:"project-dir"`
	Config     CurrentConfig     `json:"config"`
	Env        map[string]string `json:"env"`
	Modules    []ProjectModule   `json:"modules"`
}

// ProjectActionDataV1 collects data for project-scoped actions
func (sdk SDK) ProjectActionDataV1() (*ProjectActionData, error) {
	resp, err := sdk.client.R().
		SetHeader("Accept", "application/json").
		SetResult(&ProjectActionData{}).
		SetError(&APIError{}).
		Get("/v1/job/project-action-data")

	if err != nil {
		return nil, err
	} else if resp.IsSuccess() {
		return resp.Result().(*ProjectActionData), nil
	} else {
		return nil, resp.Error().(*APIError)
	}
}
