package utils

import (
	"fmt"
	"io"
	"log"
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

func UseModel(model, finalPath string) {
	_, file := GetData(model)
	Extract(file, finalPath)
}

func AddTasks(name, url string) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if !DataExists(db, name){
		InsertData(name, url)
		DownloadModel(name)
	} else {
		fmt.Printf("'%s' already exists\n", name)
	}

}

func DeleteTasks(name string) {
	_, file := GetData(name)

	if IsDirectory(file) {
		RemoveData(name)
		DeleteDir(file)
	} else {
		fmt.Println(file, IsDirectory(file))
		fmt.Printf("%s not found\n", file)
	}
}