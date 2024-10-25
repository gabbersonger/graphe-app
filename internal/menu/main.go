package menu

import (
	"graphe/internal"
	"graphe/internal/logger"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type MenuManager struct {
	logger    *logger.Logger
	app       *application.App
	shortcuts *internal.SettingsValues_Shortcuts
}

func NewMenuManager(logger *logger.Logger) *MenuManager {
	return &MenuManager{
		logger:    logger,
		app:       nil,
		shortcuts: nil,
	}
}

func (m *MenuManager) assert(cond bool, msg string) {
	m.logger.Assert("MenuManager", cond, msg)
}

func (m *MenuManager) log(msg string) {
	m.logger.Log("MenuManager", msg)
}

func (m *MenuManager) SetApp(app *application.App) {
	m.assert(app != nil, "App is nil")
	m.app = app
	m.updateMenu()
}

func (m *MenuManager) SetShortcuts(shortcuts *internal.SettingsValues_Shortcuts) {
	m.shortcuts = shortcuts
	m.updateMenu()
}
