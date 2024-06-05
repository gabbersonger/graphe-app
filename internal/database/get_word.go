package database

import (
	"fmt"
	"graphe/internal/scripture"
)

type ScriptureWordDataFields map[string]string

type ScriptureWordData struct {
	Ref         scripture.ScriptureRef    `json:"ref"`
	WordNumber  int                       `json:"word_number"`
	Text        string                    `json:"text"`
	Fields      ScriptureWordDataFields   `json:"fields"`
	Collections []ScriptureWordDataFields `json:"collections"`
}

func (g *GrapheDB) GetScriptureWord(v scripture.ScriptureVersion, r scripture.ScriptureRef, w int) ScriptureWordData {
	d := ScriptureWordData{Ref: r, WordNumber: w}
	d.Fields = make(ScriptureWordDataFields)
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
