package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: program_name <content> <filename>")
		return
	}

	if strings.HasSuffix(os.Args[2], ".txt") {
		file, err := os.OpenFile(os.Args[2], os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			return
		}
		defer file.Close()

		if _, err := file.WriteString(os.Args[1]); err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			return
		}

		fmt.Printf("File %s has been generated.\n", os.Args[2])
	} else {
		if os.Args[2] == "--upper" {
			n ,_ := os.ReadFile(os.Args[1])
			n2 := string(n)

			fileContent := strings.ToUpper(n2)
			if err := os.WriteFile(os.Args[1], []byte(fileContent), 0644); err != nil {
				fmt.Printf("Error writing to file: %v\n", err)
				return
			}
			fmt.Printf("File %s has been modified to uppercase.\n", os.Args[1])
		} else {
			fmt.Println("Invalid operation.")
		}
	}
}

