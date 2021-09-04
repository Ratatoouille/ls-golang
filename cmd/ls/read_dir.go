package ls

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
	f "unix-ls/cmd/files"
)

func readDir(root string, flag bool) (f.FilesSlice, error) {
	if exist, err := exists(root); exist && err == nil {
		if isDirectory, err := isDir(root); isDirectory && err == nil {
			var files f.FilesSlice

			fileInfo, err := ioutil.ReadDir(root)
			if err != nil {
				return files, err
			}

			osName := runtime.GOOS

			if osName != "windows" {
				for _, file := range fileInfo {
					if flag {
						files.AddItem(file.Name())
					} else if hidden := isHiddenUnix(file.Name()); !hidden {
						files.AddItem(file.Name())
					}
				}
			} else {
				// TODO if file hidden in windows it's return err
				for _, file := range fileInfo {
					//if hidden, err := isHiddenWindows(file.Name()); !hidden && err == nil {
					files.AddItem(file.Name())
					//} else if err != nil {
					//	return nil, err
					//}
				}
			}

			return files, nil
		} else if err != nil {
			return nil, err
		} else {
			return f.FilesSlice{filepath.Base(root)}, nil
		}
	} else {
		return nil, err
	}
}
