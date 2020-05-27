package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// foldersCmd represents the folders command
var foldersCmd = &cobra.Command{
	Use:   "folders",
	Short: "Move pdf and tex files into their own folders at the current level",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("folders called")
	},
}
