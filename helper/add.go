package helper

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

type LVCSAddManager struct {
	lvcsPath      string
	lvcsStagePath string
}

// NewLVCSAddManager creates a new LVCSAdd instance
func NewLVCSAddManager(lvcsPath string) *LVCSAddManager {
	return &LVCSAddManager{
		lvcsPath:      lvcsPath,
		lvcsStagePath: lvcsPath + "/stage.txt",
	}
}

// GetAbsolutePath returns the absolute path of a file given its relative path
func (lvcsAdd *LVCSAddManager) getAbsolutePath(relativePath string) (string, error) {
	// Get the absolute path of the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	// Get the absolute path of the file by joining the working directory and the relative path
	absolutePath := filepath.Join(wd, relativePath)
	return absolutePath, nil
}

func (lvcsAdd *LVCSAddManager) keyExists(target string) (bool, error) {
	// Open the file
	file, err := os.Open(lvcsAdd.lvcsStagePath)
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
func (lvcsAdd *LVCSAddManager) Add(file string, lvcsPath string) error {

	absPath, err := lvcsAdd.getAbsolutePath(file)
	if err != nil {
		return err
	}

	isIn, err := lvcsAdd.keyExists(absPath)
	if err != nil {

		return err
	}

	// if it already exists, don't re-add.
	if isIn {
		return nil
	}

	// if it DNE then hash object it
	lvcsFileHashIO := NewLVCSFileHashIOManager(lvcsAdd.lvcsPath)
	oid, err := lvcsFileHashIO.HashObject(file)
	if err != nil {
		return err
	}

	// open the file
	stageFile, err := os.OpenFile(lvcsAdd.lvcsStagePath, os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer stageFile.Close()

	// fileName := filepath.Base(file)
	// Append absolute file path, oid, filename into stage.txt
	content := absPath + " " + string(oid) + "\n"
	_, err = stageFile.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
