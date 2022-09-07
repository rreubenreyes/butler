package build

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestMustBuildTargetFromManifest_Happy(t *testing.T) {
	var target Target
	manifest := []byte(strings.TrimSpace(`
  {
    "entry": "test",
    "dry_run": false,
    "runtime": "nodejs",
    "build_artifact_type": "test"
  }
  `))

	err := json.Unmarshal(manifest, &target)
	if err != nil {
		t.Errorf("could not unmarshal expected valid manifest: " + err.Error())
	}
}

func TestMustBuildTargetFromManifest_InvalidRuntime(t *testing.T) {
	var target Target
	manifest := []byte(strings.TrimSpace(`
  {
    "entry": "test",
    "dry_run": false,
    "runtime": "invalid",
    "build_artifact_type": "test"
  }
  `))

	err := json.Unmarshal(manifest, &target)
	if err == nil {
		t.Errorf("did not catch expected invalid manifest: " + err.Error())
	}
}
