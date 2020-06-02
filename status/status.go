package status

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// Error ... Output an error
func Error(err error, message string) {
	fmt.Println(err)
	color.Red(message)
	os.Exit(0)
}

// Step ... Output a step
func Step(message string) {
	fmt.Println(message)
}

// Success ... Output a success
func Success(message string) {
	color.Green(message)
}
