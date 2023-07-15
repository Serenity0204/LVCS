package utils_test

import (
	"testing"

	"github.com/Serenity0204/LVCS/internal/utils"
)

func TestAdd(t *testing.T) {
	lvcsInit := utils.NewLVCSInitManager(lvcsTestDir)
	if !lvcsInit.AlreadyInit() {
		err := lvcsInit.Init()
		if err != nil {
			t.Errorf("Create LVCS DIR failed")
		}
	}

	lvcsAdd := utils.NewLVCSAddManager(lvcsTestDir)

	path := "../../test_data/a.txt"
	err := lvcsAdd.Add(path)
	if err != nil {
		t.Errorf("Failed to add %s: %s", path, err.Error())
	}

	path = "../../test_data/b.txt"
	err = lvcsAdd.Add(path)
	if err != nil {
		t.Errorf("Failed to add %s: %s", path, err.Error())
	}

	path = "../../test_data/ok/abc.txt"
	err = lvcsAdd.Add(path)
	if err != nil {
		t.Errorf("Failed to add %s: %s", path, err.Error())
	}
}
