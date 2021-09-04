package ls

import (
	"errors"
	"syscall"
)

// isHiddenWindows checks if the file is hidden in Windows
func isHiddenWindows(filename string) (bool, error) {
	pointer, err := syscall.UTF16PtrFromString(filename)
	if err != nil {
		return false, errors.New("can't get file info")
	}
	attributes, err := syscall.GetFileAttributes(pointer)
	if err != nil {
		return false, errors.New("can't get file attributes")
	}
	return attributes&syscall.FILE_ATTRIBUTE_HIDDEN != 0, nil
}
