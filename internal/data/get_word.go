package data

import (
	"fmt"
	"graphe/internal/scripture"
)

type ScriptureWordDataField struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

type ScriptureWordData struct {
	Version    scripture.ScriptureVersion `json:"version"`
	Ref        scripture.ScriptureRef     `json:"ref"`
	WordNumber int                        `json:"word_number"`
	Text       string                     `json:"text"`
	Fields     []ScriptureWordDataField   `json:"fields"`
}

func (d *DataDB) GetScriptureWord(version scripture.ScriptureVersion, ref scripture.ScriptureRef, word_num int) ScriptureWordData {
	d.assert(d.scripture_service.IsVersionValid(version), fmt.Sprintf("Invalid ref (version: `%s`, ref: %d, word_num: %d)", version, int(ref), word_num))
	d.assert(d.scripture_service.IsRefValid(ref, version), fmt.Sprintf("Invalid ref (version: `%s`, ref: %d, word_num: %d)", version, int(ref), word_num))
	d.assert(word_num > 0, fmt.Sprintf("Invalid word number (version: `%s`, ref: %d, word_num: %d)", version, int(ref), word_num))

	wd := ScriptureWordData{Version: version, Ref: ref, WordNumber: word_num}
	switch version {
	case "gnt":
		d.getScriptureWordGnt(&wd)
	case "lxx":
		d.getScriptureWordLxx(&wd)
	case "esv":
		d.getScriptureWordEsv(&wd)
	}
	return wd
}
