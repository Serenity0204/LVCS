package helper

import (
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	_, err := os.Stat(lvcsTestDir)
	// if does not exist
	if err != nil {
		err = Init(lvcsTestDir)
		if err != nil {
			t.Errorf("Create LVCS DIR failed")
		}
	}
	path := "../test_data/a.txt"
	err = Add(path, lvcsTestDir)
	if err != nil {
		t.Errorf("Failed to add %s", path)
	}
	path = "../test_data/b.txt"
	err = Add(path, lvcsTestDir)
	if err != nil {
		t.Errorf("Failed to add %s", path)
	}
}
