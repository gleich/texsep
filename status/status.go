package status

import (
	"fmt"

	"github.com/fatih/color"
)

// Step ... Output a step
func Step(message string) {
	fmt.Println(message)
}

// Success ... Output a success
func Success(message string) {
	color.Green(message)
}
