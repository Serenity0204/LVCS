package utils

import (
	"bufio"
	"bytes"
	"errors"
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

// commit will commit a new version under current version
func (lvcsCommit *LVCSCommitManager) Commit() error {
	// default is master for branchName
	lvcsBranch := NewLVCSBranchManager(lvcsCommit.lvcsPath)
	curBranchName, err := lvcsBranch.GetCurrentBranch()
	if err != nil {

		return err
	}

	// get previous tree
	tree := models.NewNaryTree()
	// read tree content
	treePath := lvcsCommit.lvcsTreePath + "/" + curBranchName + "_tree.txt"
	content, err := os.ReadFile(treePath)
	if err != nil {

		return err
	}
	treeData := string(content)

	err = tree.Deserialize(treeData)
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

	// create the commit record
	err = lvcsCommit.createNewCommitRecord(curBranchName, newVersion)
	if err != nil {
		return err
	}

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
	treeFile, err := os.OpenFile(treePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer treeFile.Close()
	_, err = treeFile.WriteString(newTreeData)
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
	lvcsBranch := NewLVCSBranchManager(lvcsCommit.lvcsPath)
	curBranchName, err := lvcsBranch.GetCurrentBranch()
	if err != nil {
		return "", err
	}

	// get tree
	tree := models.NewNaryTree()
	// read tree content
	treePath := lvcsCommit.lvcsTreePath + "/" + curBranchName + "_tree.txt"
	content, err := os.ReadFile(treePath)
	if err != nil {
		return "", err
	}
	treeData := string(content)

	err = tree.Deserialize(treeData)
	if err != nil {
		return "", err
	}
	if tree.GetNaryTreeRoot() == nil {
		return string("Empty Tree"), nil
	}
	return tree.NaryTreeString(), nil
}

func (lvcsCommit *LVCSCommitManager) LCA(version1 string, version2 string) (string, error) {
	lvcsBranch := NewLVCSBranchManager(lvcsCommit.lvcsPath)
	curBranchName, err := lvcsBranch.GetCurrentBranch()
	if err != nil {
		return "", err
	}

	// get tree
	tree := models.NewNaryTree()
	// read tree content
	treePath := lvcsCommit.lvcsTreePath + "/" + curBranchName + "_tree.txt"
	content, err := os.ReadFile(treePath)
	if err != nil {
		return "", err
	}
	treeData := string(content)

	err = tree.Deserialize(treeData)
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
