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
	Ref        ScriptureRef                   `json:"ref"`
	WordNumber int                            `json:"word_number"`
	Text       string                         `json:"text"`
	Translit   string                         `json:"translit"`
	English    string                         `json:"english"`
	Dictionary []ScriptureWordData_Dictionary `json:"dictionary"`
	Strongs    []ScriptureWordData_Strongs    `json:"strongs"`
}

func getGNTScriptureWordText(a *App, wg *sync.WaitGroup, w *ScriptureWordData) {
	db := <-a.db.pool

	stmt, err := db.getQuery("GetGNTScriptureWordText")
	a.check(err)

	err = stmt.Bind(int(w.Ref), int(w.WordNumber))
	a.check(err)

	hasRow, err := stmt.Step()
	a.check(err)
	if !hasRow {
		a.Throw(fmt.Sprintf("Could not find value using GetGNTScriptureWordText for (ref=%d, word_num=%d)", int(w.Ref), w.WordNumber))
	}

	err = stmt.Scan(&(w.Text))
	a.check(err)

	stmt.Reset()
	a.db.pool <- db
	wg.Done()
}

func getGNTScriptureWordBasicInfo(a *App, wg *sync.WaitGroup, w *ScriptureWordData) {
	db := <-a.db.pool

	stmt, err := db.getQuery("GetGNTScriptureWordBasicInfo")
	a.check(err)

	err = stmt.Bind(int(w.Ref), int(w.WordNumber))
	a.check(err)

	hasRow, err := stmt.Step()
	a.check(err)
	if !hasRow {
		a.Throw(fmt.Sprintf("Could not find value using GetGNTScriptureWordBasicInfo for (ref=%d, word_num=%d)", int(w.Ref), w.WordNumber))
	}

	err = stmt.Scan(&(w.Translit), &(w.English))
	a.check(err)

	stmt.Reset()
	a.db.pool <- db
	wg.Done()
}

func getGNTScriptureWordDictionaryValues(a *App, wg *sync.WaitGroup, w *ScriptureWordData) {
	db := <-a.db.pool

	stmt, err := db.getQuery("GetGNTScriptureWordDictionaryInfo")
	a.check(err)

	err = stmt.Bind(int(w.Ref), int(w.WordNumber))
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

func getGNTScriptureWordStrongsValues(a *App, wg *sync.WaitGroup, w *ScriptureWordData) {
	db := <-a.db.pool

	stmt, err := db.getQuery("GetGNTScriptureWordStrongsInfo")
	a.check(err)

	err = stmt.Bind(int(w.Ref), int(w.WordNumber))
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

func (a *App) GetScriptureWord(ref ScriptureRef, word_num int) ScriptureWordData {
	// TODO: remove
	if ref <= 40_000_000 {
		ref = 66020004
	}
	w := ScriptureWordData{Ref: ref, WordNumber: word_num}

	wg := new(sync.WaitGroup)
	wg.Add(4)
	go getGNTScriptureWordText(a, wg, &w)
	go getGNTScriptureWordBasicInfo(a, wg, &w)
	go getGNTScriptureWordDictionaryValues(a, wg, &w)
	go getGNTScriptureWordStrongsValues(a, wg, &w)
	wg.Wait()
	return w
}
