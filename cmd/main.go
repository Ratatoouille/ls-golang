package main

import (
	"log"
	"os"
	"unix-ls/pkg/ls"
)

func main() {
	var path []string

	if len(os.Args) >= 2 {
		path = append(path, os.Args[1:]...)
	} else {
		cDir, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		path = append(path, cDir)
	}

	err := ls.Ls(path)
	if err != nil {
		log.Println(err)
	}
}
