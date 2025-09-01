package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Welcome message
	fmt.Println("Welcome to your first Go program!")
	fmt.Print("What is your name? ")

	// Read user input
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading input", err)
		return
	}

	fmt.Printf("What is your favorite language? ")
	language, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("An error occurred", err)
		return
	}

	// Greet the user
	name = strings.TrimSpace(name)
	language = strings.TrimSpace(language)
	fmt.Printf("Hello, %v! Soon, %v will be your favorite language over JavaScript/TypeScript.\n", name, language)

	fmt.Printf("Do you want to continue learning Go? (yes/no): ")
	answer, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error occurred while reading input", err)
		return
	}

	// Control flow
	answer = strings.TrimSpace(answer)
	if strings.ToLower(answer) == "yes" {
		fmt.Printf("That is right %v! Let's keep going.\n", name)
	} else {
		fmt.Printf("Unfortunately, you will come back again, %v.\n", name)
	}

}
