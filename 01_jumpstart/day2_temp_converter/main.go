package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Day 2 - Temperature Converter")
	fmt.Print("Enter temperature: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading input:", err)
		return
	}

	input = strings.TrimSpace(input)
	temp, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal("Invalid numbers:", err)
		return
	}

	fmt.Print("Convert to (C)elsius or (F)ahrenheit? ")
	unit, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading input:", err)
		return
	}

	unit = strings.TrimSpace(strings.ToUpper(unit))
	switch unit {
	case "C":
		// F -> C
		celsius := (temp - 32) * 5 / 9
		fmt.Printf("%.2f°F is %.2f°C\n", temp, celsius)
	case "F":
		// C -> F
		fahrenheit := (temp * 9 / 5) + 32
		fmt.Printf("%.2f°C is %.2f°F\n", temp, fahrenheit)
	default:
		fmt.Println("Invalid choice! Please enter C or F.")
	}

}
