package utils

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/Serenity0204/LVCS/internal/models"
)

type LVCSCommitManager struct {
	lvcsBaseManager
}

// creates a new LVCSCommit instance
func NewLVCSCommitManager(lvcsPath string) *LVCSCommitManager {
	return &LVCSCommitManager{
		lvcsBaseManager: newLVCSBaseManager(lvcsPath),
	}
}

func (lvcsCommit *LVCSCommitManager) versionExists(branchName string, version string) (bool, error) {
	files, err := os.ReadDir(lvcsCommit.lvcsCommitPath + "/" + branchName)
	if err != nil {
		return false, err
	}
	for _, file := range files {
		// if it's not a file then it's error
		if file.IsDir() {
			return false, errors.New("read A Dir Inside commits")
		}
		// it's a file
		fileName := file.Name()
		// check for error
		if !strings.HasPrefix(fileName, "v") {
			return false, errors.New("wrong version number")
		}
		versionStr := strings.TrimSuffix(fileName, ".txt")
		if version == versionStr {
			return true, nil
		}
	}
	return false, nil
}

func (lvcsCommit *LVCSCommitManager) GetLatestVersion() (string, error) {
	lvcsBranch := NewLVCSBranchManager(lvcsCommit.lvcsPath)
	curBranchName, err := lvcsBranch.GetCurrentBranch()
	if err != nil {
		return "", err
	}
	ver, err := lvcsCommit.getLatestVersion(curBranchName)
	if err != nil {
		return "", err
	}
	if ver == -1 {
		return string("HEAD"), nil
	}
	return string("v" + strconv.Itoa(ver)), nil
}

func (lvcsCommit *LVCSCommitManager) GetCurrentVersion() (string, error) {
	version, err := lvcsCommit.getCurrentVersion()
	if err != nil {
		return "", err
	}
	if version == -1 {
		return string("HEAD"), nil
	}
	return string("v" + strconv.Itoa(version)), nil
}

