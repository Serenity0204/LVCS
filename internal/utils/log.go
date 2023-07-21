package utils

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type LVCSLogManager struct {
	lvcsBaseManager
}

// creates a new LVCSLogManager instance
func NewLVCSLogManager(lvcsPath string) *LVCSLogManager {
	return &LVCSLogManager{
		lvcsBaseManager: newLVCSBaseManager(lvcsPath),
	}
}

// will log out all of the version's content of the current branch
func (lvcsLogger *LVCSLogManager) Log() (string, error) {
	lvcsBranch := NewLVCSBranchManager(lvcsLogger.lvcsPath)
	curBranch, err := lvcsBranch.GetCurrentBranch()
	if err != nil {
		return "", err
	}
	curBranchPath := lvcsLogger.lvcsCommitPath + "/" + curBranch
	files, err := os.ReadDir(curBranchPath)
	if err != nil {
		return "", err
	}
	logContent := "All commits history: \n\n"
	if len(files) == 0 {
		logContent += "Empty"
		return logContent, nil
	}
	// Iterate through each file and read its content.
	for i, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(curBranchPath, file.Name())
			content, err := os.ReadFile(filePath)
			if err != nil {
				return "", err
			}
			version := strings.TrimSuffix(file.Name(), ".txt")
			if len(content) == 0 {
				logContent += strconv.Itoa(i+1) + ".\n" + "version:" + version + "\nEmpty\n\n"
				continue
			}
			logContent += strconv.Itoa(i+1) + ".\n" + "version:" + version + "\n" + string(content) + "\n\n"
		}
	}
	return logContent, nil
}

// will log out the content for that version in the current branch
func (lvcsLogger *LVCSLogManager) LogByVersion(version string) (string, error) {
	lvcsBranch := NewLVCSBranchManager(lvcsLogger.lvcsPath)
	curBranch, err := lvcsBranch.GetCurrentBranch()
	if err != nil {
		return "", err
	}
	lvcsCommitMan := NewLVCSCommitManager(lvcsLogger.lvcsPath)
	exist, err := lvcsCommitMan.versionExists(curBranch, version)
	if err != nil {
		return "", err
	}
	if !exist {
		return "", errors.New("version:" + version + " does not exist")
	}
	versionPath := lvcsLogger.lvcsCommitPath + "/" + curBranch + "/" + version + ".txt"

	content, err := os.ReadFile(versionPath)
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
