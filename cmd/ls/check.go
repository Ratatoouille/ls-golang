package ls

import (
	"errors"
	"os"
)

// exists check file is exist
func exists(root string) (bool, error) {
	_, err := os.Stat(root)
	if err != nil {
		if os.IsNotExist(err) {
			return false, errors.New("file not exist")
		}
		return false, errors.New("can't get file info")
	}

	return true, nil
}

// isDir check file is directory
func isDir(root string) (bool, error) {
	fileInfo, err := os.Stat(root)
	if err != nil {
		return false, errors.New("can't get file info")
	}

	return fileInfo.IsDir(), nil
}
