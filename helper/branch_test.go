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
			t.Errorf("create branch:" + main + " failed")
		}
	}
	if !lvcsBranch.BranchExists(test2) {
		err := lvcsBranch.CreateBranch(test2)
		if err != nil {
			t.Errorf("create branch:" + test2 + " failed")
		}
	}
	err := lvcsBranch.DeleteBranch(main)
	if err != nil {
		t.Errorf("delete branch:" + main + " failed")
	}
	test1 := "test1"
	err = lvcsBranch.CheckoutBranch(test1)
	if err != nil {
		t.Errorf("checkout branch:" + test1 + " failed")
	}
}
