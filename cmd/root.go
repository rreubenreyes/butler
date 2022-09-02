package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "butler",
	Short: "serverless build tool",
	Long:  "serverless build tool",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var cmdBuild = &cobra.Command{
	Use:   "build [OPTIONS]",
	Short: "Build a serverless function",
	Long:  "Build a serverless function",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("running build command with args: " + strings.Join(args, " "))
	},
}

func initCommands() {
	// build
	cmdBuild.Flags()
	rootCmd.AddCommand(cmdBuild)
}

func Execute() {
	initCommands()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
