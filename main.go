package main

import (
	"fmt"
	"os"
	"yamaha/yui/utils"
)

func checkArgs(n int) {
	if len(os.Args) < n {
		fmt.Println("Usage: yui <command> [arguments]\nTry using help for more information")
		os.Exit(1)
	}
}

func main() {

	user := "niXman"
	repo := "mingw-builds-binaries"
	substringAsset := "win32-seh-msvcrt"

	checkArgs(2)

	cmd := os.Args[1]

	switch cmd {
		
	case "setup":
		utils.Setup(user, repo, substringAsset)
		
	case "help":
		utils.HelpMessage()

	case "update":
		utils.MingwUpdate(user, repo, substringAsset)
	
	case "create":
		checkArgs(2)
		name := os.Args[2]
		model := os.Args[3]
		utils.CreateProject(name, model)
	
	case "delete":
		checkArgs(2)
		name := os.Args[2]
		utils.DeleteProject(name)
		
	case "model":
		checkArgs(4)
		extraArgs := os.Args[2]
		switch extraArgs {
		case "add":
			name := os.Args[3]
			url := os.Args[4]
			utils.AddTasks(name, url)
		case "remove":
			checkArgs(3)
			name := os.Args[3]
			utils.DeleteTasks(name)
		}

	case "version":
		utils.PrintVersion()

	default:
		fmt.Println("Unknown command:", cmd)
		os.Exit(1)
	}

}