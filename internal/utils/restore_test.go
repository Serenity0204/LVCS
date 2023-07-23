package utils_test

import (
	"os"
	"testing"

	"github.com/Serenity0204/LVCS/internal/utils"
)

func TestRestore(t *testing.T) {
	lvcsInit := utils.NewLVCSInitManager(lvcsTestDir)
	if !lvcsInit.AlreadyInit() {
		err := lvcsInit.Init()
		if err != nil {
			t.Errorf("create LVCS DIR failed")
		}
	}

	// branches
	wtf := "wtf"
	lvcsBranch := utils.NewLVCSBranchManager(lvcsTestDir)

	if !lvcsBranch.BranchExists(wtf) {
		err := lvcsBranch.CreateBranch(wtf)
		if err != nil {
			t.Errorf("create branch:" + wtf + " failed")
		}
	}
	err := lvcsBranch.CheckoutBranch(wtf)
	if err != nil {
		t.Errorf(err.Error())
	}
	lvcsStage := utils.NewLVCSStageManager(lvcsTestDir)
	err = lvcsStage.Add("../../test_data/a.txt")
	if err != nil {
		t.Errorf(err.Error())
	}

	lvcsCommit := utils.NewLVCSCommitManager(lvcsTestDir)

	err = lvcsCommit.Commit(true)
	if err != nil {
		t.Errorf(err.Error())
	}

	// after add and commit, restore
	lvcsRestore := utils.NewLVCSRestoreManager(lvcsTestDir)
	err = lvcsRestore.Restore("v0")
	if err != nil {
		t.Errorf(err.Error())
	}
	os.Remove("wtf_v0.zip")
	os.RemoveAll("..\\test_data")
}
