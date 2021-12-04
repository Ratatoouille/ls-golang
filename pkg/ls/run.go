package ls

import (
	"flag"
	"fmt"
	"os"
)

var all *bool

func Ls(path []string) error {
	all = flag.Bool("a", false, "display all files")
	flag.Parse()

	for _, p := range path {
		if p == "-a" {
			continue
		}
		dirEntry, err := os.ReadDir(p)
		if err != nil {
			return err
		}

		if len(path) > 1 {
			fmt.Printf("%v:\n", p)
		}

		if *all {
			printAll(dirEntry)
		} else {
			print(dirEntry)
		}

		fmt.Printf("\n\n")
	}

	return nil
}
