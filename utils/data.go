package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Data struct {
	Name string
	URL  string
	Path string
}

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "C:\\Yui\\Files\\libraries.db")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS data (
            name TEXT PRIMARY KEY,
            url TEXT,
            path TEXT
        )
    `)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InsertData(name, url string) {

	path := "C:\\Yui\\Files\\" + name
	CreateDir(path)

	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO data VALUES (?, ?, ?)", name, url, path)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: data.name" {
			fmt.Printf("%s Already exists\n", name)
		} else {
			log.Fatal(err)
		}
		return
	}

	fmt.Printf("Data inserted for %s\n", name)
}


func GetData(name string) (string, string) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if !DataExists(db, name) {
		fmt.Printf("No data found for %s\n", name)
		return "", ""
	}

	row := db.QueryRow("SELECT url, path FROM data WHERE name = ?", name)

	var retrievedURL, retrievedPath string
	err = row.Scan(&retrievedURL, &retrievedPath)
	if err != nil {
		log.Fatal(err)
		return "", ""
	}

	return retrievedURL, retrievedPath
}

func ModifyURL(name, newURL string) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if !DataExists(db, name) {
		fmt.Printf("No data found for %s\n", name)
		return
	}

	_, err = db.Exec("UPDATE data SET url = ? WHERE name = ?", newURL, name)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("URL modified for %s to %s\n", name, newURL)
}

func RemoveData(name string) {

	path := "C:\\Yui\\Files\\" + name
	DeleteDir(path)

	db, err := ConnectDB()
	if err != nil {
		log.Fatal("erro:", err)
	}
	defer db.Close()

	if !DataExists(db, name) {
		fmt.Printf("No data found for %s\n", name)
		return
	}

	_, err = db.Exec("DELETE FROM data WHERE name = ?", name)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Data removed for %s\n", name)
}

func DataExists(db *sql.DB, name string) bool {
	row := db.QueryRow("SELECT name FROM data WHERE name = ?", name)

	var retrievedName string
	err := row.Scan(&retrievedName)
	if err != nil {
		return false
	}

	return true
}