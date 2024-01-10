package utils

import (
	"fmt"
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
