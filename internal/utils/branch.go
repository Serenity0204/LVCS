package utils

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

type LVCSBranchManager struct {
	lvcsBaseManager
}

// creates a new LVCSCommit instance
func NewLVCSBranchManager(lvcsPath string) *LVCSBranchManager {
	return &LVCSBranchManager{
		lvcsBaseManager: newLVCSBaseManager(lvcsPath),
	}
}

func (lvcsBranch *LVCSBranchManager) BranchExists(branchName string) bool {
	_, err := os.Stat(lvcsBranch.lvcsCommitPath + "/" + branchName)
	// if err is nil then branch exists
	return err == nil
}

func (lvcsBranch *LVCSBranchManager) CreateBranch(branchName string) error {
	newBranchPath := lvcsBranch.lvcsCommitPath + "/" + branchName
	newTreePath := lvcsBranch.lvcsTreePath + "/" + branchName + "_tree.txt"
	err := os.Mkdir(newBranchPath, 0755)
	if err != nil {
		return errors.New("failed to create branch folder:" + branchName)
	}
	file, err := os.Create(newTreePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func (lvcsBranch *LVCSBranchManager) DeleteBranch(branchName string) error {
	err := os.RemoveAll(lvcsBranch.lvcsCommitPath + "/" + branchName)
	if err != nil {
		return err
	}
	err = os.Remove(lvcsBranch.lvcsTreePath + "/" + branchName + "_tree.txt")
	if err != nil {
		return err
	}
	return nil
}

func (lvcsBranch *LVCSBranchManager) CheckoutBranch(branchName string) error {
	file, err := os.OpenFile(lvcsBranch.lvcsCurrentRefPath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// get latest version
	lvcsCommit := NewLVCSCommitManager(lvcsBranch.lvcsPath)
	version, err := lvcsCommit.getLatestVersion(branchName)
	if err != nil {
		return err
	}
	versionStr := ""
	if version != -1 {
		versionStr += "v" + strconv.Itoa(version)
	} else {
		versionStr += "HEAD"
	}

	_, err = file.WriteString(branchName + "\n" + versionStr + "\n")
	if err != nil {
		return err
	}
	return nil
}

func (lvcsBranch *LVCSBranchManager) GetCurrentBranch() (string, error) {
	content, err := os.ReadFile(lvcsBranch.lvcsCurrentRefPath)
	if err != nil {
		return "", err
	}
	currentBranch := strings.Split(string(content), "\n")[0]
	return currentBranch, nil
}

func (lvcsBranch *LVCSBranchManager) GetAllBranch() ([]string, error) {
	branches := []string{}
	dirEntries, err := os.ReadDir(lvcsBranch.lvcsCommitPath)
	if err != nil {
		return []string{}, err
	}
	for _, dirEntry := range dirEntries {
		if !dirEntry.IsDir() {
			continue
		}
		branches = append(branches, dirEntry.Name())
	}
	return branches, nil
}
