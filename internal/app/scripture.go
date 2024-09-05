package app

import (
	data "graphe/internal/data"
	"graphe/internal/scripture"
)

func (a *App) GetScriptureSection(r scripture.ScriptureRange) []data.ScriptureSection {
	return a.data.GetScriptureSection(r)
}

func (a *App) GetScriptureWord(v scripture.ScriptureVersion, r scripture.ScriptureRef, w int) data.ScriptureWordData {
	return a.data.GetScriptureWord(v, r, w)
}
