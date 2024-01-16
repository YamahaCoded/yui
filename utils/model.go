package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadModel(name string) error {
	url, path := GetData(name)

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Downloaded '%s' to use later\n", name)

	return nil
}

func AddModel(model, finalPath string) {
	_, file := GetData(model)
	Extract(file, finalPath)
}