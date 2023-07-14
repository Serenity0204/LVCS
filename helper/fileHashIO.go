package helper

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"io"
	"os"
)

type LVCSFileHashIO struct {
	lvcsPath    string
	lvcsObjPath string
}

// NewLVCSAdd creates a new LVCSInit instance
func NewLVCSFileHashIO(lvcsPath string) *LVCSFileHashIO {
	return &LVCSFileHashIO{
		lvcsPath:    lvcsPath,
		lvcsObjPath: lvcsPath + "/objects",
	}
}

func (lvcsFileHashIO *LVCSFileHashIO) CatFile(oid string) (string, error) {
	relativePath := lvcsFileHashIO.lvcsObjPath + "/" + oid
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

func (lvcsFileHashIO *LVCSFileHashIO) HashObject(file string) (string, error) {
	info, err := os.Stat(file)
	if err != nil {
		return "", err
	}

	if info.IsDir() {
		return "", errors.New("cannot add a directory")
	}

	content, err := os.ReadFile(file)
	if err != nil {
		return "", nil
	}

	dataBytes := []byte(content)
	hash := sha1.New()
	hash.Write(dataBytes)
	oid := hex.EncodeToString(hash.Sum(nil))

	relativePath := lvcsFileHashIO.lvcsObjPath + "/" + oid
	_, err = os.Stat(relativePath)
	// if already exists
	if err == nil {
		return oid, nil
	}
	// else create
	newFile, err := os.Create(relativePath)
	if err != nil {
		return "", err
	}
	defer newFile.Close()

	// Write the content into the file
	_, err = io.Copy(newFile, bytes.NewReader(content))
	if err != nil {
		// Clean up the file if writing fails
		os.Remove(relativePath)
		return "", err
	}
	return oid, nil
}
