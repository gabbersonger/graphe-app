package app

import "sync"

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
			a.Throw("Unknown version passed to GetScriptureSections")
		}
	}
	wg.Wait()
	return data
}
