package cmd

import (
	"time"

	"github.com/Matt-Gleich/logoru"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/Matt-Gleich/texsep/pkg/check"
	"github.com/Matt-Gleich/texsep/pkg/sort"
	"github.com/Matt-Gleich/texsep/pkg/status"
	"github.com/spf13/cobra"
)

// mirrorCmd represents the mirror command
var mirrorCmd = &cobra.Command{
	Use:   "mirror",
	Short: "Create a mirror strcutrue for the PDFs\n\t\tSee https://bit.ly/2X8B5iH for more info",
	Run: func(cmd *cobra.Command, args []string) {
		loop, err := cmd.Flags().GetBool("loop")
		if err != nil {
			statuser.Error("Failed to get loop flag", err, 1)
		}
		check.ProjectRoot()
		if loop {
			for {
				files := sort.Files()
				logoru.Success("Successfully got all latex and pdf files")
				sort.MoveFiles(files)
				time.Sleep(10 * time.Second)
			}
		} else {
			files := sort.Files()
			status.Success("Successfully got all latex and pdf files")
			sort.MoveFiles(files)
		}
	},
}

func init() {
	rootCmd.AddCommand(mirrorCmd)
	mirrorCmd.Flags().BoolP("loop", "l", false, "Loop the program every 10 seconds")
}
