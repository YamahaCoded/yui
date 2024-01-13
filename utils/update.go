package utils

import "fmt"

func MingwUpdate(user, repo, substringAsset string) {
	_, finalAsset, err := GetAssetDetails(user, repo, substringAsset)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	DeleteDir("C:\\Yui")
	
	CreateDir("C:\\Yui")
	CreateDir("C:\\Yui\\Files")
	CreateDir("C:\\Yui\\Compiler")

	DownloadAsset(user, repo, substringAsset)
	Extract(finalAsset, "C:\\Yui\\Compiler")
	
	CopyFile("./yui.exe", "C:\\Yui\\yui.exe")
	CopyFile("C:\\Yui\\Compiler\\mingw64\\bin\\mingw32-make.exe", "C:\\Yui\\Compiler\\mingw64\\bin\\make.exe")
}
