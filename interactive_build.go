package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	root := "." // Set the root directory

	for {
		printMenu()
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			buildAllFiles(root)
		case "2":
			fmt.Println("Enter the directory path:")
			scanner.Scan()
			dirPath := scanner.Text()
			buildAllFiles(dirPath)
		case "3":
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		fmt.Println()
	}
}

func printMenu() {
	fmt.Println("Menu:")
	fmt.Println("1. Build all Go files in the current directory")
	fmt.Println("2. Build all Go files in a specific directory")
	fmt.Println("3. Exit")
	fmt.Print("Enter your choice: ")
}

func buildAllFiles(dirPath string) {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %s: %v\n", path, err)
			return nil
		}

		if !info.IsDir() && filepath.Ext(path) == ".go" {
			// Execute 'go build' for each Go file
			cmd := exec.Command("go", "build", path)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			if err != nil {
				fmt.Printf("Error building file %s: %v\n", path, err)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %s: %v\n", dirPath, err)
	}
}
