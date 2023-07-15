package helper

import (
	"errors"
	"os"
)

type LVCSBranchManager struct {
	lvcsPath       string
	lvcsCommitPath string
	lvcsStagePath  string
}

// NewLVCSAdd creates a new LVCSCommit instance
func NewLVCSBranchManager(lvcsPath string) *LVCSBranchManager {
	return &LVCSBranchManager{
		lvcsPath:       lvcsPath,
		lvcsCommitPath: lvcsPath + "/commits",
		lvcsStagePath:  lvcsPath + "/stage.txt",
	}
}

func (lvcsBranch *LVCSBranchManager) BranchExists(branchName string) bool {
	_, err := os.Stat(lvcsBranch.lvcsCommitPath + "/" + branchName)
	// if err is nil then branch exists
	return err == nil
}

func (lvcsBranch *LVCSBranchManager) CreateBranch(branchName string) error {
	newBranchPath := lvcsBranch.lvcsCommitPath + "/" + branchName
	err := os.Mkdir(newBranchPath, 0755)
	if err != nil {
		return errors.New("failed to create branch folder:" + branchName)
	}
	return nil
}
