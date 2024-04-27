package app

import (
	"sync"

	"github.com/bvinc/go-sqlite-lite/sqlite3"
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
	// start := time.Now()

	data := make([]ScriptureSection, 0, len(ran))
	for _, r := range ran {
		data = append(data, ScriptureSection{Range: r})
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(ran))
	for i := 0; i < len(ran); i++ {
		go getScriptureSection(a, wg, &data[i])
	}
	wg.Wait()

	// duration := time.Since(start)
	// fmt.Println(duration)
	return data
}

const DEFAULT_VERSE_COUNT = 35
const DEFAULT_WORD_COUNT = 30

func getScriptureSection(a *App, wg *sync.WaitGroup, s *ScriptureSection) {
	s.Blocks = make([]ScriptureBlock, 0, 1)

	db := <-a.db.pool
	var stmt *sqlite3.Stmt
	switch s.Range.Version {
	case "gnt":
		stmt = db.queries.GntSection
	case "lxx":
		stmt = db.queries.LxxSection
	default:
		a.Throw("Unknown version (" + string(s.Range.Version) + ") passed to getScriptureSection")
	}

	err := stmt.Bind(int(s.Range.Start), int(s.Range.End))
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
			newBlock.Verses = make([]ScriptureVerse, 0, DEFAULT_VERSE_COUNT)
			s.Blocks = append(s.Blocks, newBlock)
		}
		lastBlock := len(s.Blocks) - 1

		// Add verse if needed
		if lastRef != ref || len(s.Blocks[lastBlock].Verses) == 0 {
			lastRef = ref

			newVerse := ScriptureVerse{}
			newVerse.Ref = ScriptureRef(ref)
			newVerse.Words = make([]ScriptureWord, 0, DEFAULT_WORD_COUNT)
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
