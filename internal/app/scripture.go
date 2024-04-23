package app

import (
	"sync"
)

type ScriptureRef int
type ScriptureVersion string

type ScriptureWord struct {
	WordNumber int    `json:"word_num"`
	Text       string `json:"text"`
	Pre        string `json:"pre"`
	Post       string `json:"post"`
}

type ScriptureVerse struct {
	Ref   ScriptureRef    `json:"ref"`
	Words []ScriptureWord `json:"words"`
}

type ScriptureRange struct {
	Version ScriptureVersion `json:"version"`
	Start   ScriptureRef     `json:"start"`
	End     ScriptureRef     `json:"end"`
}

type ScriptureBlock struct {
	Range  ScriptureRange   `json:"range"`
	Verses []ScriptureVerse `json:"verses"`
}

type ScriptureSection struct {
	Range  ScriptureRange   `json:"range"`
	Blocks []ScriptureBlock `json:"blocks"`
}

func (a *App) GetScriptureSections(ran []ScriptureRange) []ScriptureSection {
	data := make([]ScriptureSection, 0, len(ran))
	for _, r := range ran {
		data = append(data, ScriptureSection{Range: r})
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(ran))
	for i := 0; i < len(ran); i++ {
		switch data[i].Range.Version {
		case "gnt":
			go getGNTScriptureSection(a, wg, &data[i])
		case "lxx":
			go getLXXScriptureSection(a, wg, &data[i])
		default:
			a.Throw("Unknown version (" + string(data[i].Range.Version) + ") passed to GetScriptureSections")
		}
	}
	wg.Wait()
	return data
}

type ScriptureWordData_Dictionary struct {
	Form  string `json:"form"`
	Gloss string `json:"gloss"`
}

type ScriptureWordData_Strongs struct {
	Num     string `json:"num"`
	Grammar string `json:"grammar"`
}

type ScriptureWordData struct {
	Ref        ScriptureRef                   `json:"ref"`
	WordNumber int                            `json:"word_number"`
	Text       string                         `json:"text"`
	Translit   string                         `json:"translit"`
	English    string                         `json:"english"`
	Dictionary []ScriptureWordData_Dictionary `json:"dictionary"`
	Strongs    []ScriptureWordData_Strongs    `json:"strongs"`
}

func (a *App) GetScriptureWord(version ScriptureVersion, ref ScriptureRef, word_num int) ScriptureWordData {
	w := ScriptureWordData{Ref: ref, WordNumber: word_num}

	switch version {
	case "gnt":
		getGNTScriptureWord(a, &w)
	case "lxx":
		getLXXScriptureWord(a, &w)
	default:
		a.Throw("Unknown version (" + string(version) + ") passed to GetScriptureSections")
	}

	return w
}
