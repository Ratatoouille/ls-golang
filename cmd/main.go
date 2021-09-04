package main

import (
	"log"
	"os"
	"unix-ls/cmd/ls"
)

func main() {
	if len(os.Args) >= 2 {
		path := os.Args[2:]

		err := ls.Run(os.Stdout, path)
		if err != nil {
			log.Println(err)
		}
		return
	}
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	err = ls.Run(os.Stdout, []string{currentDir})
	if err != nil {
		log.Println(err)
	}
}
