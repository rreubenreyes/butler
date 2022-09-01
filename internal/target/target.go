package target

// Target defines a deployment artifact build target.
type Target struct {
	// Entrypoint is an absolute path to the serverless function's entrypoint file.
	//
	// An "entrypoint file" is defined as the code which is directly imported and/or executed
	// at runtime by the host machine.
	Entrypoint string
}
