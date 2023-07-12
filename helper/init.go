package helper

import (
	"errors"
	"os"
)

const LvcsDir string = ".lvcs"
const lvcsTestDir string = "../.lvcs"

func AlreadyInit(lvcsPath string) bool {
	_, err := os.Stat(lvcsPath)
	// exists then err == nil
	return err == nil
}

func Init(lvcsPath string) error {
	if AlreadyInit(lvcsPath) {
		return errors.New(".lvcs directory already exists")
	}

	err := os.Mkdir(lvcsPath, 0755)
	if err != nil {
		return errors.New("failed to create .lvcs")
	}
	lvcsObjPath := lvcsPath + "/objects"
	err = os.Mkdir(lvcsObjPath, 0755)
	if err != nil {
		return errors.New("failed to create .lvcs/objects")
	}

	lvcsCommitPath := lvcsPath + "/commits"
	err = os.Mkdir(lvcsCommitPath, 0755)
	if err != nil {
		return errors.New("failed to create .lvcs/commits")
	}

	lvcsStagePath := lvcsPath + "/stage.txt"
	_, err = os.Create(lvcsStagePath)
	if err != nil {
		return errors.New("failed to create .lvcs/stage.txt")
	}
	return nil
}
