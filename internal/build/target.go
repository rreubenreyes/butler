package build

import (
	"encoding/json"
	"errors"

	"github.com/rreubenreyes/butler/internal/build/nodejs"
)

type RequiredString string
type Runtime string

// Target defines a deployment artifact build target.
type Target struct {
	// Entry is an absolute path to the serverless function's entrypoint file.
	//
	// An "entrypoint file" is defined as the code which is directly imported and/or executed
	// at runtime by the host machine.
	Entry RequiredString `json:"entry"`

	// DryRun is a flag which determines whether or not to execute side effects.
	DryRun bool `json:"dry_run"`

	// Runtime defines the runtime which will be used for this build.
	Runtime Runtime `json:"runtime"`

	// BuildArtifactType determines type of output package of the generated build artifact.
	BuildArtifactType RequiredString `json:"build_artifact_type"`

	// NodeJS defines the set of parameters which are accepted by a NodeJS build target.
	NodeJS nodejs.Options `json:"nodejs"`
}

func (r *Runtime) UnmarshalJSON(data []byte) error {
	type Runtime2 Runtime
	var r2 Runtime2
	if err := json.Unmarshal(data, &r2); err != nil {
		return err
	}

	switch r2 {
	case "nodejs":
		*r = Runtime("nodejs")
		return nil
	default:
		return errors.New("invalid runtime")
	}
}

func (rs *RequiredString) UnmarshalJSON(data []byte) error {
	type RequiredString2 RequiredString
	var rs2 RequiredString2
	if err := json.Unmarshal(data, &rs2); err != nil {
		return err
	}

	*rs = RequiredString(rs2)
	return nil
}
