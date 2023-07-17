package internal

import (
	"errors"
	"strconv"

	"github.com/Serenity0204/LVCS/internal/utils"
)

type LVCSManager struct {
	lvcsPath string
	lvcsMan  map[string]interface{}
}

// NewLVCSManager creates a new LVCSAdd instance
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
	case "add":
		// at least 1 subcommand
		// return success messaage if no error
		addMan, ok := lvcsManager.lvcsMan["add"].(*utils.LVCSAddManager) // change the receiver later
		if !ok {
			return "", errors.New("failed to execute add")
		}
		if len(subcommands) < 1 {
			return "", errors.New("number of argumment not correct, expected at least one path input but not found")
		}
		for _, file := range subcommands {
			err := addMan.Add(file)
			if err != nil {
				return "", errors.New("failed to add:" + file)
			}
		}
		trackedFiles := "Added files:\n"
		for i, fileName := range subcommands {
			trackedFiles += strconv.Itoa(i+1) + ":" + fileName + "\n"
		}
		return trackedFiles, nil
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
		// TBD
		return "", nil
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
			return "", errors.New("invalid:number of arguments")
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
	default:
		// unknown command
		return "", errors.New("unknown command:" + command)
	}
}
