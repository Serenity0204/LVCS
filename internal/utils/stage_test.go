package utils_test

import (
	"testing"

	"github.com/Serenity0204/LVCS/internal/utils"
)

func TestStage(t *testing.T) {
	lvcsInit := utils.NewLVCSInitManager(lvcsTestDir)
	if !lvcsInit.AlreadyInit() {
		err := lvcsInit.Init()
		if err != nil {
			t.Errorf("create LVCS DIR failed")
		}
	}

	lvcsStage := utils.NewLVCSStageManager(lvcsTestDir)

	path := "../../test_data/a.txt"
	err := lvcsStage.Add(path)
	if err != nil {
		t.Errorf("failed to add %s: %s", path, err.Error())
	}

	path = "../../test_data/b.txt"
	err = lvcsStage.Add(path)
	if err != nil {
		t.Errorf("failed to add %s: %s", path, err.Error())
	}

	path = "../../test_data/ok/abc.txt"
	err = lvcsStage.Add(path)
	if err != nil {
		t.Errorf("failed to add %s: %s", path, err.Error())
	}

	err = lvcsStage.RemoveStageContent(path)
	if err != nil {
		t.Errorf(err.Error())
	}
	err = lvcsStage.RemoveAllStageContent()
	if err != nil {
		t.Errorf(err.Error())
	}
}
