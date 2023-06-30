package helper

import (
	"os"
	"testing"
)

func TestCatFile(t *testing.T) {
	_, err := os.Stat(lvcsTestDir)
	if err != nil {
		return
	}
	oid := "533c92f19449072cbc46ba5719362362d8db2e3c"
	content, err := CatFile(oid, lvcsTestDir)
	expectedContent := "I am BBBBB"
	if err != nil {
		t.Errorf("Failed to open the file in objects dir, %s DNE", string(oid))
	}
	if content != expectedContent {
		t.Errorf("Content Wrong, Expected %s but received %s", expectedContent, content)
	}

	oid = "84e3a8e13916a5e48349e49fe16cfab6a384b4a9"
	expectedContent = "To implement functionality similar to cat-file in Git, where you convert an object ID (OID) back to its corresponding string content"
	content, err = CatFile(oid, lvcsTestDir)
	if err != nil {
		t.Errorf("Failed to open the file in objects dir, %s DNE", string(oid))
	}
	if content != expectedContent {
		t.Errorf("Content Wrong, Expected %s but received %s", expectedContent, content)
	}
}
