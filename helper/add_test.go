package helper

import (
	"testing"
)

func TestAdd(t *testing.T) {
	lvcsInit := NewLVCSInitManager(lvcsTestDir)
	if !lvcsInit.AlreadyInit() {
		err := lvcsInit.Init()
		if err != nil {
			t.Errorf("Create LVCS DIR failed")
		}
	}

	lvcsAdd := NewLVCSAddManager(lvcsTestDir)

	path := "../test_data/a.txt"
	err := lvcsAdd.Add(path)
	if err != nil {
		t.Errorf("Failed to add %s: %s", path, err.Error())
	}

	path = "../test_data/b.txt"
	err = lvcsAdd.Add(path)
	if err != nil {
		t.Errorf("Failed to add %s: %s", path, err.Error())
	}

	path = "../test_data/ok/abc.txt"
	err = lvcsAdd.Add(path)
	if err != nil {
		t.Errorf("Failed to add %s: %s", path, err.Error())
	}
}
