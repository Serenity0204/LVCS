package helper

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

func keyExists(target string, lvcsPath string) (bool, error) {
	// Open the file
	file, err := os.Open(lvcsPath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Iterate over each line
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line by white space
		fields := strings.Fields(line)
		// Invalid line format, skip to the next line
		if len(fields) != 3 {
			continue
		}

		key := fields[0]
		// Check if the key matches the target
		if key == target {
			return true, nil
		}
	}
	// Check if there was any error during scanning
	if err := scanner.Err(); err != nil {
		return false, err
	}
	// Target not found
	return false, nil
}

// need to check if file exist or not, if yes then do early return
func Add(file string, lvcsPath string) error {

	absPath, err := getAbsolutePath(file)
	if err != nil {
		return err
	}
	relativePath := lvcsPath + "/stage.txt"

	isIn, err := keyExists(absPath, relativePath)
	if err != nil {
		
		return err
	}

	// if it already exists, don't re-add.
	if isIn {
		fmt.Println("WTF")
		return nil
	}

	oid, err := HashObject(file, lvcsPath)
	if err != nil {
		return err
	}

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

	fileName := filepath.Base(file)
	// Append absolute file path, oid, filename into stage.txt
	content := absPath + " " + string(oid) + " " + fileName + "\n"
	_, err = stageFile.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
