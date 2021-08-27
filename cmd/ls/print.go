package ls

import (
	"errors"
	"fmt"
	"io"
)

func printFiles(w io.Writer, files []string) error {
	for _, file := range files {
		_, err := fmt.Fprintf(w, "%v\n", file)
		if err != nil {
			return errors.New("print error")
		}
	}

	return nil
}
