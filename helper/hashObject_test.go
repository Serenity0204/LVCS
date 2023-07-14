package helper

import (
	"testing"
)

func TestHashObject(t *testing.T) {
	lvcsInit := NewLVCSInit(lvcsTestDir)
	if !lvcsInit.AlreadyInit() {
		err := lvcsInit.Init()
		if err != nil {
			t.Errorf("Create LVCS DIR failed")
		}
	}

	path := "../test_data/a.txt"
	_, err := HashObject(path, lvcsTestDir)
	if err != nil {
		t.Errorf("Failed to hash object at %s", path)
	}
	path = "../test_data/b.txt"
	_, err = HashObject(path, lvcsTestDir)
	if err != nil {
		t.Errorf("Failed to hash object at %s", path)
	}

	path = "../test_data/ok"
	_, err = HashObject(path, lvcsTestDir)
	if err == nil {
		t.Errorf("Error is not supposed to be nil at %s", path)
	}
}
