package app

import "graphe/internal/settings"

func (a *App) GetSettings() settings.SettingsValues {
	return a.settings.GetSettings()
}

func (a *App) UpdateSetting(field []string, value interface{}) bool {
	return a.settings.UpdateSetting(field, value)
}
