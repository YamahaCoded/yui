package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CreateDir(directoryName string) {
	fmt.Printf("Creating directory %s...\n", directoryName)

	path := filepath.Join("C:\\", directoryName)

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		fmt.Printf("'%s' created\n", directoryName)
	} else if err != nil {
		fmt.Println("Something went wrong:", err)
	} else {
		fmt.Printf("'%s' already exists\n", directoryName)
	}
}

func CopyFile(src, dest string) {
	sourceFile, err := os.Open(src)
	if err != nil {
		fmt.Printf("Error due to: %s", err)
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dest)
	if err != nil {
		fmt.Printf("Error due to: %s\n", err)
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		fmt.Printf("Error due to: %s\n", err)
	}
}