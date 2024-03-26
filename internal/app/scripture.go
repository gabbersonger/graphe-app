package app

import (
	"errors"
	"fmt"
	"sync"
)

type ScriptureRef int
type ScriptureWord string
type ScriptureVersion string

type ScriptureVerse struct {
	Ref   ScriptureRef    `json:"ref"`
	Words []ScriptureWord `json:"words"`
}

type ScriptureRange struct {
	Start ScriptureRef `json:"start"`
	End   ScriptureRef `json:"end"`
}

type ScriptureBlock struct {
	Range  ScriptureRange   `json:"range"`
	Verses []ScriptureVerse `json:"verses"`
}

type ScriptureSection struct {
	Version ScriptureVersion `json:"version"`
	Range   ScriptureRange   `json:"range"`
	Blocks  []ScriptureBlock `json:"blocks"`
}

func (a *App) getScriptureSection(out chan *ScriptureSection, ver ScriptureVersion, ran ScriptureRange) {
	s := ScriptureSection{Version: ver, Range: ran}
	s.Blocks = make([]ScriptureBlock, 0, 1)

	db := <-a.db_pool
	stmt, err := db.getQuery("GetScriptureSection")
	check(a.ctx, err)
	err = stmt.Bind(int(ran.Start), int(ran.End))
	check(a.ctx, err)

	var ref, word_num int
	var text string
	createNextBlock := true
	lastRef := 0

	for {
		hasRow, err := stmt.Step()
		check(a.ctx, err)
		if !hasRow {
			break
		}
		err = stmt.Scan(&ref, &word_num, &text)
		check(a.ctx, err)

		// Add block if needed
		if createNextBlock {
			createNextBlock = false
			if len(s.Blocks) > 0 {
				s.Blocks[len(s.Blocks)-1].Range.End = ScriptureRef(lastRef)
			}

			newBlock := ScriptureBlock{}
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

		// Add word
		newWord := ScriptureWord(text)
		s.Blocks[lastBlock].Verses[lastVerse].Words = append(s.Blocks[lastBlock].Verses[lastVerse].Words, newWord)

		runes := []rune(text)
		max := len(runes) - 1
		for i := range runes {
			if runes[max-i] == '¶' {
				createNextBlock = true
			}
		}
	}
	s.Range.End = ScriptureRef(ref)
	s.Blocks[len(s.Blocks)-1].Range.End = ScriptureRef(ref)

	stmt.Reset()
	a.db_pool <- db
	out <- &s
}

func (a *App) GetScriptureSections(ver ScriptureVersion, ran []ScriptureRange) []ScriptureSection {
	feed := make(chan *ScriptureSection, len(ran))
	for _, r := range ran {
		go a.getScriptureSection(feed, ver, r)
	}

	data := make([]ScriptureSection, 0, len(ran))
	for i := 0; i < len(ran); i++ {
		data = append(data, *(<-feed))
	}
	return data
}

type ScriptureWordValue_Dictionary struct {
	Form  string `json:"form"`
	Gloss string `json:"gloss"`
}

type ScriptureWordValue_Strongs struct {
	Num     string `json:"num"`
	Grammar string `json:"grammar"`
}

type ScriptureWordValue struct {
	Translit    string                          `json:"translit"`
	English     string                          `json:"english"`
	ConjoinWord string                          `json:"conjoin_word"`
	SubMeaning  string                          `json:"sub_meaning"`
	Dictionary  []ScriptureWordValue_Dictionary `json:"dictionary"`
	Strongs     []ScriptureWordValue_Strongs    `json:"strongs"`
}

func (a *App) getScriptureWordBasicInfo(wg *sync.WaitGroup, w *ScriptureWordValue, ref ScriptureRef, word int) {
	db := <-a.db_pool

	stmt, err := db.getQuery("GetScriptureWordBasicInfo")
	check(a.ctx, err)

	err = stmt.Bind(int(ref), int(word))
	check(a.ctx, err)

	hasRow, err := stmt.Step()
	check(a.ctx, err)
	if !hasRow {
		check(a.ctx, errors.New(fmt.Sprintf("Could not find value using GetScriptureWordBasicInfo for (ref=%d, word_num=%d)", ref, word)))
	}

	err = stmt.Scan(&(w.Translit), &(w.English), &(w.ConjoinWord), &(w.SubMeaning))
	check(a.ctx, err)

	stmt.Reset()
	a.db_pool <- db
	wg.Done()
}

func (a *App) getScriptureWordDictionaryValues(wg *sync.WaitGroup, w *ScriptureWordValue, ref ScriptureRef, word int) {
	db := <-a.db_pool

	stmt, err := db.getQuery("GetScriptureWordDictionaryInfo")
	check(a.ctx, err)

	err = stmt.Bind(int(ref), int(word))
	check(a.ctx, err)

	w.Dictionary = make([]ScriptureWordValue_Dictionary, 0, 1)
	var form, gloss string
	for {
		hasRow, err := stmt.Step()
		check(a.ctx, err)
		if !hasRow {
			break
		}

		err = stmt.Scan(&form, &gloss)
		check(a.ctx, err)

		w.Dictionary = append(w.Dictionary, ScriptureWordValue_Dictionary{Form: form, Gloss: gloss})
	}

	stmt.Reset()
	a.db_pool <- db
	wg.Done()
}

func (a *App) getScriptureWordStrongsValues(wg *sync.WaitGroup, w *ScriptureWordValue, ref ScriptureRef, word int) {
	db := <-a.db_pool

	stmt, err := db.getQuery("GetScriptureWordStrongsInfo")
	check(a.ctx, err)

	err = stmt.Bind(int(ref), int(word))
	check(a.ctx, err)

	w.Strongs = make([]ScriptureWordValue_Strongs, 0, 1)
	var num, grammar string
	for {
		hasRow, err := stmt.Step()
		check(a.ctx, err)
		if !hasRow {
			break
		}

		err = stmt.Scan(&num, &grammar)
		check(a.ctx, err)

		w.Strongs = append(w.Strongs, ScriptureWordValue_Strongs{Num: num, Grammar: grammar})
	}

	stmt.Reset()
	a.db_pool <- db
	wg.Done()
}

func (a *App) GetScriptureWord(ref ScriptureRef, word int) ScriptureWordValue {
	w := ScriptureWordValue{}

	wg := new(sync.WaitGroup)
	wg.Add(3)
	a.getScriptureWordBasicInfo(wg, &w, ref, word)
	a.getScriptureWordDictionaryValues(wg, &w, ref, word)
	a.getScriptureWordStrongsValues(wg, &w, ref, word)
	wg.Wait()
	return w
}
