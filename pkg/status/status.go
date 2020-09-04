package status

import (
	"fmt"

	"github.com/fatih/color"
)

// Output a step
func Step(message string) {
	fmt.Println(message)
}

// Output a success
func Success(message string) {
	color.Green(message)
}
