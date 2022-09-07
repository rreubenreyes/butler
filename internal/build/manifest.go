package build

import (
	"context"
	"encoding/json"
	"os"

	"github.com/rreubenreyes/butler/internal/log"
)

func MustBindTargetFromManifest(ctx context.Context, path string) *Target {
	_, logger := log.ForContextWith(ctx, "manifest_path", path)

	logger.Debug().Msg("reading manifest")
	data, err := os.ReadFile(path)
	if err != nil {
		logger.Error().Msg(err.Error())
		panic(err)
	}

	var t *Target
	err = json.Unmarshal(data, t)
	if err != nil {
		logger.Error().Msg(err.Error())
		panic(err)
	}

	logger.Debug().RawJSON("manifest", data).Msg("read manifest")

	return t
}
