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