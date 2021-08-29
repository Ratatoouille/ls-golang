package ls

import (
	"fmt"
	"io"
)

func Run(w io.Writer, path []string) error {
	for _, p := range path {
		files, err := readDir(p)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintf(w, "%v:\n", p)
		if err != nil {
			return err
		}
		err = files.PrintFiles(w)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(w, "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
