package utils_test

import (
	"os"
	"testing"

	"github.com/Serenity0204/LVCS/internal/utils"
)

func TestCommit(t *testing.T) {
	lvcsInit := utils.NewLVCSInitManager(lvcsTestDir)
	if !lvcsInit.AlreadyInit() {
		err := lvcsInit.Init()
		if err != nil {
			t.Errorf("Create LVCS DIR failed")
		}
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
	lvcsBranch := utils.NewLVCSBranchManager(lvcsTestDir)

	if !lvcsBranch.BranchExists(master) {
		err := lvcsBranch.CreateBranch(master)
		if err != nil {
			t.Errorf("Create branch:" + master + " failed")
		}
	}
	if !lvcsBranch.BranchExists(test1) {
		err := lvcsBranch.CreateBranch(test1)
		if err != nil {
			t.Errorf("Create branch:" + test1 + " failed")
		}
	}

	lvcsCommit := utils.NewLVCSCommitManager(lvcsTestDir)
	err = lvcsCommit.Commit(master)
	if err != nil {
		t.Errorf("Failed to commit " + master)
	}
	err = lvcsCommit.Commit(test1)
	if err != nil {
		t.Errorf("Failed to commit " + test1)
	}
}
