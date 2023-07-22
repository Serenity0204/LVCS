package internal

import (
	"errors"
	"strconv"

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

func (lvcsManager *LVCSManager) Execute(command string, subcommands []string) (string, error) {
	switch command {
	case "hashObject":
		// 1 subcommand
		// return oid if no error
		fileHashIOMan, ok := lvcsManager.lvcsMan["fileHashIO"].(*utils.LVCSFileHashIOManager)
		if !ok {
			return "", errors.New("failed to execute hash-object")
		}
		if len(subcommands) != 1 {
			return "", errors.New("number of argumment not correct, expected a path to a file but not found")
		}

		oid, err := fileHashIOMan.HashObject(subcommands[0])
		if err != nil {
			return "", err
		}
		return oid, nil
	case "catFile":
		// 1 subcommand
		// return file content if no error
		fileHashIOMan, ok := lvcsManager.lvcsMan["fileHashIO"].(*utils.LVCSFileHashIOManager)
		if !ok {
			return "", errors.New("failed to execute cat-file")
		}
		if len(subcommands) != 1 {
			return "", errors.New("number of argumment not correct, expected an oid but not found")
		}
		content, err := fileHashIOMan.CatFile(subcommands[0])
		if err != nil {
			return "", err
		}
		return content, nil
	case "init":
		// 0 subcommand
		if len(subcommands) != 0 {
			return "", errors.New("too many args for init")
		}
		// return success message if no error
		initMan, ok := lvcsManager.lvcsMan["init"].(*utils.LVCSInitManager)
		if !ok {
			return "", errors.New("failed to execute init")
		}
		err := initMan.Init()
		if err != nil {
			return "", err
		}
		return string("Init .lvcs directory at:" + lvcsManager.lvcsPath + " success"), nil
	case "commit":
		// 0, 1, 2, or 3 subcommands
		// commit
		commitMan, ok := lvcsManager.lvcsMan["commit"].(*utils.LVCSCommitManager)
		if !ok {
			return "", errors.New("failed to execute commit")
		}
		if len(subcommands) == 0 {
			err := commitMan.Commit(false)
			if err != nil {
				return "", err
			}
			return string("commit success"), nil
		}

		if len(subcommands) == 1 {
			// inherit commit
			if subcommands[0] == "inherit" {
				err := commitMan.Commit(true)
				if err != nil {
					return "", err
				}
				return string("commit success"), nil
			}
			// commit latest
			if subcommands[0] == "latest" {
				latest, err := commitMan.GetLatestVersion()
				if err != nil {
					return "", err
				}
				return latest, nil
			}
			// commit current
			if subcommands[0] == "current" {
				current, err := commitMan.GetCurrentVersion()
				if err != nil {
					return "", err
				}
				return current, nil
			}
			// commit tree
			if subcommands[0] == "tree" {
				tree, err := commitMan.CommitTree()
				if err != nil {
					return "", err
				}
				return tree, nil
			}
			return "", errors.New("unknown subcommands:" + subcommands[0])
		}
		// commit switch <version number>
		if len(subcommands) == 2 && subcommands[0] == "switch" {
			err := commitMan.SwitchCommitVersion(subcommands[1])
			if err != nil {
				return "", err
			}
			return string("Switch to " + subcommands[1] + " success"), nil
		}
		if len(subcommands) == 3 && subcommands[0] == "lca" {
			version1 := subcommands[1]
			version2 := subcommands[2]
			lca, err := commitMan.LCA(version1, version2)
			if err != nil {
				return "", err
			}
			return lca, nil
		}
		return "", errors.New("invalid:number of arguments")
	case "branch":
		// 0, or 2 sub commands
		branchMan, ok := lvcsManager.lvcsMan["branch"].(*utils.LVCSBranchManager)
		if !ok {
			return "", errors.New("failed to execute branch")
		}

		// list all the branches
		if len(subcommands) == 0 {
			branches, err := branchMan.GetAllBranch()
			if err != nil {
				return "", errors.New("failed to retrieve all of the branches")
			}
			allBranchNames := "All branches:\n"
			for i, branchName := range branches {
				allBranchNames += strconv.Itoa(i+1) + ":" + branchName + "\n"
			}
			return allBranchNames, nil
		}
		if len(subcommands) == 1 {
			// get current branch
			if subcommands[0] == "current" {
				curBranch, err := branchMan.GetCurrentBranch()
				if err != nil {
					return "", err
				}
				return string("Current branch is:" + curBranch), nil
			}
			return "", errors.New("unknown subcommands:" + subcommands[0])
		}
		if len(subcommands) == 2 {
			branchName := subcommands[1]
			exists := branchMan.BranchExists(branchName)
			// check if branch exists
			if subcommands[0] == "exists" {
				return strconv.FormatBool(exists), nil
			}
			// create branch
			if subcommands[0] == "create" {
				if exists {
					return "", errors.New("branch:" + branchName + " already exists")
				}
				err := branchMan.CreateBranch(branchName)
				if err != nil {
					return "", err
				}
				return string("Create branch:" + branchName + " success"), nil
			}
			// checkout branch
			if subcommands[0] == "checkout" {
				if !exists {
					return "", errors.New("branch:" + branchName + " does not exist")
				}
				err := branchMan.CheckoutBranch(branchName)
				if err != nil {
					return "", err
				}
				return string("Checkout branch:" + branchName + " success"), nil
			}
			// delete branch
			if subcommands[0] == "delete" {
				// cannot delete branch that DNE
				if !exists {
					return "", errors.New("branch:" + branchName + " does not exist")
				}
				// cannot delete current branch
				curBranch, err := branchMan.GetCurrentBranch()
				if err != nil {
					return "", err
				}
				if branchName == curBranch {
					return "", errors.New("cannot delete current working branch:" + branchName)
				}
				err = branchMan.DeleteBranch(branchName)
				if err != nil {
					return "", err
				}
				return string("Delete branch:" + branchName + " success"), nil
			}
			// if 2 args and it's not one of the above then it's an error
			return "", errors.New("unknown subcommands:" + subcommands[0])
		}
		return "", errors.New("invalid:number of arguments")
	case "stage":
		// 0, 1, or 2 arguments
		stageMan, ok := lvcsManager.lvcsMan["stage"].(*utils.LVCSStageManager)
		if !ok {
			return "", errors.New("failed to execute stage")
		}
		// Show staging content
		if len(subcommands) == 0 {
			content, err := stageMan.GetStageContent()
			if err != nil {
				return "", err
			}
			return content, nil
		}
		// Untrack All
		if len(subcommands) == 1 && subcommands[0] == "untrack" {
			err := stageMan.RemoveAllStageContent()
			if err != nil {
				return "", err
			}
			return string("Untracked all staged files success"), nil
		}

		// stage add, need to have add as subcommand, and at least another arg
		if len(subcommands) >= 2 {
			// Untrack file
			if len(subcommands) == 2 && subcommands[0] == "untrack" {
				err := stageMan.RemoveStageContent(subcommands[1])
				if err != nil {
					return "", err
				}
				return string("Untracked " + subcommands[1] + " success"), nil
			}
			if subcommands[0] == "add" {
				for i, file := range subcommands {
					if i == 0 {
						continue
					}
					err := stageMan.Add(file)
					if err != nil {
						return "", errors.New("failed to add:" + file)
					}
				}
				trackedFiles := "Added files:\n"
				for i, fileName := range subcommands {
					if i == 0 {
						continue
					}
					trackedFiles += strconv.Itoa(i) + ":" + fileName + "\n"
				}
				return trackedFiles, nil
			}
			return "", errors.New("unknown subcommands:" + subcommands[0])
		}
		return "", errors.New("invalid:number of arguments")
	case "log":
		// 0 or 1 args
		logMan, ok := lvcsManager.lvcsMan["log"].(*utils.LVCSLogManager)
		if !ok {
			return "", errors.New("failed to execute log")
		}
		if len(subcommands) == 0 {
			logHistory, err := logMan.Log()
			if err != nil {
				return "", err
			}
			return logHistory, nil
		}
		if len(subcommands) == 1 {
			version := subcommands[0]
			content, err := logMan.LogByVersion(version)
			if err != nil {
				return "", err
			}
			// if empty log empty
			if len(content) == 0 {
				logContent := "version:" + version + "\nEmpty\n\n"
				return logContent, nil
			}
			logContent := "version:" + version + "\n" + string(content) + "\n\n"
			return logContent, nil
		}
		return "", errors.New("invalid:number of arguments")
	default:
		// unknown command
		return "", errors.New("unknown command:" + command)
	}
}
