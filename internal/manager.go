package internal

import (
	"errors"

	"github.com/Serenity0204/LVCS/internal/utils"
)

type LVCSManager struct {
	lvcsMan map[string]interface{}
}

// NewLVCSAddManager creates a new LVCSAdd instance
func NewLVCSManager(lvcsPath string) *LVCSManager {
	manager := make(map[string]interface{})
	addMan := utils.NewLVCSAddManager(lvcsPath)
	branchMan := utils.NewLVCSBranchManager(lvcsPath)
	commitMan := utils.NewLVCSCommitManager(lvcsPath)
	fileHashIOMan := utils.NewLVCSFileHashIOManager(lvcsPath)
	initMan := utils.NewLVCSInitManager(lvcsPath)
	manager["add"] = addMan
	manager["branch"] = branchMan
	manager["commit"] = commitMan
	manager["fileHashIO"] = fileHashIOMan
	manager["init"] = initMan
	return &LVCSManager{
		lvcsMan: manager,
	}
}

func (lvcsManager *LVCSManager) Execute(command string, subcommands []string) error {
	switch command {
	case "hash-object":
		// TBD
	case "cat-file":
		// TBD
	case "add":
		// TBD
	case "init":
		initMan, ok := lvcsManager.lvcsMan["init"].(*utils.LVCSInitManager)
		if !ok {
			return errors.New("Failed to execute init")
		}
		err := initMan.Init()
		if err != nil {
			return err
		}
	case "commit":

	default:
		// TBD
	}

	return nil
}
