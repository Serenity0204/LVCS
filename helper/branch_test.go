package helper

import "testing"

func TestBranch(t *testing.T) {
	if !AlreadyInit(lvcsTestDir) {
		err := Init(lvcsTestDir)
		if err != nil {
			t.Errorf("Create LVCS DIR failed")
		}
	}

	main := "main"
	test2 := "test2"
	commitFolderPath := lvcsTestDir + "/commits/"
	if !BranchExists(commitFolderPath, main) {
		err := CreateBranch(commitFolderPath, main)
		if err != nil {
			t.Errorf("Create branch:" + main + " failed")
		}
	}
	if !BranchExists(commitFolderPath, test2) {
		err := CreateBranch(commitFolderPath, test2)
		if err != nil {
			t.Errorf("Create branch:" + test2 + " failed")
		}
	}
}
