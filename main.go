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
	cmdArg := os.Args[2]
	name := os.Args[3]

	switch cmd {
	case "setup":
		utils.Setup(user, repo, substringAsset)
		
	case "help":
		utils.HelpMessage()

	case "update":
		utils.MingwUpdate(user, repo, substringAsset)
	
	case "create":
		switch cmdArg {
		case "project":
			utils.CreateProject(name)
		}

	case "delete":
		switch cmdArg{
		case "project":
			utils.DeleteProject(name)
		}

	default:
		fmt.Println("Unknown command:", cmd)
		os.Exit(1)
	}

}