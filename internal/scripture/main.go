package scripture

import (
	"graphe/internal/logger"
)

type ScriptureService struct {
	logger *logger.Logger
}

func NewScriptureService(logger *logger.Logger) *ScriptureService {
	return &ScriptureService{
		logger: logger,
	}
}

func (s *ScriptureService) assert(cond bool, msg string) {
	s.logger.Assert("ScriptureService", cond, msg)
}
