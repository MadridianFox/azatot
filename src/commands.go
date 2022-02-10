package src

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func checkAndLoadHC(homeConfigPath string) (*HomeConfig, error) {
	err := CheckHomeConfigIsEmpty(homeConfigPath)
	if err != nil {
		return nil, err
	}
	hc, err := LoadHomeConfig(homeConfigPath)
	if err != nil {
		return nil, err
	}

	return hc, nil
}

func getWorkspaceConfig(homeConfigPath string) (*MainConfig, error) {
	hc, err := checkAndLoadHC(homeConfigPath)
	if err != nil {
		return nil, err
	}

	wsPath, err := hc.GetCurrentWsPath()
	if err != nil {
		return nil, err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	cfg := NewConfig(wsPath, cwd)
	err = cfg.LoadFromFile()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func addStartFlags(fs *flag.FlagSet, params *SvcStartParams) {
	fs.StringVar(&params.Mode, "mode", "default", "tag for dependencies selecting")
	fs.BoolVar(&params.Force, "force", false, "force start dependencies")
}

func addComposeFlags(fs *flag.FlagSet, params *SvcComposeParams) {
	fs.StringVar(&params.SvcName, "svc", "", "name of service")
}

func addExecFlags(fs *flag.FlagSet, params *SvcExecParams) {
	fs.IntVar(&params.UID, "uid", os.Getuid(), "user id")
}

func NeedHelp(args []string, usage string, lines []string) bool {
	if len(args) > 0 && (args[0] == "-h" || args[0] == "--help" || args[0] == "help") {
		fmt.Printf("Usage: %s %s\n", os.Args[0], usage)
		if lines != nil {
			fmt.Println("")
			fmt.Println(strings.Join(lines, "\n"))
			fmt.Println("")
		}
		return true
	}
	return false
}

func CmdWorkspaceList(homeConfigPath string, args []string) error {
	if NeedHelp(args, "workspace list", []string{
		"Show list of registered workspaces.",
	}) {
		return nil
	}
	hc, err := checkAndLoadHC(homeConfigPath)
	if err != nil {
		return err
	}

	for _, workspace := range hc.Workspaces {
		fmt.Printf("%-10s %s\n", workspace.Name, workspace.Path)
	}

	return nil
}

func CmdWorkspaceAdd(homeConfigPath string, args []string) error {
	if NeedHelp(args, "workspace add NAME PATH", []string{
		"Register new workspace.",
	}) {
		return nil
	}
	hc, err := checkAndLoadHC(homeConfigPath)
	if err != nil {
		return err
	}

	if len(args) != 2 {
		return errors.New("command requires exactly 2 arguments")
	}

	name := args[0]
	wsPath := args[1]

	ws := hc.findWorkspace(name)
	if ws != nil {
		return errors.New(fmt.Sprintf("workspace with name '%s' already exists", name))
	}

	err = hc.AddWorkspace(name, wsPath)
	if err != nil {
		return err
	}

	fmt.Printf("workspace '%s' is added\n", name)
	return nil
}

func CmdWorkspaceSelect(homeConfigPath string, args []string) error {
	if NeedHelp(args, "workspace select NAME", []string{
		"Set workspace with name NAME as current.",
	}) {
		return nil
	}
	hc, err := checkAndLoadHC(homeConfigPath)
	if err != nil {
		return err
	}

	if len(args) != 1 {
		return errors.New("command requires exactly 1 argument")
	}

	name := args[0]

	ws := hc.findWorkspace(name)
	if ws == nil {
		return errors.New(fmt.Sprintf("workspace with name '%s' is not defined", name))
	}

	hc.CurrentWorkspace = name
	err = SaveHomeConfig(hc)
	if err != nil {
		return err
	}

	fmt.Printf("active workspace changed to '%s'\n", name)
	return nil
}

func CmdWorkspaceShow(homeConfigPath string, args []string) error {
	if NeedHelp(args, "workspace show", []string{
		"Print current workspace name.",
	}) {
		return nil
	}
	hc, err := checkAndLoadHC(homeConfigPath)
	if err != nil {
		return err
	}
	fmt.Println(hc.CurrentWorkspace)

	return nil
}

func CmdWorkspaceHelp() error {
	NeedHelp([]string{"--help"}, "workspace COMMAND", []string{
		"Available commands:",
		"  ls, list - list available workspaces",
		"  show     - show current workspace name",
		"  add      - add new workspace",
		"  select   - select workspace as current",
	})
	return nil
}

func CmdServiceStart(homeConfigPath string, args []string) error {
	if NeedHelp(args, "start [OPTIONS] [NAMES...]", []string{
		"Start one or more services.",
		"By default starts service found with current directory, but you can pass one or more service names instead.",
		"",
		"Available options:",
		fmt.Sprintf("  %-10s - %s", "--force", "force start dependencies, even if service already started"),
		fmt.Sprintf("  %-10s - %s", "--mode", "start only dependencies with specified mode, by default starts 'default' dependencies"),
	}) {
		return nil
	}
	fs := flag.NewFlagSet("start", flag.ContinueOnError)
	startParams := &SvcStartParams{}
	addStartFlags(fs, startParams)
	err := fs.Parse(args)
	if err != nil {
		return err
	}

	cfg, err := getWorkspaceConfig(homeConfigPath)
	if err != nil {
		return err
	}

	svcNames := fs.Args()
	if len(svcNames) > 0 {
		for _, svcName := range svcNames {
			svc, err := CreateFromSvcName(cfg, svcName)
			if err != nil {
				return err
			}

			err = svc.Start(startParams)
			if err != nil {
				return err
			}
		}
	} else {
		svcName, err := cfg.FindServiceByPath()
		if err != nil {
			return err
		}

		svc, err := CreateFromSvcName(cfg, svcName)
		if err != nil {
			return err
		}

		err = svc.Start(startParams)
		if err != nil {
			return err
		}
	}

	return nil
}

func CmdServiceStop(homeConfigPath string, args []string) error {
	if NeedHelp(args, "stop [NAMES...]", []string{
		"Stop one or more services.",
		"By default stops service found with current directory, but you can pass one or more service names instead.",
	}) {
		return nil
	}
	cfg, err := getWorkspaceConfig(homeConfigPath)
	if err != nil {
		return err
	}

	svcNames := args
	if len(svcNames) > 0 {
		for _, svcName := range svcNames {
			svc, err := CreateFromSvcName(cfg, svcName)
			if err != nil {
				return err
			}

			err = svc.Stop()
			if err != nil {
				return err
			}
		}
	} else {
		svcName, err := cfg.FindServiceByPath()
		if err != nil {
			return err
		}

		svc, err := CreateFromSvcName(cfg, svcName)
		if err != nil {
			return err
		}

		err = svc.Stop()
		if err != nil {
			return err
		}
	}

	return nil
}

func CmdServiceDestroy(homeConfigPath string, args []string) error {
	if NeedHelp(args, "destroy [NAMES...]", []string{
		"Stop and remove containers of one or more services.",
		"By default destroys service found with current directory, but you can pass one or more service names instead.",
	}) {
		return nil
	}
	cfg, err := getWorkspaceConfig(homeConfigPath)
	if err != nil {
		return err
	}

	svcNames := args
	if len(svcNames) > 0 {
		for _, svcName := range svcNames {
			svc, err := CreateFromSvcName(cfg, svcName)
			if err != nil {
				return err
			}

			err = svc.Destroy()
			if err != nil {
				return err
			}
		}
	} else {
		svcName, err := cfg.FindServiceByPath()
		if err != nil {
			return err
		}

		svc, err := CreateFromSvcName(cfg, svcName)
		if err != nil {
			return err
		}

		err = svc.Destroy()
		if err != nil {
			return err
		}
	}

	return nil
}

func CmdServiceRestart(homeConfigPath string, args []string) error {
	if NeedHelp(args, "restart [OPTIONS] [NAMES...]", []string{
		"Restart one or more services.",
		"By default restart service found with current directory, but you can pass one or more service names instead.",
		"",
		"Available options:",
		fmt.Sprintf("  %-10s - %s", "--hard", "destroy service instead of stopping it"),
	}) {
		return nil
	}
	fs := flag.NewFlagSet("restart", flag.ContinueOnError)
	restartParams := &SvcRestartParams{}
	fs.BoolVar(&restartParams.Hard, "hard", false, "destroy container instead of stop it before start")
	err := fs.Parse(args)
	if err != nil {
		return err
	}

	cfg, err := getWorkspaceConfig(homeConfigPath)
	if err != nil {
		return err
	}

	svcNames := fs.Args()
	if len(svcNames) > 0 {
		for _, svcName := range svcNames {
			svc, err := CreateFromSvcName(cfg, svcName)
			if err != nil {
				return err
			}

			err = svc.Restart(restartParams)
			if err != nil {
				return err
			}
		}
	} else {
		svcName, err := cfg.FindServiceByPath()
		if err != nil {
			return err
		}

		svc, err := CreateFromSvcName(cfg, svcName)
		if err != nil {
			return err
		}

		err = svc.Restart(restartParams)
		if err != nil {
			return err
		}
	}

	return nil
}

func CmdServiceVars(homeConfigPath string, args []string) error {
	if NeedHelp(args, "vars [NAME]", []string{
		"Print all variables computed for service.",
		"By default uses service found with current directory, but you can pass name of another service instead.",
	}) {
		return nil
	}
	cfg, err := getWorkspaceConfig(homeConfigPath)
	if err != nil {
		return err
	}

	var svcName string

	if len(args) > 0 {
		svcName = args[0]
	} else {
		svcName, err = cfg.FindServiceByPath()
		if err != nil {
			return err
		}
	}

	svc, err := CreateFromSvcName(cfg, svcName)
	if err != nil {
		return err
	}

	err = svc.DumpVars()
	if err != nil {
		return err
	}

	return nil
}

func CmdServiceCompose(homeConfigPath string, args []string) (int, error) {
	if NeedHelp(args, "compose [OPTIONS] COMMAND [ARGS]", []string{
		"Run docker-compose command.",
		"By default uses service found with current directory.",
		"",
		"Available options:",
		fmt.Sprintf("  %-10s - %s", "--svc", "name of another service instead of current"),
	}) {
		return 0, nil
	}
	fs := flag.NewFlagSet("compose", flag.ContinueOnError)
	composeParams := &SvcComposeParams{}
	addComposeFlags(fs, composeParams)
	err := fs.Parse(args)
	if err != nil {
		return 0, err
	}

	composeParams.Cmd = fs.Args()

	cfg, err := getWorkspaceConfig(homeConfigPath)
	if err != nil {
		return 0, err
	}

	if composeParams.SvcName == "" {
		composeParams.SvcName, err = cfg.FindServiceByPath()
		if err != nil {
			return 0, err
		}
	}

	svc, err := CreateFromSvcName(cfg, composeParams.SvcName)
	if err != nil {
		return 0, err
	}

	returnCode, err := svc.Compose(composeParams)
	if err != nil {
		return 0, err
	}

	return returnCode, nil
}

func CmdServiceExec(homeConfigPath string, args []string) (int, error) {
	if NeedHelp(args, "[OPTIONS] COMMAND [ARGS]", []string{
		"Execute command in container. For module uses container of linked service.",
		"By default uses service/module found with current directory. Starts service if it is not running.",
		"",
		"Available options:",
		fmt.Sprintf("  %-10s - %s", "--force", "force start dependencies, even if service already started"),
		fmt.Sprintf("  %-10s - %s", "--svc=NAME", "name of another service or module instead of current"),
		fmt.Sprintf("  %-10s - %s", "--mode=MODE", "start only dependencies wit specified tag, by default starts 'default' dependencies"),
		fmt.Sprintf("  %-10s - %s", "--uid=UID", "use another uid, by default uses uid of current user"),
	}) {
		return 0, nil
	}
	fs := flag.NewFlagSet("exec", flag.ContinueOnError)
	execParams := &SvcExecParams{}
	addComposeFlags(fs, &execParams.SvcComposeParams)
	addStartFlags(fs, &execParams.SvcStartParams)
	addExecFlags(fs, execParams)
	err := fs.Parse(args)
	if err != nil {
		return 0, err
	}

	execParams.Cmd = fs.Args()

	cfg, err := getWorkspaceConfig(homeConfigPath)
	if err != nil {
		return 0, err
	}

	var mdl *ModuleConfig

	if execParams.SvcName == "" {
		mdl, err = cfg.FindModuleByPath()
		if err == nil {
			execParams.SvcName = mdl.HostedIn
		} else {
			execParams.SvcName, err = cfg.FindServiceByPath()
			if err != nil {
				return 0, err
			}
		}
	} else {
		mdl, err := cfg.FindModuleByName(execParams.SvcName)
		if err == nil {
			execParams.SvcName = mdl.HostedIn
		}
	}

	if mdl != nil {
		execParams.WorkingDir, err = cfg.renderPath(mdl.ExecPath)
		if err != nil {
			return 0, err
		}
	}

	svc, err := CreateFromSvcName(cfg, execParams.SvcName)
	if err != nil {
		return 0, err
	}

	returnCode, err := svc.Exec(execParams)
	if err != nil {
		return 0, err
	}

	return returnCode, nil
}

func CmdServiceSetHooks(args []string) error {
	if NeedHelp(args, "set-hooks HOOKS_PATH", []string{
		"Install hooks from specified folder to .git/hooks.",
		"HOOKS_PATH must contain subdirectories with names as git hooks, eg. 'pre-commit'.",
		"One subdirectory can contain one or many scripts with .sh extension.",
		"Every script wil be wrapped with 'elc --tag=hook' command.",
	}) {
		return nil
	}
	if len(args) != 1 {
		return errors.New("command requires exactly 1 argument")
	}
	hooksFolder := args[0]
	err := SetGitHooks(hooksFolder, os.Args[0])
	if err != nil {
		return err
	}

	return nil
}
