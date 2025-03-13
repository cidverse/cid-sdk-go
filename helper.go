package cidsdk

import (
	"encoding/json"
)

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

		// overwrite from env
		PopulateFromEnv(cfg, nil)
	}

	return ModuleActionData{ProjectDir: config.ProjectDir, Config: *config, Module: *module, Env: env}, nil
}

// ProjectAction is a utility function to prepare to run a project-scoped action
func (sdk SDK) ProjectAction(cfg any) (ProjectActionData, error) {
	config, err := sdk.CurrentConfig()
	if err != nil {
		return ProjectActionData{}, err
	}

	env, err := sdk.Env()
	if err != nil {
		return ProjectActionData{}, err
	}

	if config.Config != "" && cfg != nil {
		err := json.Unmarshal([]byte(config.Config), cfg)
		if err != nil {
			return ProjectActionData{}, err
		}

		// overwrite from env
		PopulateFromEnv(cfg, nil)
	}

	modules, err := sdk.Modules()
	if err != nil {
		return ProjectActionData{}, err
	}

	return ProjectActionData{ProjectDir: config.ProjectDir, Config: *config, Modules: *modules, Env: env}, nil
}
