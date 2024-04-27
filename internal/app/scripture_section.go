package app

import (
	"fmt"
	"sync"

	"github.com/bvinc/go-sqlite-lite/sqlite3"
)

func (a *App) GetScriptureSection(r ScriptureRange) []ScriptureSection {
	if !r.isValidRange() {
		a.Throw("Invalid range (version=" + string(r.Version) + ", start=" + fmt.Sprint(r.Start) + ", end:" + fmt.Sprint(r.End) + ") passed to `GetScriptureSection`")
	}

	data := makeScriptureSections(a, r)
	wg := new(sync.WaitGroup)
	wg.Add(len(data))
	for i := 0; i < len(data); i++ {
		go getScriptureSection(a, wg, &data[i])
	}
	wg.Wait()
	return data
}

func makeScriptureSections(a *App, r ScriptureRange) []ScriptureSection {
	version_index := getVersionIndex(r.Version)
	start_book_index := getVersionBookIndex(r.Version, r.Start.getBook())
	end_book_index := getVersionBookIndex(r.Version, r.End.getBook())

	data := make([]ScriptureSection, 0, end_book_index-start_book_index+1)
	for i := start_book_index; i <= end_book_index; i++ {
		book_data := versionData[version_index].books[i]
		var book_start int
		var err error
		if book_data.prologue > 0 {
			book_start, err = createRef(book_data.book_number, 0, 1)
		} else {
			book_start, err = createRef(book_data.book_number, 1, 1)
		}
		a.check(err)
		book_end, err := createRef(book_data.book_number, book_data.num_chapters, book_data.num_verses[len(book_data.num_verses)-1])
		a.check(err)
		data = append(data, ScriptureSection{Range: ScriptureRange{
			Version: r.Version,
			Start:   ScriptureRef(book_start),
			End:     ScriptureRef(book_end),
		}})
	}
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
