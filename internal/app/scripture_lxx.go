package app

import (
	"fmt"
	"sync"
)

func getLXXScriptureSection(a *App, wg *sync.WaitGroup, s *ScriptureSection) {
	s.Blocks = make([]ScriptureBlock, 0, 1)

	db := <-a.db.pool
	stmt, err := db.getQuery("GetLXXScriptureSection")
	a.check(err)
	err = stmt.Bind(int(s.Range.Start), int(s.Range.End))
	a.check(err)

	var ref, word_num int
	var text, pre, post string
	createNextBlock := true
	lastRef := 0

	for {
		hasRow, err := stmt.Step()
		a.check(err)
		if !hasRow {
			break
		}
		err = stmt.Scan(&ref, &word_num, &text, &pre, &post)
		a.check(err)

		// Add block if needed
		if createNextBlock {
			createNextBlock = false
			if len(s.Blocks) > 0 {
				s.Blocks[len(s.Blocks)-1].Range.End = ScriptureRef(lastRef)
			}

			newBlock := ScriptureBlock{}
			newBlock.Range.Version = s.Range.Version
			newBlock.Range.Start = ScriptureRef(ref)
			newBlock.Verses = make([]ScriptureVerse, 0, 20) // TODO: pick the right value
			s.Blocks = append(s.Blocks, newBlock)
		}
		lastBlock := len(s.Blocks) - 1

		// Add verse if needed
		if lastRef != ref || len(s.Blocks[lastBlock].Verses) == 0 {
			lastRef = ref

			newVerse := ScriptureVerse{}
			newVerse.Ref = ScriptureRef(ref)
			newVerse.Words = make([]ScriptureWord, 0, 40) // TODO: pick the right value
			s.Blocks[lastBlock].Verses = append(s.Blocks[lastBlock].Verses, newVerse)
		}
		lastVerse := len(s.Blocks[lastBlock].Verses) - 1

		// Check for paragraph break
		n := -1
		runes := []rune(post)
		for i, rune := range runes {
			if rune == '¶' {
				createNextBlock = true
				n = i
				break
			}
		}
		if n >= 0 {
			runes = append(runes[:n], runes[n+1:]...)
			post = string(runes)
		}

		// Add word
		newWord := ScriptureWord{word_num, text, pre, post}
		s.Blocks[lastBlock].Verses[lastVerse].Words = append(s.Blocks[lastBlock].Verses[lastVerse].Words, newWord)
	}
	s.Range.End = ScriptureRef(ref)
	s.Blocks[len(s.Blocks)-1].Range.End = ScriptureRef(ref)

	stmt.Reset()
	a.db.pool <- db
	wg.Done()
}

func _getLXXScriptureWordText(a *App, wg *sync.WaitGroup, w *ScriptureWordData) {
	db := <-a.db.pool

	stmt, err := db.getQuery("GetLXXScriptureWordText")
	a.check(err)

	err = stmt.Bind(int(w.Ref), int(w.WordNumber))
	a.check(err)

	hasRow, err := stmt.Step()
	a.check(err)
	if !hasRow {
		a.Throw(fmt.Sprintf("Could not find value using GetLXXScriptureWordText for (ref=%d, word_num=%d)", int(w.Ref), w.WordNumber))
	}

	err = stmt.Scan(&(w.Text))
	a.check(err)

	stmt.Reset()
	a.db.pool <- db
	wg.Done()
}

func _getLXXScriptureWordInfo(a *App, wg *sync.WaitGroup, w *ScriptureWordData) {
	db := <-a.db.pool

	stmt, err := db.getQuery("GetLXXScriptureWordBasicInfo")
	a.check(err)

	err = stmt.Bind(int(w.Ref), int(w.WordNumber))
	a.check(err)

	hasRow, err := stmt.Step()
	a.check(err)
	if !hasRow {
		a.Throw(fmt.Sprintf("Could not find value using GetLXXScriptureWordBasicInfo for (ref=%d, word_num=%d)", int(w.Ref), w.WordNumber))
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

	stmt, err := db.getQuery("GetLXXScriptureWordInflectedCount")
	a.check(err)

	err = stmt.Bind(int(w.Ref), int(w.WordNumber))
	a.check(err)

	hasRow, err := stmt.Step()
	a.check(err)
	if !hasRow {
		a.Throw(fmt.Sprintf("Could not find value using GetLXXScriptureWordInflectedCount for (ref=%d, word_num=%d)", int(w.Ref), w.WordNumber))
	}

	err = stmt.Scan(&(w.InflectedCount))
	a.check(err)

	stmt.Reset()
	a.db.pool <- db
	wg.Done()
}

func _getLXXScriptureWordLexemeCount(a *App, wg *sync.WaitGroup, w *ScriptureWordData) {
	db := <-a.db.pool

	stmt, err := db.getQuery("GetLXXScriptureWordLexemeCount")
	a.check(err)

	err = stmt.Bind(int(w.Ref), int(w.WordNumber))
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
