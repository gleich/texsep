package cmd

import (
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "texsep",
	Long: `
Separate your pdfs from your tex files
	            ___           ___            ___           ___           ___
                   /\__\         /|  |          /\__\         /\__\         /\  \
      ___         /:/ _/_       |:|  |         /:/ _/_       /:/ _/_       /::\  \
     /\__\       /:/ /\__\      |:|  |        /:/ /\  \     /:/ /\__\     /:/\:\__\
    /:/  /      /:/ /:/ _/_   __|:|__|       /:/ /::\  \   /:/ /:/ _/_   /:/ /:/  /
   /:/__/      /:/_/:/ /\__\ /::::\__\_____ /:/_/:/\:\__\ /:/_/:/ /\__\ /:/_/:/  /
  /::\  \      \:\/:/ /:/  / ~~~~\::::/___/ \:\/:/ /:/  / \:\/:/ /:/  / \:\/:/  /
 /:/\:\  \      \::/_/:/  /      |:|~~|      \::/ /:/  /   \::/_/:/  /   \::/__/
 \/__\:\  \      \:\/:/  /       |:|  |       \/_/:/  /     \:\/:/  /     \:\  \
      \:\__\      \::/  /        |:|__|         /:/  /       \::/  /       \:\__\
       \/__/       \/__/         |/__/          \/__/         \/__/         \/__/
	`,
}

// Execute the main command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		statuser.Error("Failed to execute root command", err, 1)
	}
}
