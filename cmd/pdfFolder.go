package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pdfFolderCmd represents the localize command
var pdfFolderCmd = &cobra.Command{
	Use:   "pdfFolder",
	Short: "Create the pdf folder at the same level of the tex file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("localize called")
	},
}

func init() {
	rootCmd.AddCommand(pdfFolderCmd)
}
