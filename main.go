package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// FahrenheitToCelsius converts a temperature from Fahrenheit to Celsius
func FahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

// CelsiusToFahrenheit converts a temperature from Celsius to Fahrenheit
func CelsiusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter temperature (e.g., 32 F or 100 C) or type 'exit' to quit:")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Exit the program if the user types 'exit'
		if strings.ToLower(input) == "exit" {
			break
		}

		// Split the input into temperature and unit
		parts := strings.Fields(input)
		if len(parts) != 2 {
			fmt.Println("Invalid input. Please enter a temperature followed by a unit (C or F).")
			continue
		}

		// Parse the temperature
		temp, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			fmt.Println("Invalid temperature. Please enter a numeric value.")
			continue
		}

		// Get the unit
		unit := strings.ToUpper(parts[1])

		// Perform the conversion based on the unit
		switch unit {
		case "C":
			// Convert Celsius to Fahrenheit
			fahrenheit := CelsiusToFahrenheit(temp)
			fmt.Printf("%.2f C is %.2f F\n", temp, fahrenheit)
		case "F":
			// Convert Fahrenheit to Celsius
			celsius := FahrenheitToCelsius(temp)
			fmt.Printf("%.2f F is %.2f C\n", temp, celsius)
		default:
			fmt.Println("Invalid unit. Please enter either 'C' for Celsius or 'F' for Fahrenheit.")
		}
	}
}
