package cmd

import (
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/spf13/cobra"
)

// foldersCmd represents the folders command
var foldersCmd = &cobra.Command{
	Use:   "folders",
	Short: "Move pdf and tex files into their own folders at the current level\n\t\tSee https://bit.ly/2XHmg5Q for more info",
	Run: func(cmd *cobra.Command, args []string) {
		statuser.ErrorMsg("This command is not yet implemented", 1)
	},
}

func init() {
	rootCmd.AddCommand(foldersCmd)
}
