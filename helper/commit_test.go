package helper

import (
	"os"
	"testing"
)

func TestCommit(t *testing.T) {
	// check if file exists
	if !AlreadyInit(lvcsTestDir) {
		Init(lvcsTestDir)
		return
	}

	fileInfo, err := os.Stat(lvcsTestDir + "/stage.txt")
	if err != nil {
		t.Errorf("Init failed inside commit")
	}

	// Check if file size is 0
	if fileInfo.Size() == 0 {
		return
	}
	// branches
	master := "master"
	test1 := "test1"

	commitFolderPath := lvcsTestDir + "/commits/"
	if !BranchExists(commitFolderPath, master) {
		CreateBranch(commitFolderPath, master)
	}
	if !BranchExists(commitFolderPath, test1) {
		CreateBranch(commitFolderPath, test1)
	}
	// not 0 then commit
	err = Commit(lvcsTestDir, master)
	if err != nil {
		t.Errorf("Failed to commit " + master)
	}
	err = Commit(lvcsTestDir, test1)
	if err != nil {
		t.Errorf("Failed to commit " + test1)
	}
}
