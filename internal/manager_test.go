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
		t.Errorf("failed to delete existing dir")
	}

	lvcsMan := internal.NewLVCSManager(lvcsManTestDir)
	err = lvcsMan.Execute("init", []string{})
	if err != nil {
		t.Errorf("create .lvcs_test DIR failed")
	}
}
