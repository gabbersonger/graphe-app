package app

import (
	"graphe/internal/database"
	"graphe/internal/scripture"
)

func (a *App) GetScriptureSection(r scripture.ScriptureRange) []database.ScriptureSection {
	return a.db.GetScriptureSection(r)
}

func (a *App) GetScriptureWord(v scripture.ScriptureVersion, r scripture.ScriptureRef, w int) database.ScriptureWordData {
	return a.db.GetScriptureWord(v, r, w)
}
