package cidsdk

import (
	"encoding/json"
)

type ModuleActionData struct {
	ProjectDir string
	Module     ProjectModule
	Config     CurrentConfig
	Env        map[string]string
}

type ProjectActionData struct {
	ProjectDir string
	Config     CurrentConfig
	Env        map[string]string
}

// ModuleAction is a utility function to prepare to run a module-scoped action
func (sdk SDK) ModuleAction(cfg any) (ModuleActionData, error) {
	config, err := sdk.CurrentConfig()
	if err != nil {
		return ModuleActionData{}, err
	}

	module, err := sdk.CurrentModule()
	if err != nil {
		return ModuleActionData{}, err
	}

	env, err := sdk.Env()
	if err != nil {
		return ModuleActionData{}, err
	}

	if config.Config != "" && cfg != nil {
		err := json.Unmarshal([]byte(config.Config), cfg)
		if err != nil {
			return ModuleActionData{}, err
		}

		OverwriteFromEnv(&cfg)
	}

	return ModuleActionData{ProjectDir: (*config).ProjectDir, Config: *config, Module: *module, Env: env}, nil
}

// ProjectAction is a utility function to prepare to run a project-scoped action
func (sdk SDK) ProjectAction(cfg any) (ProjectActionData, error) {
	config, err := sdk.CurrentConfig()
	if err != nil {
		return ProjectActionData{}, err
	}

	if config.Config != "" && cfg != nil {
		err := json.Unmarshal([]byte(config.Config), cfg)
		if err != nil {
			return ProjectActionData{}, err
		}

		OverwriteFromEnv(&cfg)
	}

	return ProjectActionData{ProjectDir: (*config).ProjectDir, Config: *config}, nil
}
