package utils

import (
	"fmt"

	"github.com/gen2brain/go-unarr"
)

func Extract(filename, path string) {
	fmt.Println("Extracting file...")

	file, err := unarr.NewArchive("./x86_64-13.2.0-release-win32-seh-msvcrt-rt_v11-rev0.7z")
	if err != nil {
		panic(err)
	}

	file.Extract(path)
	defer file.Close()
	
}