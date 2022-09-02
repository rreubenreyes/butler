package nodejs

import (
	"context"

	"github.com/rreubenreyes/butler/internal/log"
)

// Target defines a deployment artifact build target.
type Target struct {
	// Entry is an absolute path to the serverless function's entrypoint file.
	//
	// An "entrypoint file" is defined as the code which is directly imported and/or executed
	// at runtime by the host machine.
	Entry string `json:"entry"`

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

	// BuildArtifactType defines what build artifact should be expected from this build.
	BuildArtifactType string `json:"build_artifact_type"`

	// DryRun is a flag which determines whether or not to execute side effects.
	DryRun bool `json:"dry_run"`
}

func (t *Target) Build() error {
	ctx := log.ForContext(context.Background())

  b := &Builder{ctx: ctx, target: t}
	b.Build()
}
