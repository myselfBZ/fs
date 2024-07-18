package main

import (
	"fmt"
	"strings"
)

func showEntries() {
	for name := range currentDir.Directories {
		fmt.Println(name + "/")
	}

	for name := range currentDir.Files {
		fmt.Println(name)
	}

}

func pwd() {
	var path string
	for _, dir := range paths.dirs {
		path += dir.Name
		if dir.Name != "/" {
			path += "/"
		}
	}
	fmt.Println(path)
}

func changeDir(path string) {
    newPath := strings.Split(path, "/")
    for _, path := range newPath{
        if path == ".."{
            err := paths.Pop() 
            if err != nil{
                fmt.Println("You are in the root dir")
            }
        } else{
            dir, ok := currentDir.Directories[path]
            if !ok{
                fmt.Println(DirDoesntExist)
                return 
            } 
            currentDir = dir
            
        }
    }
}

func echo(filename, text string) {
	file, ok := currentDir.Files[filename]
	if !ok {
		fmt.Println("File doesn't exist")
		return
	}
	file.Content = fmt.Sprintf("%s\n%s", file.Content, text)
}

func cat(filename string) {
	file, ok := currentDir.Files[filename]
	if !ok {
		fmt.Println("file doesn't exist")
		return
	}
	fmt.Println(file.Content)
}
