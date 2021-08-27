package ls

import (
	"io"
)

func Run(w io.Writer, path string) error {
	files, err := readDir(path)
	if err != nil {
		return err
	}

	err = printFiles(w, files)
	if err != nil {
		return err
	}

	return nil
}
