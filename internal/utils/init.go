package utils

import (
	"errors"
	"os"
)

type LVCSInitManager struct {
	lvcsBaseManager
}

// creates a new LVCSInit instance
func NewLVCSInitManager(lvcsPath string) *LVCSInitManager {
	return &LVCSInitManager{
		lvcsBaseManager: newLVCSBaseManager(lvcsPath),
	}
}

func (lvcsInit *LVCSInitManager) AlreadyInit() bool {
	_, err := os.Stat(lvcsInit.lvcsPath)
	// exists then err == nil
	return err == nil
}

func (lvcsInit *LVCSInitManager) Init() error {
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
	// create default branch
	err = os.Mkdir(lvcsInit.lvcsCommitPath+"/master", 0755)
	if err != nil {
		return errors.New("failed to create the default branch: .lvcs/commits/master")
	}

	// create trees
	err = os.Mkdir(lvcsInit.lvcsTreePath, 0755)
	if err != nil {
		return errors.New("failed to create .lvcs/trees")
	}
	// create default branch tree
	file, err := os.Create(lvcsInit.lvcsTreePath + "/master_tree.txt")
	if err != nil {
		return errors.New("failed to create the default branch: .lvcs/trees/master_tree.txt")
	}
	defer file.Close()

	file, err = os.Create(lvcsInit.lvcsStagePath)
	if err != nil {
		return errors.New("failed to create .lvcs/stage.txt")
	}
	file.Close()

	file, err = os.OpenFile(lvcsInit.lvcsCurrentRefPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return errors.New("failed to open .lvcs/currentRef.txt")
	}
	defer file.Close()

	// Write "master" to the file as default branch
	_, err = file.WriteString("master\nHEAD\n")
	if err != nil {
		return errors.New("failed to write to .lvcs/currentRef.txt")
	}
	return nil
}
