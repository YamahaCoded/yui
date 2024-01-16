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
		fmt.Println("Usage: yui <command>\nTry using help for more information")
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
	
	case "create":
		name := os.Args[2]
		utils.CreateProject(name)
	
	case "delete":
		name := os.Args[2]
		utils.DeleteProject(name)
		
	case "library":
		extraArgs := os.Args[2]
		switch extraArgs {
		case "add":
			name := os.Args[3]
			url := os.Args[4]
			utils.InsertData(name, url)
		case "remove":
			name := os.Args[3]
			utils.RemoveData(name)
		}

	default:
		fmt.Println("Unknown command:", cmd)
		os.Exit(1)
	}

}