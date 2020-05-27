package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// mirrorCmd represents the mirror command
var mirrorCmd = &cobra.Command{
	Use:   "mirror",
	Short: "Create a mirror strcutrue for the PDFs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mirror called")
	},
}
