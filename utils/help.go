package utils

import "fmt"

const helpString = `Yui is a command line tool to easily setup Mingw and manage C/C++ project models.

Usage:
	yui <command> [arguments]

Commands:
	setup                        - Installs the latest version of MinGW.
	update                       - Updates MinGW and Yui.
	create [name] [model]        - Creates a new project using a model.
	delete [name]                - Deletes the project.
	model add [name] [url]       - Creates a model with a name and URL.
	model remove [name]          - Deletes an existing model.
	help                         - Displays this message.
	version                      - Displays current version of Yui.

Examples:
	yui setup
	yui update
	yui model add myModel https://your.file//your_model.7z
	yui model remove myModel
	yui create myProject myModel
	yui delete myProject

`

func HelpMessage(){
	fmt.Printf(helpString)
	
}