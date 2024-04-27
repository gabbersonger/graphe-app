package app

import (
	"fmt"
	"sync"
)

func _getLXXScriptureWordText(a *App, wg *sync.WaitGroup, w *ScriptureWordData) {
	db := <-a.db.pool
	stmt := db.queries.LxxWordText
	err := stmt.Bind(int(w.Ref), int(w.WordNumber))
	a.check(err)

	hasRow, err := stmt.Step()
	a.check(err)
	if !hasRow {
		a.Throw(fmt.Sprintf("Could not find value using query `LxxWordText` for (ref=%d, word_num=%d)", int(w.Ref), w.WordNumber))
	}

	err = stmt.Scan(&(w.Text))
	a.check(err)

	stmt.Reset()
	a.db.pool <- db
	wg.Done()
}

func _getLXXScriptureWordInfo(a *App, wg *sync.WaitGroup, w *ScriptureWordData) {
	db := <-a.db.pool
	stmt := db.queries.LxxWordBasicInfo
	err := stmt.Bind(int(w.Ref), int(w.WordNumber))
	a.check(err)

	hasRow, err := stmt.Step()
	a.check(err)
	if !hasRow {
		a.Throw(fmt.Sprintf("Could not find value using query `LxxWordBasicInfo` for (ref=%d, word_num=%d)", int(w.Ref), w.WordNumber))
	}

	w.Strongs = make([]ScriptureWordData_Strongs, 1)
	w.Dictionary = make([]ScriptureWordData_Dictionary, 1)

	err = stmt.Scan(&(w.Translit), &(w.English), &(w.Strongs[0].Num), &(w.Strongs[0].Grammar), &(w.Dictionary[0].Form), &(w.Dictionary[0].Gloss))
	a.check(err)

	stmt.Reset()
	a.db.pool <- db
	wg.Done()
}

func _getLXXScriptureWordInflectedCount(a *App, wg *sync.WaitGroup, w *ScriptureWordData) {
	db := <-a.db.pool
	stmt := db.queries.LxxWordInflectedCount
	err := stmt.Bind(int(w.Ref), int(w.WordNumber))
	a.check(err)

	hasRow, err := stmt.Step()
	a.check(err)
	if !hasRow {
		a.Throw(fmt.Sprintf("Could not find value using query `LxxWordInflectedCount` for (ref=%d, word_num=%d)", int(w.Ref), w.WordNumber))
	}

	err = stmt.Scan(&(w.InflectedCount))
	a.check(err)

	stmt.Reset()
	a.db.pool <- db
	wg.Done()
}

func _getLXXScriptureWordLexemeCount(a *App, wg *sync.WaitGroup, w *ScriptureWordData) {
	db := <-a.db.pool
	stmt := db.queries.LxxWordLexemeCount
	err := stmt.Bind(int(w.Ref), int(w.WordNumber))
	a.check(err)

	hasRow, err := stmt.Step()
	a.check(err)
	if !hasRow {
		a.Throw(fmt.Sprintf("Could not find value using GetLXXScriptureWordLexemeCount for (ref=%d, word_num=%d)", int(w.Ref), w.WordNumber))
	}

	err = stmt.Scan(&(w.LexemeCount))
	a.check(err)

	stmt.Reset()
	a.db.pool <- db
	wg.Done()
}

func getLXXScriptureWord(a *App, w *ScriptureWordData) {
	wg := new(sync.WaitGroup)
	wg.Add(4)
	go _getLXXScriptureWordText(a, wg, w)
	go _getLXXScriptureWordInfo(a, wg, w)
	go _getLXXScriptureWordInflectedCount(a, wg, w)
	go _getLXXScriptureWordLexemeCount(a, wg, w)
	wg.Wait()
}
