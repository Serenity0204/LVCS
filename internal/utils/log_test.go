package utils_test

import (
	"testing"

	"github.com/Serenity0204/LVCS/internal/utils"
)

func TestLog(t *testing.T) {
	lvcsInit := utils.NewLVCSInitManager(lvcsTestDir)
	if !lvcsInit.AlreadyInit() {
		err := lvcsInit.Init()
		if err != nil {
			t.Errorf("create LVCS DIR failed")
		}
	}

	lvcsLogger := utils.NewLVCSLogManager(lvcsTestDir)

	_, err := lvcsLogger.Log()
	if err != nil {
		t.Errorf(err.Error())
	}
}
