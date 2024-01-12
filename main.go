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
		fmt.Println("Usage: yui <command>")
		os.Exit(1)
	}

	cmd := os.Args[1]
	
	switch cmd {
	case "setup":
		utils.Setup(user, repo, substringAsset)
		
	case "help":
		utils.HelpMessage()

	case "update":
		utils.MingwUpdate(user, repo, substringAsset)

	default:
		fmt.Println("Unknown command:", cmd)
		os.Exit(1)
	}

}