package internal

import (
	"errors"
	"os"

	"github.com/Serenity0204/LVCS/internal/ui"
	"github.com/Serenity0204/LVCS/internal/utils"
)

type LVCSManager struct {
	lvcsPath string
	lvcsMan  map[string]interface{}
}

// NewLVCSManager creates a new LVCSAdd instance
func NewLVCSManager(lvcsPath string) *LVCSManager {
	manager := make(map[string]interface{})
	branchMan := utils.NewLVCSBranchManager(lvcsPath)
	commitMan := utils.NewLVCSCommitManager(lvcsPath)
	fileHashIOMan := utils.NewLVCSFileHashIOManager(lvcsPath)
	initMan := utils.NewLVCSInitManager(lvcsPath)
	stageMan := utils.NewLVCSStageManager(lvcsPath)
	logMan := utils.NewLVCSLogManager(lvcsPath)

	manager["branch"] = branchMan
	manager["commit"] = commitMan
	manager["fileHashIO"] = fileHashIOMan
	manager["init"] = initMan
	manager["stage"] = stageMan
	manager["log"] = logMan
	return &LVCSManager{
		lvcsPath: lvcsPath,
		lvcsMan:  manager,
	}
}

func (lvcsManager *LVCSManager) LVCSExists() (bool, error) {
	initMan, ok := lvcsManager.lvcsMan["init"].(*utils.LVCSInitManager)
	if !ok {
		return false, errors.New("failed to check lvcs existence")
	}
	// AlreadyInit() == LVCSExists()
	return initMan.AlreadyInit(), nil
}

func (lvcsManager *LVCSManager) GetRandomASCIIArt() (string, error) {
	ascii := ui.NewASCIIArtGenerator()
	art, err := ascii.GetRandASCIIArt()
	if err != nil {
		return "", err
	}
	return art, nil
}
func (lvcsManager *LVCSManager) Dump() error {
	exist, err := lvcsManager.LVCSExists()
	if err != nil {
		return err
	}
	if !exist {
		return errors.New(".lvcs directory does not exist")
	}
	err = os.RemoveAll(lvcsManager.lvcsPath)
	if err != nil {
		return err
	}
	return nil
}

func (lvcsManager *LVCSManager) Execute(command string, subcommands []string) (string, error) {
	switch command {
	case "hashObject":
		return lvcsManager.hashObjectHandler(subcommands)
	case "catFile":
		return lvcsManager.catFileHandler(subcommands)
	case "init":
		return lvcsManager.initHandler(subcommands)
	case "commit":
		return lvcsManager.commitHandler(subcommands)
	case "branch":
		return lvcsManager.branchHandler(subcommands)
	case "stage":
		return lvcsManager.stageHandler(subcommands)
	case "log":
		return lvcsManager.logHandler(subcommands)
	default:
		// unknown command
		return "", errors.New("unknown command:" + command)
	}
}
