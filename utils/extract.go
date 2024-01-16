package utils

import (
	"fmt"

	"github.com/gen2brain/go-unarr"
)

func Extract(filename, path string) {
	fmt.Println("Extracting file...")

	file, err := unarr.NewArchive(filename)
	if err != nil {
		panic(err)
	}

	file.Extract(path)
	defer file.Close()

}