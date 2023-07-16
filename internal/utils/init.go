package utils

import (
	"errors"
	"os"
)

type LVCSInitManager struct {
	lvcsPath           string
	lvcsObjPath        string
	lvcsCommitPath     string
	lvcsStagePath      string
	lvcsCurrentRefPath string
}

// creates a new LVCSInit instance
func NewLVCSInitManager(lvcsPath string) *LVCSInitManager {
	return &LVCSInitManager{
		lvcsPath:           lvcsPath,
		lvcsObjPath:        lvcsPath + "/objects",
		lvcsCommitPath:     lvcsPath + "/commits",
		lvcsStagePath:      lvcsPath + "/stage.txt",
		lvcsCurrentRefPath: lvcsPath + "/currentRef.txt",
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

	_, err = os.Create(lvcsInit.lvcsStagePath)
	if err != nil {
		return errors.New("failed to create .lvcs/stage.txt")
	}

	file, err := os.OpenFile(lvcsInit.lvcsCurrentRefPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
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
