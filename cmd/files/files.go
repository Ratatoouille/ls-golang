package files

import (
	"errors"
	"fmt"
	"io"
)

type FilesSlice []string

// AddItem add item to FilesSlice
func (f *FilesSlice) AddItem(filename string) {
	*f = append(*f, filename)
}

// PrintFiles print FilesSlice to io.Writer
func (f FilesSlice) PrintFiles(w io.Writer) error {
	for _, file := range f {
		_, err := fmt.Fprintf(w, "%v\n", file)
		if err != nil {
			return errors.New("print error")
		}
	}

	return nil
}
