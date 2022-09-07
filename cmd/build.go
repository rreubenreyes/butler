package cmd

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/rreubenreyes/butler/internal/build"
	"github.com/rreubenreyes/butler/internal/log"
	"github.com/spf13/cobra"
)

// base build command
var buildManifestPath string
var target = &build.Target{}

var cmdBuild = &cobra.Command{
	Use:       "build",
	Short:     "Build a serverless function",
	Long:      "Build a serverless function",
	ValidArgs: []string{"nodejs"},
	Args:      cobra.ExactValidArgs(1),
}

func initCmdBuild() {
	cmdBuild.PersistentFlags().
		BoolVar(&target.DryRun, "dry-run", false, "if set, no build side effects will occur")
	cmdBuild.PersistentFlags().
		StringVarP(&buildManifestPath, "manifest", "m", "", "path to manifest file")

	// cmdBuild.MarkFlagRequired("build-artifact-type")
	rootCmd.AddCommand(cmdBuild)
}

var cmdBuildNodeJS = &cobra.Command{
	Use:   "nodejs [entrypoint]",
	Short: "Build a NodeJS serverless function",
	Long:  "Build a NodeJS serverless function",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, logger := log.ForContextWith(context.Background(), "cmd", "build/nodejs")
		logger.Trace().Msg("starting nodejs command")

		if buildManifestPath != "" {
			// TODO: might have to resolve absolute path from manifest here
			target = build.MustBindTargetFromManifest(ctx, buildManifestPath)
			if target.Runtime != "nodejs" {
				panic(errors.New("invalid nodejs manifest"))
			}
		} else {
			target.Entry = args[0]
			target.Runtime = "nodejs"
		}

		t, _ := json.Marshal(target)

		logger.Debug().RawJSON("target", t).Msg("bound build arguments")
	},
}

func initCmdBuildNodeJS() {
	cmdBuildNodeJS.Flags().
		StringVarP(&target.BuildArtifactType, "build-artifact-type", "t", "", "the type of build artifact to generate")
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

	cmdBuildNodeJS.MarkFlagRequired("build-artifact-type")
	cmdBuild.AddCommand(cmdBuildNodeJS)
}
