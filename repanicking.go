package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// processInput simulates processing of user input
func processInput(input string) {
	defer func() {
		if r := recover(); r != nil {
			// Log the panic for debugging purposes
			log.Printf("Recovered in processInput: %v", r)
			// Perform any necessary cleanup
			fmt.Println("Cleanup after panic...")
			// Re-panic to ensure the calling function
			// can handle it or fail gracefully
			panic(r)
		}
	}()

	// Simulate a panic scenario based on input
	// If the input is not a number, panic
	if _, err := strconv.ParseInt(input, 10, 64); err != nil {
		panic(err)
	}

	fmt.Println("Processed input successfully:", input)
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Main recovered from panic:", r)
		}
	}()

	// Get a number from the user
	fmt.Print("Enter a number: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read from standard input:", err)
		return
	}
	input = strings.Trim(input, "\n")

	fmt.Println("Processing input:", input)
	processInput(input)

	fmt.Println("Main function continues after processing input.")
}
