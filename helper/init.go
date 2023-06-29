package helper

import (
	"errors"
	"os"
)

const LVCS_DIR string = ".lvcs"
const LVCS_OBJECTS string = ".lvcs/objects"

func Init() error {
	_, err := os.Stat(LVCS_DIR)
	if err == nil {
		return errors.New(".lvcs directory already exists")
	}

	err = os.Mkdir(LVCS_DIR, 0755)
	if err != nil {
		return errors.New("Failed to create .lvcs")
	}
	err = os.Mkdir(LVCS_OBJECTS, 0755)
	if err != nil {
		return errors.New("Failed to create .lvcs/objects")
	}

	return nil
}
