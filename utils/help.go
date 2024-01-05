package utils

import "fmt"

const helpString = `Yui is a command line tool to easily setup Mingw and manage C/C++ libraries.

Usage:
	yui <command> [arguments]

Commands:
	setup                        - Installs the latest version of mingw.
	update                       - Updates mingw and the yui utility.
	create [name] [library]      - Creates a new project using a library.
	delete [name]                - Deletes the project.
	library add [name] [url]     - Creates a library with a name and URL.
	library remove [name]        - Deletes an existing library.
	help                         - Displays this message.
	version                      - Displays current version of yui and mingw.

Examples:
	yui setup
	yui update
	yui library add myLibrary https://github.com/example/libraryTemplate.git
	yui library remove myLibrary
	yui create myProject myLibrary
	yui delete myProject

`

func HelpMessage(){
	fmt.Printf("%s", helpString)
}