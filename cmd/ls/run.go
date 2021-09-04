package ls

import (
	"flag"
	"fmt"
	"io"
)

func Run(w io.Writer, path []string) error {
	boolPtr := flag.Bool("all", false, "print all files")
	flag.Parse()

	for _, p := range path {
		files, err := readDir(p, *boolPtr)
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
