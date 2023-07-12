package helper

import (
	"errors"
	"os"
)

func BranchExists(commitFolderPath string, branchName string) bool {
	_, err := os.Stat(commitFolderPath + "/" + branchName)
	// if err is nil then branch exists
	return err == nil
}

// need to move later
func CreateBranch(commitFolderPath string, branchName string) error {
	newBranchPath := commitFolderPath + "/" + branchName
	err := os.Mkdir(newBranchPath, 0755)
	if err != nil {
		return errors.New("failed to create branch folder:" + branchName)
	}
	return nil
}
