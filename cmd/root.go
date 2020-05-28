package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

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

// Execute ... Excute the main command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
