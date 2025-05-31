package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Open the file for reading
	messages, fileOpenError := os.Open("messages.txt")
	if fileOpenError != nil {
		fmt.Printf("failed to open file: %v\n", fileOpenError)
		return
	}
	defer messages.Close()

	// Create a buffer to hold 8 bytes as you read the file
	buffer := make([]byte, 8)

	// Read the file
	for {
		offset, fileReadError := messages.Read(buffer)
		if fileReadError != nil {
			if fileReadError == io.EOF {
				break
			}
			fmt.Printf("Error reading file: %v\n", fileReadError)
			return
		}

		if offset > 0 {
			fmt.Printf("read: %s\n", buffer[:offset])
		}
	}
}
