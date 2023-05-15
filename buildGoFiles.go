package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
)

func main() {
    // Get the current working directory
    currentDir, err := os.Getwd()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Print the menu to the console
    fmt.Println("1. Build all .go files")
    fmt.Println("2. Build all .go files in the current directory")
    fmt.Println("3. Build a specific .go file")
    fmt.Println("4. Exit")

    // Read user input from the console
    reader := bufio.NewReader(os.Stdin)
    choice, _, err := reader.ReadLine()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    switch string(choice) {
    case "1":
        // Recursively go through all directories from the current location
        err = filepath.Walk(currentDir, func(path string, info os.FileInfo, err error) error {
            // Check if the file is a regular file and has a ".go" extension
            if err == nil && !info.IsDir() && filepath.Ext(path) == ".go" {
                // Build the file with the go command
                cmd := exec.Command("go", "build", path)
                output, err := cmd.CombinedOutput()
                if err != nil {
                    fmt.Println("Error:", err)
                    return err
                }
                fmt.Printf("Built %s:\n%s\n", path, output)
            }
            return nil
        })

        if err != nil {
            fmt.Println("Error:", err)
            return
        }

    case "2":
        // Compile all ".go" files in the root directory as well
        rootFiles, err := filepath.Glob("*.go")
        if err != nil {
            fmt.Println("Error:", err)
            return
        }

        for _, file := range rootFiles {
            // Build the file with the go command
            cmd := exec.Command("go", "build", file)
            output, err := cmd.CombinedOutput()
            if err != nil {
                fmt.Println("Error:", err)
                return
            }
            fmt.Printf("Built %s:\n%s\n", file, output)
        }

    case "3":
        // Read the path of the file to build from the user
        fmt.Print("Enter the path of the .go file: ")
        path, _, err := reader.ReadLine()
        if err != nil {
            fmt.Println("Error:", err)
            return
        }

        // Build the file with the go command
        cmd := exec.Command("go", "build", string(path))
        output, err := cmd.CombinedOutput()
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        fmt.Printf("Built %s:\n%s\n", string(path), output)

    case "4":
        // Exit the program
        return

    default:
        fmt.Println("Invalid choice")
        return
    }
}