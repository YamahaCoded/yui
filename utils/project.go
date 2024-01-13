package utils

import (
	"os"
	"path/filepath"
)

func CreateProject(name string) {
	docs, err := os.UserHomeDir()
	if err != nil {
		return
	}

	path := filepath.Join(docs, "Documents\\Projects\\")
	finalPath := filepath.Join(path, name)

	CreateDir(finalPath)
}

func DeleteProject() {

}
