package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

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

		_, err = io.Copy(file, response.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Downloaded %s to: %s\n", finalAsset, finalAsset)
	} else {
		fmt.Printf("No assets with \"%s\".\n", substringAsset)
	}
}
