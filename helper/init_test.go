package helper

import (
	"testing"
)

func TestInit(t *testing.T) {
	lvcsInit := NewLVCSInit(lvcsTestDir)
	if !lvcsInit.AlreadyInit() {
		err := lvcsInit.Init()
		if err != nil {
			t.Errorf("Create LVCS DIR failed")
		}
	}
}
