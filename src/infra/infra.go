package infra

import (
	"rms-backend/src/core"

	"github.com/rs/zerolog"
)

func Logger() zerolog.Logger {
	return core.CoreLogger().With().Str("layer", "infra").Logger()
}
