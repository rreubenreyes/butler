package nodejs

// Target defines a deployment artifact build target.
type Target struct {
  // Entry is an absolute path to the serverless function's entrypoint file.
  //
  // An "entrypoint file" is defined as the code which is directly imported and/or executed
  // at runtime by the host machine.
  Entry string

  // ProjectRoot is an _optional_ absolute path to  the project root to be used
  // when building the serverless function defined by this Target. If not specified,
  // then we attempt to detect the project root at build time.
  ProjectRoot string

  // RemoveUnusedImports is a flag which defines whether or not to remove unused imported code
  // from the final build artifact. If set, modules which are defined in this Target's
  // dependency file and/or lock file, but are not used in the built code, are removed from
  // the dependency file and/or lock file.
  RemoveUnusedImports bool

  // AdditionalFiles is an array of absolute paths to additional files which should be built
  // alongside the Target's entrypoint. All files specified in this array must be children
  // of the Target's ProjectRoot.
  AdditionalFiles []string

  // IgnoreLockFile is a flag which defines whether or not to ignore this project's lock file
  // when building. By default, the lock file is always used. The dependency file will instead
  // be used to install modules.
  IgnoreLockFile bool
}

func discoverProjectRoot() error {
  // TODO: implement
}

func discoverUsedImports() error {

}

func removeUnusedImports(ignoreLockFile bool) error {

}

func buildFromLockFile() error {
  // TODO: implement
}

func buildFromDependencyFile() error {
  // TODO: implement
}

func (t *Target) Build() error {
  if !t.IgnoreLockFile {
    // build from package.json
    return buildFromLockFile()
  }
  // build from package-lock.json
  return buildFromDependencyFile()
}
