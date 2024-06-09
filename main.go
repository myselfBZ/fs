package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/myselfBZ/fs/entries"
)

var (
	root       = entries.RootDir()
	currentDir = root
	paths      = []*entries.Directory{root}
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		input, err := reader.ReadString('\n')
		if err != nil {
			panic("I don't know how it can happen")
		}
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}
		tokens := strings.Split(input, " ")
		argc := len(tokens)
		switch tokens[0] {
		case "mkdir":
			if argc == 2 {
				currentDir.Create(tokens[1])
			} else if argc < 2 {
				fmt.Println("Not enough arguments")
			} else {
				fmt.Println("too many arguments, required only 2")
			}
		case "ls":
			showEntries()
		case "cd":
			if argc != 2 {
				fmt.Println("Usage: cd name of the dir")
			} else {
				changeDir(tokens[1])
			}
		case "pwd":
			pwd()
		case "touch":
			if argc != 2 {
				fmt.Println("Provdie the name of a single file")
			} else {
				currentDir.CreateFile(tokens[1])
			}
		case "echo":
			echo(tokens[1], tokens[2])
		case "cat":
			cat(tokens[1])
		default:
			fmt.Println("Command not found")
		}

	}
}
