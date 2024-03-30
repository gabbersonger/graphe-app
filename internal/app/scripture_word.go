package app

import (
	"fmt"
	"sync"
)

type ScriptureWordData_Dictionary struct {
	Form  string `json:"form"`
	Gloss string `json:"gloss"`
}

type ScriptureWordData_Strongs struct {
	Num     string `json:"num"`
	Grammar string `json:"grammar"`
}

type ScriptureWordData struct {
	ref         ScriptureRef
	word_num    int
	Translit    string                         `json:"translit"`
	English     string                         `json:"english"`
	ConjoinWord string                         `json:"conjoin_word"`
	SubMeaning  string                         `json:"sub_meaning"`
	Dictionary  []ScriptureWordData_Dictionary `json:"dictionary"`
	Strongs     []ScriptureWordData_Strongs    `json:"strongs"`
}

func getScriptureWordBasicInfo(a *App, wg *sync.WaitGroup, w *ScriptureWordData) {
	db := <-a.db.pool

	stmt, err := db.getQuery("GetScriptureWordBasicInfo")
	a.check(err)

	err = stmt.Bind(int(w.ref), int(w.word_num))
	a.check(err)

	hasRow, err := stmt.Step()
	a.check(err)
	if !hasRow {
		a.Throw(fmt.Sprintf("Could not find value using GetScriptureWordBasicInfo for (ref=%d, word_num=%d)", int(w.ref), w.word_num))
	}

	err = stmt.Scan(&(w.Translit), &(w.English), &(w.ConjoinWord), &(w.SubMeaning))
	a.check(err)

	stmt.Reset()
	a.db.pool <- db
	wg.Done()
}

func getScriptureWordDictionaryValues(a *App, wg *sync.WaitGroup, w *ScriptureWordData) {
	db := <-a.db.pool

	stmt, err := db.getQuery("GetScriptureWordDictionaryInfo")
	a.check(err)

	err = stmt.Bind(int(w.ref), int(w.word_num))
	a.check(err)

	w.Dictionary = make([]ScriptureWordData_Dictionary, 0, 1)
	var form, gloss string
	for {
		hasRow, err := stmt.Step()
		a.check(err)
		if !hasRow {
			break
		}

		err = stmt.Scan(&form, &gloss)
		a.check(err)

		w.Dictionary = append(w.Dictionary, ScriptureWordData_Dictionary{Form: form, Gloss: gloss})
	}

	stmt.Reset()
	a.db.pool <- db
	wg.Done()
}

func getScriptureWordStrongsValues(a *App, wg *sync.WaitGroup, w *ScriptureWordData) {
	db := <-a.db.pool

	stmt, err := db.getQuery("GetScriptureWordStrongsInfo")
	a.check(err)

	err = stmt.Bind(int(w.ref), int(w.word_num))
	a.check(err)

	w.Strongs = make([]ScriptureWordData_Strongs, 0, 1)
	var num, grammar string
	for {
		hasRow, err := stmt.Step()
		a.check(err)
		if !hasRow {
			break
		}

		err = stmt.Scan(&num, &grammar)
		a.check(err)

		w.Strongs = append(w.Strongs, ScriptureWordData_Strongs{Num: num, Grammar: grammar})
	}

	stmt.Reset()
	a.db.pool <- db
	wg.Done()
}

func (a *App) GetScriptureWord(ref ScriptureRef, word int) ScriptureWordData {
	w := ScriptureWordData{ref: ref, word_num: word}

	wg := new(sync.WaitGroup)
	wg.Add(3)
	go getScriptureWordBasicInfo(a, wg, &w)
	go getScriptureWordDictionaryValues(a, wg, &w)
	go getScriptureWordStrongsValues(a, wg, &w)
	wg.Wait()
	return w
}
