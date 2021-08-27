package main

import (
	"fmt"
	"os"
	"unix-ls/cmd/ls"
)

func main() {
	if len(os.Args) == 1 {
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}

		err = ls.Run(os.Stdout, currentDir)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		path := os.Args[1]

		err := ls.Run(os.Stdout, path)
		if err != nil {
			fmt.Println(err)
		}
	}
}
