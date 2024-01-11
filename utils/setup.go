package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/schollz/progressbar/v3"
)

func Tasks(finalAsset string) {
	Extract(finalAsset, "C:\\Yui\\Compiler")
	CheckPath("C:\\Yui")
	CheckPath("C:\\Yui\\Compiler\\mingw64\\bin")
	CopyFile("./yui.exe", "C:\\Yui\\yui.exe")
	CopyFile("C:\\Yui\\Compiler\\mingw64\\bin\\mingw32-make.exe", "C:\\Yui\\Compiler\\mingw64\\bin\\make.exe")
}

func LastVersion(user, repo string) string {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", user, repo)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	var result map[string]interface{}
	if err := decodeJSON(res.Body, &result); err != nil {
		fmt.Println(err)
		return ""
	}

	return result["tag_name"].(string)
}

func assetExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func DownloadAsset(user, repo, substringAsset string) {
	latestVersion := LastVersion(user, repo)
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/tags/%s", user, repo, latestVersion)

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	var release map[string]interface{}
	if err := decodeJSON(res.Body, &release); err != nil {
		fmt.Println(err)
		return
	}

	var assetURL, finalAsset string
	assets := release["assets"].([]interface{})
	for _, a := range assets {
		asset := a.(map[string]interface{})
		if strings.Contains(asset["name"].(string), substringAsset) {
			assetURL = asset["browser_download_url"].(string)
			finalAsset = path.Join("./", asset["name"].(string))
			break
		}
	}

	if assetURL != "" {
		if assetExists(finalAsset) {
			fmt.Printf("%s already downloaded, skipping download...\n", finalAsset)
			Tasks(finalAsset)
			return
		}

		response, err := http.Get(assetURL)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer response.Body.Close()

		file, err := os.Create(finalAsset)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		bar := progressbar.DefaultBytes(
			response.ContentLength,
			"Downloading",
		)

		_, err = io.Copy(io.MultiWriter(file, bar), response.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("\nDownloaded %s to: %s\n", finalAsset, finalAsset)
		Tasks(finalAsset)
	} else {
		fmt.Printf("No assets with \"%s\".\n", substringAsset)
	}

}