func (lvcsCommit *LVCSCommitManager) getCurrentVersion() (int, error) {
	file, err := os.Open(lvcsCommit.lvcsCurrentRefPath)
	if err != nil {
		return -1, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	success := scanner.Scan()
	if !success {
		return -1, errors.New("scanned failed")
	}
	success = scanner.Scan()
	if !success {
		return -1, errors.New("scanned failed")
	}
	curVersion := scanner.Text()
	if curVersion == "HEAD" {
		return -1, nil
	}
	version, err := strconv.Atoi(strings.TrimPrefix(curVersion, "v"))
	if err != nil {
		return -1, err
	}
	return version, nil
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

func (lvcsCommit *LVCSCommitManager) removeDuplicateParentContent(branchName string, version string) error {
	versionPath := lvcsCommit.lvcsCommitPath + "/" + branchName + "/" + version + ".txt"
	file, err := os.Open(versionPath)
	if err != nil {
		return err
	}
	defer file.Close()
	// Create a map to store the latest key-value pairs
	latestKeys := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)
		if len(parts) != 2 {
			// Skip lines that are not in the expected format (path OID)
			continue
		}
		filePath := parts[0]
		oid := parts[1]
		// Update the latest value for the key
		latestKeys[filePath] = oid
	}
	err = scanner.Err()
	if err != nil {
		return err
	}
	file.Close()

	// Reopen the file in write mode to overwrite its content
	file, err = os.Create(versionPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the latest key-value pairs to the file
	for filePath, oid := range latestKeys {
		_, err := fmt.Fprintf(file, "%s %s\n", filePath, oid)
		if err != nil {
			return err
		}
	}
	return nil
}

func (lvcsCommit *LVCSCommitManager) createNewCommitRecord(branchName string, version int, inherit bool) error {
	// get the stage content
	content, err := os.ReadFile(lvcsCommit.lvcsStagePath)
	if err != nil {
		return err
	}
	// get the parent's node's content
	tree, err := lvcsCommit.getCommitTree()
	if err != nil {
		return err
	}

	parentInfo := ""
	if inherit && version != 0 {
		parent, err := tree.GetParentNode(string("v" + strconv.Itoa(version)))
		if err != nil {
			return err
		}
		// read parent's commit
		logMan := NewLVCSLogManager(lvcsCommit.lvcsPath)
		parentInfo, err = logMan.LogByVersion(parent.Value)
		if err != nil {
			return err
		}
	}

	// Combine parentInfo and content into a single byte slice
	var combinedContent bytes.Buffer
	combinedContent.WriteString(parentInfo)
	combinedContent.Write(content)

	branchPath := lvcsCommit.lvcsCommitPath + "/" + branchName + "/v" + strconv.Itoa(version) + ".txt"

	// create
	versionFile, err := os.Create(branchPath)
	if err != nil {
		return err
	}
	defer versionFile.Close()

	// Write the content into the file
	_, err = io.Copy(versionFile, &combinedContent)
	if err != nil {
		// Clean up the file if writing fails
		os.Remove(branchPath)
		return err
	}
	// remove stage content
	stageMan := NewLVCSStageManager(lvcsCommit.lvcsPath)
	err = stageMan.RemoveAllStageContent()
	if err != nil {
		return err
	}
	// remove duplicate
	err = lvcsCommit.removeDuplicateParentContent(branchName, string("v"+strconv.Itoa(version)))
	if err != nil {
		return err
	}
	return nil
}

func (lvcsCommit *LVCSCommitManager) getCommitTree() (*models.NaryTree, error) {
	lvcsBranch := NewLVCSBranchManager(lvcsCommit.lvcsPath)
	curBranchName, err := lvcsBranch.GetCurrentBranch()
	if err != nil {
		return nil, err
	}
	// get previous tree
	tree := models.NewNaryTree()
	// read tree content
	treePath := lvcsCommit.lvcsTreePath + "/" + curBranchName + "_tree.txt"
	content, err := os.ReadFile(treePath)
	if err != nil {
		return nil, err
	}
	treeData := string(content)

	err = tree.Deserialize(treeData)
	if err != nil {
		return nil, err
	}
	return tree, nil
}

// commit will commit a new version under current version
func (lvcsCommit *LVCSCommitManager) Commit(inherit bool) error {
	lvcsBranch := NewLVCSBranchManager(lvcsCommit.lvcsPath)
	curBranchName, err := lvcsBranch.GetCurrentBranch()
	if err != nil {
		return err
	}

	tree, err := lvcsCommit.getCommitTree()
	if err != nil {
		return err
	}

	// get current curVersion
	curVersion, err := lvcsCommit.getCurrentVersion()
	if err != nil {
		return err
	}
	// get the latest version number
	latestVersion, err := lvcsCommit.getLatestVersion(curBranchName)
	if err != nil {
		return err
	}
	newVersion := latestVersion + 1

	// insert
	// get parent and cur version string
	curVersionStr := "v" + strconv.Itoa(newVersion)

	if curVersion != -1 {
		// get parent node if parent is not -1
		parentVersionStr := "v" + strconv.Itoa(curVersion)
		parent, err := tree.GetNode(parentVersionStr)
		if err != nil {
			return err
		}
		// insert cur under parent
		err = tree.Insert(parent, curVersionStr)
		if err != nil {
			return err
		}
	} else {
		err = tree.Insert(nil, curVersionStr)
		if err != nil {
			return err
		}
	}
	// serialize
	newTreeData, err := tree.Serialize()
	if err != nil {

		return err
	}

	// write it back to file
	treePath := lvcsCommit.lvcsTreePath + "/" + curBranchName + "_tree.txt"
	treeFile, err := os.OpenFile(treePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer treeFile.Close()
	_, err = treeFile.WriteString(newTreeData)
	if err != nil {

		return err
	}

	// create the commit record
	err = lvcsCommit.createNewCommitRecord(curBranchName, newVersion, inherit)
	if err != nil {
		return err
	}
	// update current ref
	// default is master for branchName
	file, err := os.OpenFile(lvcsCommit.lvcsCurrentRefPath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(curBranchName + "\n" + curVersionStr + "\n")
	if err != nil {
		return err
	}

	return nil
}

func (lvcsCommit *LVCSCommitManager) SwitchCommitVersion(version string) error {
	// default is master for branchName
	lvcsBranch := NewLVCSBranchManager(lvcsCommit.lvcsPath)
	curBranchName, err := lvcsBranch.GetCurrentBranch()
	if err != nil {
		return err
	}
	exist, err := lvcsCommit.versionExists(curBranchName, version)

	if err != nil {
		return err
	}
	if !exist {
		return errors.New("version:" + version + " does not exist")
	}

	file, err := os.OpenFile(lvcsCommit.lvcsCurrentRefPath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(curBranchName + "\n" + version + "\n")
	if err != nil {
		return err
	}
	return nil
}

func (lvcsCommit *LVCSCommitManager) CommitTree() (string, error) {
	tree, err := lvcsCommit.getCommitTree()
	if err != nil {
		return "", err
	}
	if tree.GetNaryTreeRoot() == nil {
		return string("Empty Tree"), nil
	}
	return tree.NaryTreeString(), nil
}

func (lvcsCommit *LVCSCommitManager) LCA(version1 string, version2 string) (string, error) {
	tree, err := lvcsCommit.getCommitTree()
	if err != nil {
		return "", err
	}
	if tree.GetNaryTreeRoot() == nil {
		return string("Empty Tree"), nil
	}
	lca, err := tree.LCA(version1, version2)
	if err != nil {
		return "", err
	}
	lcaOutput := tree.NaryTreeString() + "\n" + "LCA:" + lca + "\n"
	return lcaOutput, nil
}

// // Will remove the commit including it's children
// func (lvcsCommit *LVCSCommitManager) RemoveCommit(version string) error {
// 	tree, err := lvcsCommit.getCommitTree()
// 	if err != nil {
// 		return err
// 	}
// 	tree.Remove(version)
// 	return nil
// }
