package internal_test

import (
	"os"
	"testing"

	"github.com/Serenity0204/LVCS/internal"
)

const lvcsManTestDir string = "../.lvcs_test"

func removeLVCSTestDir() error {
	err := os.RemoveAll(lvcsManTestDir)
	if err != nil {
		return err
	}
	return nil
}

func TestLVCSManager(t *testing.T) {
	err := removeLVCSTestDir()
	if err != nil {
		t.Errorf("failed to delete existing dir: %s", err.Error()) // Convert error to string
	}

	lvcsMan := internal.NewLVCSManager(lvcsManTestDir)
	_, err = lvcsMan.Execute("init", []string{})
	if err != nil {
		t.Errorf("create .lvcs_test DIR failed: %s", err.Error()) // Convert error to string
	}

	oid, err := lvcsMan.Execute("hashObject", []string{"../test_data/a.txt"})
	if err != nil || oid != "84e3a8e13916a5e48349e49fe16cfab6a384b4a9" {
		t.Errorf("hash-object failed: %s", err.Error()) // Convert error to string
	}

	content, err := lvcsMan.Execute("catFile", []string{"84e3a8e13916a5e48349e49fe16cfab6a384b4a9"})
	expectedContent := "To implement functionality similar to cat-file in Git, where you convert an object ID (OID) back to its corresponding string content"
	if err != nil || content != expectedContent {
		t.Errorf("cat-file failed: %s", err.Error()) // Convert error to string
	}

	_, err = lvcsMan.Execute("add", []string{"../test_data/a.txt", "../test_data/b.txt", "../test_data/ok/abc.txt"})
	if err != nil {
		t.Errorf("add failed: %s", err.Error()) // Convert error to string
	}
}
