package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func CheckPath(substr string) {
	path := os.Getenv("Path")
	check := strings.Contains(path, substr)
	
	if !check {
		AddToPath(substr)
	} else {
		fmt.Printf("'%s' already on path", substr)
	}
}

func AddToPath(path string) {
    newPath := `[Environment]::SetEnvironmentVariable("Path",[Environment]::GetEnvironmentVariable("Path", [EnvironmentVariableTarget]::User) + ";`+ path +`", [EnvironmentVariableTarget]::User)`
	cmd := exec.Command("powershell", newPath)
	
	err := cmd.Run()
	if err != nil {
		fmt.Println("Something went wrong:", err)
		return
	}

	fmt.Printf("Successfully added '%s' to Path\n", path)
}
