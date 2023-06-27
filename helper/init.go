package helper

import "os"

const GIT_DIR string = ".lvcs"

func Init() {
	os.Mkdir(GIT_DIR, 0755)
}
