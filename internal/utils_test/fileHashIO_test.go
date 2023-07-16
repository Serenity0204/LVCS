package utils_test

import (
	"testing"

	"github.com/Serenity0204/LVCS/internal/utils"
)

func TestHashObject(t *testing.T) {
	lvcsInit := utils.NewLVCSInitManager(lvcsTestDir)
	if !lvcsInit.AlreadyInit() {
		err := lvcsInit.Init()
		if err != nil {
			t.Errorf("create LVCS DIR failed")
		}
	}

	lvcsFileHashIO := utils.NewLVCSFileHashIOManager(lvcsTestDir)

	path := "../../test_data/a.txt"
	_, err := lvcsFileHashIO.HashObject(path)
	if err != nil {
		t.Errorf("failed to hash object at %s", path)
	}
	path = "../../test_data/b.txt"
	_, err = lvcsFileHashIO.HashObject(path)
	if err != nil {
		t.Errorf("failed to hash object at %s", path)
	}

	path = "../../test_data/ok"
	_, err = lvcsFileHashIO.HashObject(path)
	if err == nil {
		t.Errorf("error is not supposed to be nil at %s", path)
	}
}

func TestCatFile(t *testing.T) {
	lvcsInit := utils.NewLVCSInitManager(lvcsTestDir)
	if !lvcsInit.AlreadyInit() {
		err := lvcsInit.Init()
		if err != nil {
			t.Errorf("create LVCS DIR failed")
		}
	}

	lvcsFileHashIO := utils.NewLVCSFileHashIOManager(lvcsTestDir)
	oid := "533c92f19449072cbc46ba5719362362d8db2e3c"
	content, err := lvcsFileHashIO.CatFile(oid)
	expectedContent := "I am BBBBB"
	if err != nil {
		t.Errorf("failed to open the file in objects dir, %s DNE", string(oid))
	}
	if content != expectedContent {
		t.Errorf("content Wrong, Expected %s but received %s", expectedContent, content)
	}

	oid = "84e3a8e13916a5e48349e49fe16cfab6a384b4a9"
	expectedContent = "To implement functionality similar to cat-file in Git, where you convert an object ID (OID) back to its corresponding string content"
	content, err = lvcsFileHashIO.CatFile(oid)
	if err != nil {
		t.Errorf("failed to open the file in objects dir, %s DNE", string(oid))
	}
	if content != expectedContent {
		t.Errorf("content Wrong, Expected %s but received %s", expectedContent, content)
	}
}
