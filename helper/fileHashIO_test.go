package helper

import (
	"testing"
)

func TestHashObject(t *testing.T) {
	lvcsInit := NewLVCSInitManager(lvcsTestDir)
	if !lvcsInit.AlreadyInit() {
		err := lvcsInit.Init()
		if err != nil {
			t.Errorf("Create LVCS DIR failed")
		}
	}

	lvcsFileHashIO := NewLVCSFileHashIOManager(lvcsTestDir)

	path := "../test_data/a.txt"
	_, err := lvcsFileHashIO.HashObject(path)
	if err != nil {
		t.Errorf("Failed to hash object at %s", path)
	}
	path = "../test_data/b.txt"
	_, err = lvcsFileHashIO.HashObject(path)
	if err != nil {
		t.Errorf("Failed to hash object at %s", path)
	}

	path = "../test_data/ok"
	_, err = lvcsFileHashIO.HashObject(path)
	if err == nil {
		t.Errorf("Error is not supposed to be nil at %s", path)
	}
}

func TestCatFile(t *testing.T) {
	lvcsInit := NewLVCSInitManager(lvcsTestDir)
	if !lvcsInit.AlreadyInit() {
		err := lvcsInit.Init()
		if err != nil {
			t.Errorf("Create LVCS DIR failed")
		}
	}

	lvcsFileHashIO := NewLVCSFileHashIOManager(lvcsTestDir)
	oid := "533c92f19449072cbc46ba5719362362d8db2e3c"
	content, err := lvcsFileHashIO.CatFile(oid)
	expectedContent := "I am BBBBB"
	if err != nil {
		t.Errorf("Failed to open the file in objects dir, %s DNE", string(oid))
	}
	if content != expectedContent {
		t.Errorf("Content Wrong, Expected %s but received %s", expectedContent, content)
	}

	oid = "84e3a8e13916a5e48349e49fe16cfab6a384b4a9"
	expectedContent = "To implement functionality similar to cat-file in Git, where you convert an object ID (OID) back to its corresponding string content"
	content, err = lvcsFileHashIO.CatFile(oid)
	if err != nil {
		t.Errorf("Failed to open the file in objects dir, %s DNE", string(oid))
	}
	if content != expectedContent {
		t.Errorf("Content Wrong, Expected %s but received %s", expectedContent, content)
	}
}
