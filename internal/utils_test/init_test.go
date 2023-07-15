package utils_test

import (
	"testing"

	"github.com/Serenity0204/LVCS/internal/utils"
)

// const LvcsDir string = ".lvcs"
const lvcsTestDir string = "../../.lvcs"

func TestInit(t *testing.T) {
	lvcsInit := utils.NewLVCSInitManager(lvcsTestDir)
	if !lvcsInit.AlreadyInit() {
		err := lvcsInit.Init()
		if err != nil {
			t.Errorf("Create LVCS DIR failed")
		}
	}
}
