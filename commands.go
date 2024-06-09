package main

import "fmt"

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
	for _, dir := range paths {
		path += dir.Name
		if dir.Name != "/" {
			path += "/"
		}
	}
	fmt.Println(path)
}

func changeDir(path string) {
	if path == ".." {
		length := len(paths)
		currentDir = paths[length-2]
		paths = paths[0 : len(paths)-1]
		return
	}

	dir, ok := currentDir.Directories[path]

	if !ok {
		fmt.Println("Dir doesn't exit")
		return
	}
	paths = append(paths, dir)
	currentDir = dir
}
