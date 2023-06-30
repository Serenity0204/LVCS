package helper

import (
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	_, err := os.Stat(lvcsTestDir)
	// if does not exist, then init
	if err != nil {
		err = Init(lvcsTestDir)
		if err != nil {
			t.Errorf("Create LVCS DIR failed")
		}
	}
}
