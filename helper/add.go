package helper

import (
	"os"
	"path/filepath"
)

// GetAbsolutePath returns the absolute path of a file given its relative path
func getAbsolutePath(relativePath string) (string, error) {
	// Get the absolute path of the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	// Get the absolute path of the file by joining the working directory and the relative path
	absolutePath := filepath.Join(wd, relativePath)
	return absolutePath, nil
}

// need to check if file exist or not, if yes then do early return
func Add(file string, lvcsPath string) error {
	oid, err := HashObject(file, lvcsPath)
	if err != nil {
		return err
	}
	relativePath := lvcsPath + "/stage.txt"
	_, err = os.Stat(relativePath)
	// if does not exist
	if err != nil {
		return err
	}

	// open the file
	stageFile, err := os.OpenFile(relativePath, os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer stageFile.Close()

	absPath, err := getAbsolutePath(file)
	if err != nil {
		return err
	}
	// Append filename, absolute file path, oid into stage.txt
	content := file + " " + absPath + " " + string(oid) + "\n"
	_, err = stageFile.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
