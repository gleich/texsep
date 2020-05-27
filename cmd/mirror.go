package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// mirrorCmd represents the mirror command
var mirrorCmd = &cobra.Command{
	Use:   "mirror",
	Short: "Create a mirror strcutrue for the PDFs\n\t\tSee https://bit.ly/2X8B5iH for more info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mirror called")
	},
}

func init() {
	rootCmd.AddCommand(mirrorCmd)
}
