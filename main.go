package main

import (
	"fmt"
	"os"
	"yamaha/yui/utils"
)

func main() {
	user := "niXman"
	repo := "mingw-builds-binaries"
	substringAsset := "win32-seh-msvcrt"

	if len(os.Args) < 2 {
		fmt.Println("Usage: yui setup -user=<username> -repo=<repository> -substring=<substring> -destiny=<directory> [-confirm]")
		os.Exit(1)
	}

	cmd := os.Args[1]
	switch cmd {
	case "setup":
		latestVersion := utils.LastVersion(user, repo)
		fmt.Println("Latest Version:", latestVersion)

		fmt.Print("Do you want to proceed with the download? (y/n): ")
		var response string
		fmt.Scanln(&response)
		if response != "y" {
			fmt.Println("Download canceled.")
			os.Exit(0)
		} else {
			utils.DownloadAsset(user, repo, substringAsset)
		}

	default:
		fmt.Println("Unknown command:", cmd)
		os.Exit(1)
	}
}
