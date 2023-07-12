package helper

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)



func getLatestVersion(commitFolderPath string, branchName string) (int, error) {
	files, err := os.ReadDir(commitFolderPath + "/" + branchName)
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

func CreateNewCommitRecord(lvcsPath string, branchName string, version int) error {

	// get the stage content
	stagePath := lvcsPath + "/stage.txt"
	content, err := os.ReadFile(stagePath)
	if err != nil {
		return err
	}

	branchPath := lvcsPath + "/commits/" + branchName + "/v" + strconv.Itoa(version) + ".txt"

	// else create
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



func Commit(lvcsPath string, branchName string) error {
	// default is master for branchName
	commitFolderPath := lvcsPath + "/commits/"

	check := BranchExists(commitFolderPath, branchName)
	// branch does not exist
	if !check {
		return errors.New("Branch:" + branchName + " does not exist")
	}

	// get the latest version number
	version, err := getLatestVersion(commitFolderPath, branchName)
	if err != nil {
		return err
	}
	curVersion := version + 1

	// create the commit record
	err = CreateNewCommitRecord(lvcsPath, branchName, curVersion)
	if err != nil {
		return err
	}
	return nil
}
