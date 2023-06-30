package helper

import (
	"os"
	"testing"
)

func TestHashObject(t *testing.T) {
	_, err := os.Stat(lvcsTestDir)
	// if does not exist
	if err != nil {
		err = Init(lvcsTestDir)
		if err != nil {
			t.Errorf("Create LVCS DIR failed")
		}
	}

	path := "../test_data/a.txt"
	_, err = HashObject(path, lvcsTestDir)
	if err != nil {
		t.Errorf("Failed to hash object at %s", path)
	}
	path = "../test_data/b.txt"
	_, err = HashObject(path, lvcsTestDir)
	if err != nil {
		t.Errorf("Failed to hash object at %s", path)
	}
}
