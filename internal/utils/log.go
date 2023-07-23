package utils

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const dash string = "****************************************************************************************************"

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
	lvcsCommit := NewLVCSCommitManager(lvcsLogger.lvcsPath)
	exist, err := lvcsCommit.versionExists(curBranch, version)
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
	return string(content), nil
}

// will log out the detailed content for that version in the current branch
func (lvcsLogger *LVCSLogManager) LogByVersionDetail(version string) (string, error) {
	logs, err := lvcsLogger.LogByVersion(version)
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(strings.NewReader(logs))
	lvcsFileIO := NewLVCSFileHashIOManager(lvcsLogger.lvcsPath)
	logContent := "Commit History With Detailed File View:\n\n"
	if len(logs) == 0 {
		logContent += "Empty\n"
		return logContent, nil
	}

	i := 1
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) == 3 {
			filePath := parts[0]
			oid := parts[1]
			content, err := lvcsFileIO.CatFile(oid)
			if err != nil {
				return "", err
			}
			logContent += strconv.Itoa(i) + ":" + filePath + "\n" + dash + "\n" + content + "\n" + dash + "\n\n"
			i++
		}
	}
	err = scanner.Err()
	if err != nil {
		return "", err
	}

	return logContent, nil
}
