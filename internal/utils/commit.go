package utils

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

type LVCSCommitManager struct {
	lvcsPath       string
	lvcsCommitPath string
	lvcsStagePath  string
}

// creates a new LVCSCommit instance
func NewLVCSCommitManager(lvcsPath string) *LVCSCommitManager {
	return &LVCSCommitManager{
		lvcsPath:       lvcsPath,
		lvcsCommitPath: lvcsPath + "/commits",
		lvcsStagePath:  lvcsPath + "/stage.txt",
	}
}

func (lvcsCommit *LVCSCommitManager) getLatestVersion(branchName string) (int, error) {
	files, err := os.ReadDir(lvcsCommit.lvcsCommitPath + "/" + branchName)
	if err != nil {
		return 0, err
	}

	latestVersion := -1
	for _, file := range files {
		// if it's not a file then it's error
		if file.IsDir() {
			return -1, errors.New("read A Dir Inside commits")
		}
		// it's a file
		fileName := file.Name()
		// check for error
		if !strings.HasPrefix(fileName, "v") {
			return -1, errors.New("wrong version number")
		}

		// fileName will be in form like v1.txt, v2.txt
		versionStr := strings.TrimSuffix(fileName, ".txt")
		versionStr = strings.TrimPrefix(versionStr, "v")
		version, err := strconv.Atoi(versionStr)
		if err != nil {
			// Skip invalid file names
			return -1, errors.New("error While Reading Version Numbers")
		}

		if version > latestVersion {
			latestVersion = version
		}
	}
	return latestVersion, nil
}

func (lvcsCommit *LVCSCommitManager) createNewCommitRecord(branchName string, version int) error {

	// get the stage content
	content, err := os.ReadFile(lvcsCommit.lvcsStagePath)
	if err != nil {
		return err
	}

	branchPath := lvcsCommit.lvcsCommitPath + "/" + branchName + "/v" + strconv.Itoa(version) + ".txt"

	// create
	versionFile, err := os.Create(branchPath)
	if err != nil {
		return err
	}
	defer versionFile.Close()

	// Write the content into the file
	_, err = io.Copy(versionFile, bytes.NewReader(content))
	if err != nil {
		// Clean up the file if writing fails
		os.Remove(branchPath)
		return err
	}
	return nil
}

func (lvcsCommit *LVCSCommitManager) Commit(branchName string) error {
	// default is master for branchName
	lvcsBranch := NewLVCSBranchManager(lvcsCommit.lvcsPath)
	check := lvcsBranch.BranchExists(branchName)
	// branch does not exist
	if !check {
		return errors.New("Branch:" + branchName + " does not exist")
	}

	// get the latest version number
	version, err := lvcsCommit.getLatestVersion(branchName)
	if err != nil {
		return err
	}
	curVersion := version + 1

	// create the commit record
	err = lvcsCommit.createNewCommitRecord(branchName, curVersion)
	if err != nil {
		return err
	}
	return nil
}

// need to move this later too
func RemoveStageContent(lvcsPath string) error {
	stagePath := lvcsPath + "/stage.txt"
	stageFile, err := os.OpenFile(stagePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer stageFile.Close()
	return nil
}
