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

	// NodeJS defines the set of parameters which are accepted by a NodeJS build target.
	NodeJS nodejs.Options
}

func Build(t *Target) error {
	ctx := log.ForContext(context.Background())
	logger := log.FromContext(ctx)

	return nil
}
