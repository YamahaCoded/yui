package utils

import (
	"os"
	"path/filepath"
)

func CreateProject(name, model string) {
	docs, err := os.UserHomeDir()
	if err != nil {
		return
	}

	path := filepath.Join(docs, "Documents\\Projects\\")
	
	if _, err := os.Stat(path); os.IsNotExist(err) {
		CreateDir(path)
	}

	finalPath := filepath.Join(path, name)
	CreateDir(finalPath)
	UseModel(model, finalPath)
}

func DeleteProject(name string) {
	docs, err := os.UserHomeDir()
	if err != nil {
		return
	}

	path := filepath.Join(docs, "Documents\\Projects\\")
	
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return
	}

	finalPath := filepath.Join(path, name)
	DeleteDir(finalPath)
}
