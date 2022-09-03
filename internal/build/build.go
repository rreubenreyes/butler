package build

import (
	"context"

	"github.com/rreubenreyes/butler/internal/build/nodejs"
	"github.com/rreubenreyes/butler/internal/log"
)

type Runtime string

const (
	NodeJS Runtime = "nodejs"
)

// Target defines a deployment artifact build target.
type Target struct {
	// Entry is an absolute path to the serverless function's entrypoint file.
	//
	// An "entrypoint file" is defined as the code which is directly imported and/or executed
	// at runtime by the host machine.
	Entry string `json:"entry"`

	// DryRun is a flag which determines whether or not to execute side effects.
	DryRun bool `json:"dry_run"`

	// Runtime defines the runtime which will be used for this build.
	Runtime Runtime `json:"runtime"`

	// BuildArtifactType determines type of output package of the generated build artifact.
	BuildArtifactType string `json:"build_artifact_type"`

	// NodeJS defines the set of parameters which are accepted by a NodeJS build target.
	NodeJS nodejs.Options `json:"node_js"`
}

func Build(ctx context.Context, t *Target) error {
	_, logger := log.ForContext(ctx)
	logger.Trace().Msg("starting top-level build")

	return nil
}
