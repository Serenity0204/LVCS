package utils_test

import (
	"testing"

	"github.com/Serenity0204/LVCS/internal/utils"
)

func TestBranch(t *testing.T) {
	lvcsInit := utils.NewLVCSInitManager(lvcsTestDir)
	if !lvcsInit.AlreadyInit() {
		err := lvcsInit.Init()
		if err != nil {
			t.Errorf("create LVCS DIR failed")
		}
	}

	main := "main"
	test2 := "test2"
	lvcsBranch := utils.NewLVCSBranchManager(lvcsTestDir)

	if !lvcsBranch.BranchExists(main) {
		err := lvcsBranch.CreateBranch(main)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	if !lvcsBranch.BranchExists(test2) {
		err := lvcsBranch.CreateBranch(test2)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
}

func TestBranchDeleteCheckout(t *testing.T) {
	main := "main"
	lvcsBranch := utils.NewLVCSBranchManager(lvcsTestDir)

	if lvcsBranch.BranchExists(main) {
		err := lvcsBranch.DeleteBranch(main)
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	test1 := "test1"
	if lvcsBranch.BranchExists(test1) {
		err := lvcsBranch.CheckoutBranch(test1)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
}
