package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check if the correct number of arguments is provided
	if len(os.Args[1:]) < 2 || len(os.Args[1:]) > 3 {
		fmt.Println("Not enough arguments")
		return
	}

	// If the second argument ends with ".txt", create a new file with the provided content
	if strings.HasSuffix(os.Args[2], ".txt") {
		err := os.WriteFile(os.Args[2], []byte(os.Args[1]), 0o644)
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			return
		}
		fmt.Printf("File %s has been generated.\n", os.Args[2])
	} else {
		// Handle other operations (e.g., --upper, --cap, --low)
		if os.Args[2] == "--upper" {
			n, _ := os.ReadFile(os.Args[1])
			n2 := string(n)
			fileContent := strings.ToUpper(n2)
			if err := os.WriteFile(os.Args[1], []byte(fileContent), 0644); err != nil {
				fmt.Printf("Error writing to file: %v\n", err)
				return
			}
			fmt.Printf("File %s has been modified to uppercase.\n", os.Args[1])
		} else if os.Args[2] == "--cap" {
			n, _ := os.ReadFile(os.Args[1])
			n2 := string(n)
			fileContent := strings.Title(n2)
			if err := os.WriteFile(os.Args[1], []byte(fileContent), 0644); err != nil {
				fmt.Printf("Error writing to file: %v\n", err)
				return
			}
			fmt.Printf("File %s has been modified to title case.\n", os.Args[1])
		} else if os.Args[2] == "--low" {
			n, _ := os.ReadFile(os.Args[1])
			n2 := string(n)
			fileContent := strings.ToLower(n2)
			if err := os.WriteFile(os.Args[1], []byte(fileContent), 0644); err != nil {
				fmt.Printf("Error writing to file: %v\n", err)
				return
			}
			fmt.Printf("File %s has been modified to lowercase.\n", os.Args[1])
		} else {
			fmt.Println("Invalid operation.")
		}
	}
}
