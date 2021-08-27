package ls

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

func readDir(root string) ([]string, error) {
	exist, err := exists(root)
	if err != nil {
		return nil, err
	}

	if exist {
		isDirectory, err := isDir(root)
		if err != nil {
			return nil, err
		}
		if isDirectory {
			var files []string

			fileInfo, err := ioutil.ReadDir(root)
			if err != nil {
				return files, err
			}

			for _, file := range fileInfo {
				files = append(files, file.Name())
			}

			return files, nil
		} else {
			return []string{filepath.Base(root)}, nil
		}
	}

	return nil, err
}

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

func isDir(root string) (bool, error) {
	fileInfo, err := os.Stat(root)
	if err != nil {
		return false, errors.New("can't get file info")
	}

	return fileInfo.IsDir(), nil
}
