package app

import "graphe/internal/settings"

func (a *App) GetSettings() *settings.SettingsValues {
	return a.settings.GetSettings()
}

func (a *App) UpdateSetting(group string, field string, value string) bool {
	return a.settings.UpdateSetting(group, field, value)
}
