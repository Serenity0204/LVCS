package helper

import (
	"errors"
	"os"
)

const lvcsDir string = ".lvcs"
const lvcsTestDir = "../.lvcs"

func Init(lvcsPath string) error {
	_, err := os.Stat(lvcsPath)
	if err == nil {
		return errors.New(".lvcs directory already exists")
	}

	err = os.Mkdir(lvcsPath, 0755)
	if err != nil {
		return errors.New("Failed to create .lvcs")
	}
	lvcsObjPath := lvcsPath + "/objects"
	err = os.Mkdir(lvcsObjPath, 0755)
	if err != nil {
		return errors.New("Failed to create .lvcs/objects")
	}
	lvcsStagePath := lvcsPath + "/stage.txt"
	_, err = os.Create(lvcsStagePath)
	if err != nil {
		return errors.New("Failed to create .lvcs/stage.txt")
	}
	return nil
}
