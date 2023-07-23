package utils

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type LVCSStageManager struct {
	lvcsBaseManager
}

// creates a new LVCSLoggerManager instance
func NewLVCSStageManager(lvcsPath string) *LVCSStageManager {
	return &LVCSStageManager{
		lvcsBaseManager: newLVCSBaseManager(lvcsPath),
	}
}

// GetAbsolutePath returns the absolute path of a file given its relative path
func (lvcsStage *LVCSStageManager) getAbsolutePath(relativePath string) (string, error) {
	// Get the absolute path of the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	// Get the absolute path of the file by joining the working directory and the relative path
	absolutePath := filepath.Join(wd, relativePath)
	return absolutePath, nil
}

func (lvcsStage *LVCSStageManager) keyExists(target string) (bool, error) {
	// Open the file
	file, err := os.Open(lvcsStage.lvcsStagePath)
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

		filePath := fields[0]
		// Check if the filePath matches the target
		if filePath == target {
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
func (lvcsStage *LVCSStageManager) Add(file string) error {
	if lvcsStage.ignoreOrAbsPath(file) {
		return errors.New("cannot add a .lvcs elements or using absolute path")
	}
	// First hash object it
	lvcsFileHashIO := NewLVCSFileHashIOManager(lvcsStage.lvcsPath)
	// if OID DNE then will create one
	oid, err := lvcsFileHashIO.HashObject(file)
	if err != nil {
		return err
	}

	absPath, err := lvcsStage.getAbsolutePath(file)
	if err != nil {
		return err
	}

	isIn, err := lvcsStage.keyExists(absPath)
	if err != nil {
		return err
	}

	// if it already exists, delete existing one
	if isIn {
		err = lvcsStage.RemoveStageContent(file)
		if err != nil {
			return err
		}
	}

	// open the file and write the new one
	stageFile, err := os.OpenFile(lvcsStage.lvcsStagePath, os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer stageFile.Close()

	// Append absolute file path, oid, filename into stage.txt
	content := absPath + " " + string(oid) + " " + file + "\n"
	_, err = stageFile.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

func (lvcsStage *LVCSStageManager) RemoveAllStageContent() error {
	stageFile, err := os.OpenFile(lvcsStage.lvcsStagePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer stageFile.Close()
	return nil
}

func (lvcsStage *LVCSStageManager) GetStageContent() (string, error) {
	stageFile, err := os.Open(lvcsStage.lvcsStagePath)
	if err != nil {
		return "", err
	}
	defer stageFile.Close()

	var lines []string

	scanner := bufio.NewScanner(stageFile)
	i := 1
	for scanner.Scan() {
		lines = append(lines, strconv.Itoa(i)+":"+scanner.Text())
		i++
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	// Concatenate the lines into a single string separated by '\n'
	fileList := "empty"
	if len(lines) != 0 {
		fileList = strings.Join(lines, "\n")
	}
	content := "Tracked Files:\n" + fileList
	return content, nil
}

func (lvcsStage *LVCSStageManager) RemoveStageContent(file string) error {

	absPath, err := lvcsStage.getAbsolutePath(file)
	if err != nil {
		return err
	}

	isIn, err := lvcsStage.keyExists(absPath)
	if err != nil {
		return err
	}

	if !isIn {
		return errors.New("file name:" + absPath + " is not tracked")
	}

	// Read the contents of the stage file
	stageFile, err := os.Open(lvcsStage.lvcsStagePath)
	if err != nil {
		return err
	}
	defer stageFile.Close()

	var linesToKeep []string

	scanner := bufio.NewScanner(stageFile)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 3)
		if len(parts) != 3 {
			// Skip lines that are not in the expected format (path OID relativePath)
			continue
		}
		filePath := parts[0]

		// Check if the first part of the line matches the absPath
		if filePath == absPath {
			continue // Skip this line
		}

		linesToKeep = append(linesToKeep, line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Close the file after reading
	if err := stageFile.Close(); err != nil {
		return err
	}

	// Reopen the file for writing
	stageFile, err = os.OpenFile(lvcsStage.lvcsStagePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer stageFile.Close()

	// Write the filtered lines back to the file
	writer := bufio.NewWriter(stageFile)
	for _, line := range linesToKeep {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	// Flush the writer to ensure all data is written to the file
	err = writer.Flush()
	if err != nil {
		return err
	}
	stageFile.Close()
	return nil
}
