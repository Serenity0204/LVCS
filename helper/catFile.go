package helper

import "os"

func CatFile(oid string, lvcsPath string) (string, error) {
	relativePath := lvcsPath + "/objects/" + oid
	_, err := os.Stat(relativePath)
	if err != nil {
		return "", err
	}

	content, err := os.ReadFile(relativePath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
