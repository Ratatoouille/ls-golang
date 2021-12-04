package ls

import (
	"io/fs"

	"github.com/fatih/color"
)

func print(dirEntry []fs.DirEntry) {
	for _, v := range dirEntry {
		if v.Name()[0:1] != "." {
			if v.IsDir() {
				color.Blue("%v\t", v.Name())
			} else {
				color.White("%v\t", v.Name())
			}
		}
	}
}

func printAll(dirEntry []fs.DirEntry) {
	for _, v := range dirEntry {
		if v.IsDir() {
			color.Blue("%v\t", v.Name())
		} else {
			color.White("%v\t", v.Name())
		}
	}
}
