package helper

import "testing"

func TestBranch(t *testing.T) {
	lvcsInit := NewLVCSInitManager(lvcsTestDir)
	if !lvcsInit.AlreadyInit() {
		err := lvcsInit.Init()
		if err != nil {
			t.Errorf("Create LVCS DIR failed")
		}
	}

	main := "main"
	test2 := "test2"
	lvcsBranch := NewLVCSBranchManager(lvcsTestDir)

	if !lvcsBranch.BranchExists(main) {
		err := lvcsBranch.CreateBranch(main)
		if err != nil {
			t.Errorf("Create branch:" + main + " failed")
		}
	}
	if !lvcsBranch.BranchExists(test2) {
		err := lvcsBranch.CreateBranch(test2)
		if err != nil {
			t.Errorf("Create branch:" + test2 + " failed")
		}
	}
}
