package utils

import (
	"archive/zip"
	"bufio"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type LVCSRestoreManager struct {
	lvcsBaseManager
}

// creates a new LVCSRestoreManager instance
func NewLVCSRestoreManager(lvcsPath string) *LVCSRestoreManager {
	return &LVCSRestoreManager{
		lvcsBaseManager: newLVCSBaseManager(lvcsPath),
	}
}

// restore data by commit version under current branch
func (lvcsRestore *LVCSRestoreManager) Restore(version string) error {

	// check if version exists first
	lvcsBranch := NewLVCSBranchManager(lvcsRestore.lvcsPath)
	curBranch, err := lvcsBranch.GetCurrentBranch()
	if err != nil {
		return err
	}
	lvcsCommit := NewLVCSCommitManager(lvcsRestore.lvcsPath)
	exist, err := lvcsCommit.versionExists(curBranch, version)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("version:" + version + " does not exist")
	}

	// read it
	versionPath := lvcsRestore.lvcsCommitPath + "/" + curBranch + "/" + version + ".txt"
	toBeCreated, oids, err := lvcsRestore.getFilesAndOIDs(versionPath)
	if err != nil {
		return err
	}

	// create temp folder
	err = lvcsRestore.createTempFolder()
	if err != nil {
		return err
	}
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	tempDir := wd + "/temp"
	err = lvcsRestore.createDirectoriesAndFiles(toBeCreated, tempDir)
	if err != nil {
		return err
	}

	// write oid content
	err = lvcsRestore.writeOIDContent(toBeCreated, oids)
	if err != nil {
		return err
	}

	// zip it
	err = lvcsRestore.createZipFile(tempDir, version, curBranch)
	if err != nil {
		return err
	}
	// remove temp
	err = os.RemoveAll(tempDir)
	if err != nil {
		return err
	}
	return nil
}

func (lvcsRestore *LVCSRestoreManager) createTempFolder() error {
	// Get the absolute path of the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	// temp directory
	tempPath := wd + "/temp"
	_, err = os.Stat(tempPath)
	if err != nil {
		err = os.Mkdir(tempPath, 0755)
		if err != nil {
			return err
		}
	} else {
		err = os.RemoveAll(tempPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func (lvcsRestore *LVCSRestoreManager) createEmptyFile(destPath string) error {
	file, err := os.Create(destPath)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}

func (lvcsRestore *LVCSRestoreManager) createDirectoriesAndFiles(files []string, destDir string) error {
	for _, file := range files {
		destPath := filepath.Join(destDir, file)
		destDirPath := filepath.Dir(destPath)

		// Create the destination directory if it doesn't exist
		if _, err := os.Stat(destDirPath); os.IsNotExist(err) {
			err = os.MkdirAll(destDirPath, 0755)
			if err != nil {
				return err
			}
		}

		// Create an empty file in the destination directory
		err := lvcsRestore.createEmptyFile(destPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func (lvcsRestore *LVCSRestoreManager) getFilesAndOIDs(versionPath string) ([]string, []string, error) {
	versionFile, err := os.Open(versionPath)
	if err != nil {
		return []string{}, []string{}, err
	}
	defer versionFile.Close()

	// hold data
	var toBeCreated []string
	var oids []string

	scanner := bufio.NewScanner(versionFile)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 3)
		if len(parts) != 3 {
			continue
		}
		toBeCreated = append(toBeCreated, parts[2])
		oids = append(oids, parts[1])
	}
	if len(toBeCreated) == 0 && len(oids) == 0 {
		return []string{}, []string{}, errors.New("cannot restore empty commit")
	}
	return toBeCreated, oids, nil
}

func (lvcsRestore *LVCSRestoreManager) writeOIDContent(toBeCreated []string, oids []string) error {
	lvcsFileIO := NewLVCSFileHashIOManager(lvcsRestore.lvcsPath)

	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	if len(oids) != len(toBeCreated) {
		return errors.New("length not equal")
	}
	for i, file := range toBeCreated {
		filePath := wd + "/temp/" + file
		fileHandle, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		defer fileHandle.Close()
		// Write CatFile content into the file
		content, err := lvcsFileIO.CatFile(oids[i])
		if err != nil {
			return err
		}
		_, err = fileHandle.WriteString(content)
		if err != nil {
			return err
		}
		fileHandle.Close()
	}
	return nil
}

func (lvcsRestore *LVCSRestoreManager) createZipFile(srcDir string, version string, branchName string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	zipFilePath := wd + "/" + branchName + "_" + version + ".zip"
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = filepath.ToSlash(relPath)

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}
