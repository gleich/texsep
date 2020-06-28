package cmd

import (
	"github.com/Matt-Gleich/texsep/check"
	"github.com/Matt-Gleich/texsep/sort"
	"github.com/spf13/cobra"
)

// mirrorCmd represents the mirror command
var mirrorCmd = &cobra.Command{
	Use:   "mirror",
	Short: "Create a mirror strcutrue for the PDFs\n\t\tSee https://bit.ly/2X8B5iH for more info",
	Run: func(cmd *cobra.Command, args []string) {
		check.ProjectRoot()
		files := sort.Files()
		sort.MoveFiles(files)
	},
}

func init() {
	rootCmd.AddCommand(mirrorCmd)
}
