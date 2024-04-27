package app

import (
	"fmt"
	"sync"
)

func _getGNTScriptureWordText(a *App, wg *sync.WaitGroup, w *ScriptureWordData) {
	db := <-a.db.pool
	stmt := db.queries.GntWordText
	err := stmt.Bind(int(w.Ref), int(w.WordNumber))
	a.check(err)

	hasRow, err := stmt.Step()
	a.check(err)
	if !hasRow {
		a.Throw(fmt.Sprintf("Could not find value using query `GntWordText` for (ref=%d, word_num=%d)", int(w.Ref), w.WordNumber))
	}

	err = stmt.Scan(&(w.Text))
	a.check(err)

	stmt.Reset()
	a.db.pool <- db
	wg.Done()
}

func _getGNTScriptureWordBasicInfo(a *App, wg *sync.WaitGroup, w *ScriptureWordData) {
	db := <-a.db.pool
	stmt := db.queries.GntWordBasicInfo
	err := stmt.Bind(int(w.Ref), int(w.WordNumber))
	a.check(err)

	hasRow, err := stmt.Step()
	a.check(err)
	if !hasRow {
		a.Throw(fmt.Sprintf("Could not find value using query `GntWordBasicInfo` for (ref=%d, word_num=%d)", int(w.Ref), w.WordNumber))
	}

	err = stmt.Scan(&(w.Translit), &(w.English))
	a.check(err)

	stmt.Reset()
	a.db.pool <- db
	wg.Done()
}

func _getGNTScriptureWordDictionaryValues(a *App, wg *sync.WaitGroup, w *ScriptureWordData) {
	db := <-a.db.pool
	stmt := db.queries.GntWordDictionary
	err := stmt.Bind(int(w.Ref), int(w.WordNumber))
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

func _getGNTScriptureWordStrongsValues(a *App, wg *sync.WaitGroup, w *ScriptureWordData) {
	db := <-a.db.pool
	stmt := db.queries.GntWordStrongs
	err := stmt.Bind(int(w.Ref), int(w.WordNumber))
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

func _getGNTScriptureWordInflectedCount(a *App, wg *sync.WaitGroup, w *ScriptureWordData) {
	db := <-a.db.pool
	stmt := db.queries.GntWordInflectedCount
	err := stmt.Bind(int(w.Ref), int(w.WordNumber))
	a.check(err)

	hasRow, err := stmt.Step()
	a.check(err)
	if !hasRow {
		a.Throw(fmt.Sprintf("Could not find value using query `GntWordInflectedCount` for (ref=%d, word_num=%d)", int(w.Ref), w.WordNumber))
	}

	err = stmt.Scan(&(w.InflectedCount))
	a.check(err)

	stmt.Reset()
	a.db.pool <- db
	wg.Done()
}

func _getGNTScriptureWordLexemeCount(a *App, wg *sync.WaitGroup, w *ScriptureWordData) {
	db := <-a.db.pool
	stmt := db.queries.GntWordLexemeCount
	err := stmt.Bind(int(w.Ref), int(w.WordNumber))
	a.check(err)

	hasRow, err := stmt.Step()
	a.check(err)
	if !hasRow {
		a.Throw(fmt.Sprintf("Could not find value using query `GntWordLexemeCount` for (ref=%d, word_num=%d)", int(w.Ref), w.WordNumber))
	}

	err = stmt.Scan(&(w.LexemeCount))
	a.check(err)

	stmt.Reset()
	a.db.pool <- db
	wg.Done()
}

func getGNTScriptureWord(a *App, w *ScriptureWordData) {
	wg := new(sync.WaitGroup)
	wg.Add(6)
	go _getGNTScriptureWordText(a, wg, w)
	go _getGNTScriptureWordBasicInfo(a, wg, w)
	go _getGNTScriptureWordDictionaryValues(a, wg, w)
	go _getGNTScriptureWordStrongsValues(a, wg, w)
	go _getGNTScriptureWordInflectedCount(a, wg, w)
	go _getGNTScriptureWordLexemeCount(a, wg, w)
	wg.Wait()
}
