package scripture

import (
	"fmt"
	"log/slog"
)

type ScriptureService struct {
	logger *slog.Logger
}

func NewScriptureService(logger *slog.Logger) *ScriptureService {
	return &ScriptureService{
		logger: logger,
	}
}

func (s *ScriptureService) assert(cond bool, msg string) {
	if !cond {
		if s.logger != nil {
			s.logger.Error(fmt.Sprintf("[ScriptureService] %s", msg))
		} else {
			panic(msg)
		}
	}
}
