package helper

import (
	"errors"
	"os"
)

const LvcsDir string = ".lvcs"
const lvcsTestDir string = "../.lvcs"

type LVCSInit struct {
	lvcsPath       string
	lvcsObjPath    string
	lvcsCommitPath string
	lvcsStagePath  string
}

// NewLVCSAdd creates a new LVCSInit instance
func NewLVCSInit(lvcsPath string) *LVCSInit {
	return &LVCSInit{
		lvcsPath:       lvcsPath,
		lvcsObjPath:    lvcsPath + "/objects",
		lvcsCommitPath: lvcsPath + "/commits",
		lvcsStagePath:  lvcsPath + "/stage.txt",
	}
}

func (lvcsInit *LVCSInit) AlreadyInit() bool {
	_, err := os.Stat(lvcsInit.lvcsPath)
	// exists then err == nil
	return err == nil
}

func (lvcsInit *LVCSInit) Init() error {
	if lvcsInit.AlreadyInit() {
		return errors.New(".lvcs directory already exists")
	}

	err := os.Mkdir(lvcsInit.lvcsPath, 0755)
	if err != nil {
		return errors.New("failed to create .lvcs")
	}

	err = os.Mkdir(lvcsInit.lvcsObjPath, 0755)
	if err != nil {
		return errors.New("failed to create .lvcs/objects")
	}

	err = os.Mkdir(lvcsInit.lvcsCommitPath, 0755)
	if err != nil {
		return errors.New("failed to create .lvcs/commits")
	}

	_, err = os.Create(lvcsInit.lvcsStagePath)
	if err != nil {
		return errors.New("failed to create .lvcs/stage.txt")
	}
	return nil
}
