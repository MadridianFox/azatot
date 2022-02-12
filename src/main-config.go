package src

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type CoreConfig struct {
	Aliases   map[string]string         `yaml:"aliases"`
	Templates map[string]TemplateConfig `yaml:"templates"`
	Services  map[string]ServiceConfig  `yaml:"services"`
	Modules   map[string]ModuleConfig   `yaml:"modules"`
	Variables yaml.MapSlice             `yaml:"variables"`
}

type MainConfig struct {
	CoreConfig    `yaml:",inline"`
	Name          string     `yaml:"name"`
	LocalConfig   CoreConfig `yaml:"-"`
	WorkspacePath string     `yaml:"-"`
	Cwd           string     `yaml:"-"`
	WillStart     []string   `yaml:"-"`
}

func NewConfig(workspacePath string, cwd string) *MainConfig {
	cfg := MainConfig{
		WorkspacePath: workspacePath,
		Cwd:           cwd,
	}

	return &cfg
}

func (cfg *MainConfig) LoadFromFile() error {
	yamlFile, err := ioutil.ReadFile(path.Join(cfg.WorkspacePath, "workspace.yaml"))
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		return err
	}

	_, err = os.Stat(path.Join(cfg.WorkspacePath, "env.yaml"))
	if err == nil {
		yamlFile, err = ioutil.ReadFile(path.Join(cfg.WorkspacePath, "env.yaml"))
		if err != nil {
			return err
		}

		err = yaml.Unmarshal(yamlFile, &cfg.LocalConfig)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cfg *MainConfig) makeGlobalEnv() (Context, error) {
	ctx := make(Context, 0)

	ctx = ctx.add("WORKSPACE_PATH", strings.TrimRight(cfg.WorkspacePath, "/"))
	ctx = ctx.add("WORKSPACE_NAME", cfg.Name)

	for _, pair := range cfg.LocalConfig.Variables {
		value, err := substVars(pair.Value.(string), ctx)
		if err != nil {
			return nil, err
		}
		ctx = ctx.add(pair.Key.(string), value)
	}

	for _, pair := range cfg.Variables {
		value, err := substVars(pair.Value.(string), ctx)
		if err != nil {
			return nil, err
		}
		ctx = ctx.add(pair.Key.(string), value)
	}

	return ctx, nil
}

func (cfg *MainConfig) renderPath(path string) (string, error) {
	env, err := cfg.makeGlobalEnv()
	if err != nil {
		return "", err
	}
	return substVars(path, env)
}

func (cfg *MainConfig) FindServiceByPath() (string, error) {
	for name, svc := range cfg.Services {
		svcPath, err := cfg.renderPath(svc.Path)
		if err != nil {
			return "", err
		}
		if strings.HasPrefix(cfg.Cwd, svcPath) {
			return name, nil
		}
	}

	return "", errors.New("you are not in service folder")
}

func (cfg *MainConfig) FindServiceByName(name string) (*ServiceConfig, error) {
	realName := cfg.LocalConfig.resolveAlias(name)
	svc, found := cfg.Services[realName]
	if !found {
		return nil, errors.New(fmt.Sprintf("service %s not found", name))
	}

	return &svc, nil
}

func (cfg *MainConfig) FindTemplateByName(name string) (*TemplateConfig, error) {
	tpl, found := cfg.Templates[name]
	if !found {
		return nil, errors.New(fmt.Sprintf("template %s not found", name))
	}

	return &tpl, nil
}

func (cfg *MainConfig) FindModuleByName(name string) (*ModuleConfig, error) {
	realName := cfg.LocalConfig.resolveAlias(name)
	mdl, found := cfg.Modules[realName]
	if !found {
		return nil, errors.New(fmt.Sprintf("module %s not found", name))
	}

	return &mdl, nil
}

func (cfg *MainConfig) FindModuleByPath() (*ModuleConfig, error) {
	for _, mdl := range cfg.Modules {
		mdlPath, err := cfg.renderPath(mdl.Path)
		if err != nil {
			return nil, err
		}
		if strings.HasPrefix(cfg.Cwd, mdlPath) {
			return &mdl, nil
		}
	}

	return nil, errors.New("you are not in module folder")
}

func (cfg *MainConfig) GetAllSvcNames() []string {
	result := make([]string, 0)
	for name := range cfg.Services {
		result = append(result, name)
	}

	return result
}

func (ccfg *CoreConfig) resolveAlias(name string) string {
	realName, found := ccfg.Aliases[name]
	if found {
		return realName
	}

	return name
}
