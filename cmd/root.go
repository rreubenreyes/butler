package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:       "butler",
	Short:     "serverless build tool",
	Long:      "serverless build tool",
	ValidArgs: []string{"build"},
	Args:      cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("running root cmd")
	},
}

func Execute() {
	initCmdBuild()
	initCmdBuildNodeJS()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
