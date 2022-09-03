package nodejs

import (
	"context"

	"github.com/rreubenreyes/butler/internal/log"
)

// Target defines a deployment artifact build target.
type Options struct {
	// ProjectRoot is an _optional_ absolute path to  the project root to be used
	// when building the serverless function defined by this Target. If not specified,
	// then we attempt to detect the project root at build time.
	ProjectRoot string `json:"project_root"`

	// RemoveUnusedImports is a flag which defines whether or not to remove unused imported code
	// from the final build artifact. If set, modules which are defined in this Target's
	// dependency file and/or lock file, but are not used in the built code, are removed from
	// the dependency file and/or lock file.
	RemoveUnusedImports bool `json:"remove_unused_imports"`

	// AdditionalFiles is an array of absolute paths to additional files which should be built
	// alongside the Target's entrypoint. All files specified in this array must be children
	// of the Target's ProjectRoot.
	AdditionalFiles []string `json:"additional_files"`

	// IgnoreLockFile is a flag which defines whether or not to ignore this project's lock file
	// when building. By default, the lock file is always used. The dependency file will instead
	// be used to install modules.
	IgnoreLockFile bool `json:"ignore_lock_file"`

	// PreserveFileStructure is a flag which defines whether or not to preserve the source
	// file structure when creating the build artifact.
	PreserveFileStructure bool `json:"preserve_file_structure"`
}

func Build(entry string, dryRun bool, opts *Options) error {
	ctx, logger := log.ForContext(context.Background())

	b := &Builder{
		ctx:    ctx,
		entry:  entry,
		dryRun: dryRun,
		opts:   opts,
	}

	err := b.Build()
	if err != nil {
		logger.Error().
			Str("err", err.Error()).
			Msg("error building")

		return err
	}

	logger.Info().Msg("finished building")

	return nil
}
