package cmd

import (
	"context"
	"encoding/json"

	"github.com/rreubenreyes/butler/internal/build"
	"github.com/rreubenreyes/butler/internal/log"
	"github.com/spf13/cobra"
)

// base build command
var target = &build.Target{}

var cmdBuild = &cobra.Command{
	Use:       "build",
	Short:     "Build a serverless function",
	Long:      "Build a serverless function",
	ValidArgs: []string{"nodejs"},
	Args:      cobra.ExactValidArgs(1),
}

func initCmdBuild() {
	cmdBuildNodeJS.PersistentFlags().
		StringVarP(&target.BuildArtifactType, "build-artifact-type", "t", "", "the type of build artifact to generate")
	cmdBuild.PersistentFlags().
		BoolVar(&target.DryRun, "dry-run", false, "if set, no build side effects will occur")

	cmdBuild.MarkFlagRequired("build-artifact-type")
	rootCmd.AddCommand(cmdBuild)
}

var cmdBuildNodeJS = &cobra.Command{
	Use:   "nodejs [entrypoint]",
	Short: "Build a NodeJS serverless function",
	Long:  "Build a NodeJS serverless function",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := log.ForContextWith(context.Background(), "cmd", "build/nodejs")
		logger := log.FromContext(ctx)
		logger.Trace().Msg("starting nodejs command")

		target.Entry = args[0]
		target.Runtime = "nodejs"
		t, _ := json.Marshal(target)

		logger.Debug().RawJSON("target", t).Msg("binding build arguments")
	},
}

func initCmdBuildNodeJS() {
	cmdBuildNodeJS.Flags().
		StringVarP(&target.NodeJS.ProjectRoot, "project-root", "r", "", "location of project root")
	cmdBuildNodeJS.Flags().
		BoolVar(&target.NodeJS.RemoveUnusedImports, "remove-unused-imports", true, "if set, removes unused imports from build artifact")
	cmdBuildNodeJS.Flags().
		BoolVar(&target.NodeJS.IgnoreLockFile, "ignore-lock-file", false, "if set, installs dependencies from package.json instead of package-lock.json")
	cmdBuildNodeJS.Flags().
		BoolVar(&target.NodeJS.PreserveFileStructure, "preserve-file-structure", false, "if set, preserves original project file structure")
	cmdBuildNodeJS.Flags().
		StringSliceVar(&target.NodeJS.AdditionalFiles, "additional-files", []string{}, "if set, removes unused imports from build artifact")

	cmdBuild.AddCommand(cmdBuildNodeJS)
}
