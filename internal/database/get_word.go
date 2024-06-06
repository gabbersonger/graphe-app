package database

import (
	"fmt"
	"graphe/internal/scripture"
)

// type ScriptureWordDataFields map[string]string
type ScriptureWordDataField struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type ScriptureWordData struct {
	Version     scripture.ScriptureVersion `json:"version"`
	Ref         scripture.ScriptureRef     `json:"ref"`
	WordNumber  int                        `json:"word_number"`
	Text        string                     `json:"text"`
	Fields      []ScriptureWordDataField   `json:"fields"`
	Collections [][]ScriptureWordDataField `json:"collections"`
}

func (g *GrapheDB) GetScriptureWord(v scripture.ScriptureVersion, r scripture.ScriptureRef, w int) ScriptureWordData {
	d := ScriptureWordData{Version: v, Ref: r, WordNumber: w}
	switch v {
	case "gnt":
		getScriptureWord_GNT(g, &d)
	case "esv":
		getScriptureWord_ESV(g, &d)
	case "lxx":
		getScriptureWord_LXX(g, &d)
	default:
		g.throw(fmt.Sprintf("Invalid version in GetScriptureWord. args=(%s, %d, %d)", v, r, w))
	}
	return d
}
