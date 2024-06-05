package database

import (
	"fmt"
	"graphe/internal/scripture"
	"strings"
	"sync"

	"github.com/bvinc/go-sqlite-lite/sqlite3"
)

type ScriptureSection struct {
	Range  scripture.ScriptureRange `json:"range"`
	Blocks []ScriptureBlock         `json:"blocks"`
}

type ScriptureBlock struct {
	Range  scripture.ScriptureRange `json:"range"`
	Verses []ScriptureVerse         `json:"verses"`
}

type ScriptureVerse struct {
	Ref          scripture.ScriptureRef `json:"ref"`
	Words        []ScriptureWord        `json:"words"`
	Details      []ScriptureVerseDetail `json:"details,omitempty"`
	Continuation bool                   `json:"continuation,omitempty"`
}

type ScriptureVerseDetail struct {
	Type ScriptureVerseDetailType `json:"type"`
	Data string                   `json:"data,omitempty"`
}

type ScriptureWord struct {
	WordNumber       int                   `json:"word_num"`
	Text             string                `json:"text"`
	Pre              string                `json:"pre"`
	Post             string                `json:"post"`
	Details          []ScriptureWordDetail `json:"details,omitempty"`
	NoInstantDetails bool                  `json:"no_instant_details,omitempty"`
}

type ScriptureWordDetail struct {
	Type           ScriptureWordDetailType `json:"type"`
	PositionBefore bool                    `json:"position,omitempty"`
	Data           string                  `json:"data,omitempty"`
}

type ScriptureVerseDetailType int
type ScriptureWordDetailType int

const (
	Title   ScriptureVerseDetailType = iota
	Heading ScriptureVerseDetailType = iota
)

const (
	NewLine  ScriptureWordDetailType = iota
	Footnote ScriptureWordDetailType = iota
	Crossref ScriptureWordDetailType = iota
)

func (g *GrapheDB) GetScriptureSection(r scripture.ScriptureRange) []ScriptureSection {
	if !r.IsValid() {
		g.throw(fmt.Sprintf("Invalid range for GetSection. args=({version: %s, start: %d, end:%d})", string(r.Version), r.Start, r.End))
	}

	data := createEmptyBookSections(g, r)
	wg := new(sync.WaitGroup)
	wg.Add(len(data))
	for i := 0; i < len(data); i++ {
		go getScriptureSection(g, wg, &data[i])
	}
	wg.Wait()
	return data
}

func createEmptyBookSections(g *GrapheDB, r scripture.ScriptureRange) []ScriptureSection {
	v_i, err := scripture.GetVersionIndex(r.Version)
	if err != nil {
		g.throw(fmt.Sprintf("Invalid version passed to createDividedRangesSections. args=({version: %s, start: %d, end:%d})", string(r.Version), r.Start, r.End))
	}
	start_vb_i, err := scripture.GetVersionBookIndex(r.Version, r.Start.GetBook())
	if err != nil {
		g.throw(fmt.Sprintf("Invalid starting book (in specific version) passed to createDividedRangesSections. args=({version: %s, start: %d, end:%d})", string(r.Version), r.Start, r.End))
	}
	end_vb_i, err := scripture.GetVersionBookIndex(r.Version, r.End.GetBook())
	if err != nil {
		g.throw(fmt.Sprintf("Invalid ending book (in specific version) passed to createDividedRangesSections. args=({version: %s, start: %d, end:%d})", string(r.Version), r.Start, r.End))
	}

	data := make([]ScriptureSection, 0, end_vb_i-start_vb_i+1)
	for i := start_vb_i; i <= end_vb_i; i++ {
		book_data := scripture.VersionsData[v_i].Books[i]
		var book_start scripture.ScriptureRef
		var err error
		if book_data.Prologue > 0 {
			book_start, err = scripture.CreateRef(book_data.BookNumber, 0, 1)
		} else {
			book_start, err = scripture.CreateRef(book_data.BookNumber, 1, 1)
		}
		g.check(err)
		book_end, err := scripture.CreateRef(book_data.BookNumber, book_data.NumChapters, book_data.NumVerses[len(book_data.NumVerses)-1])
		g.check(err)
		data = append(data, ScriptureSection{Range: scripture.ScriptureRange{
			Version: r.Version,
			Start:   (book_start),
			End:     (book_end),
		}})
	}
	return data
}

const DEFAULT_VERSE_COUNT = 35
const DEFAULT_WORD_COUNT = 30

func getScriptureSection(g *GrapheDB, wg *sync.WaitGroup, s *ScriptureSection) {
	s.Blocks = make([]ScriptureBlock, 0, 1)

	db := <-g.Pool
	var stmt *sqlite3.Stmt
	switch s.Range.Version {
	case "gnt":
		stmt = db.queries.gntSection
	case "lxx":
		stmt = db.queries.lxxSection
	case "esv":
		stmt = db.queries.esvSection
	default:
		g.throw(fmt.Sprintf("Unknown version passed to getScriptureSection. args=(%s)", string(s.Range.Version)))
	}

	err := stmt.Bind(int(s.Range.Start), int(s.Range.End))
	g.check(err)

	var ref, word_num int
	var text, pre, post string
	has_instant_details := 1
	createNextBlock := true
	lastRef := 0

	for {
		hasRow, err := stmt.Step()
		g.check(err)
		if !hasRow {
			break
		}
		if s.Range.Version == "esv" {
			err = stmt.Scan(&ref, &word_num, &text, &pre, &post, &has_instant_details)
		} else {
			err = stmt.Scan(&ref, &word_num, &text, &pre, &post)
		}
		g.check(err)

		// Add block if needed
		if createNextBlock {
			createNextBlock = false
			if len(s.Blocks) > 0 {
				s.Blocks[len(s.Blocks)-1].Range.End = scripture.ScriptureRef(lastRef)
			}

			newBlock := ScriptureBlock{}
			newBlock.Range.Version = s.Range.Version
			newBlock.Range.Start = scripture.ScriptureRef(ref)
			newBlock.Verses = make([]ScriptureVerse, 0, DEFAULT_VERSE_COUNT)
			s.Blocks = append(s.Blocks, newBlock)
		}
		lastBlock := len(s.Blocks) - 1

		// Add verse if needed
		if lastRef != ref || len(s.Blocks[lastBlock].Verses) == 0 {
			newVerse := ScriptureVerse{}
			newVerse.Ref = scripture.ScriptureRef(ref)
			newVerse.Words = make([]ScriptureWord, 0, DEFAULT_WORD_COUNT)
			newVerse.Continuation = lastRef == ref
			newVerse.Details = make([]ScriptureVerseDetail, 0)
			if len(s.Blocks[lastBlock].Verses) == 0 {
				if lastRef != ref && newVerse.Ref.IsBookStart(s.Range.Version) {
					book_data, err := scripture.GetVersionBookData(s.Range.Version, newVerse.Ref.GetBook())
					if err != nil {
						g.throw(fmt.Sprintf("Invalid ref pulled from database for version in getScriptureSection. details=(version: %s, ref: %d)", s.Range.Version, newVerse.Ref))
					}
					newVerse.Details = append(newVerse.Details, ScriptureVerseDetail{
						Type: Title,
						Data: book_data.DisplayName,
					})
				}
				// TODO: headings
			}
			s.Blocks[lastBlock].Verses = append(s.Blocks[lastBlock].Verses, newVerse)

			lastRef = ref
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
			} else if rune == '@' {

			}
		}
		if n >= 0 {
			runes = append(runes[:n], runes[n+1:]...)
			post = string(runes)
		}

		// Replace _ with —
		if len(post) > 0 && post[0] == '_' {
			post = " —" + post[1:]
		} else if strings.ContainsAny(post, "_") {
			post = strings.ReplaceAll(post, "_", "—")
		}
		if len(pre) > 0 && pre[len(pre)-1] == '_' {
			pre = pre[:len(pre)-1] + "— "
		} else if strings.ContainsAny(pre, "_") {
			pre = strings.ReplaceAll(pre, "_", "—")
		}

		var details []ScriptureWordDetail = make([]ScriptureWordDetail, 0)

		// TODO: footnotes and crossrefs

		// Check for poetry line-breaks
		if len(post) > 0 && post[len(post)-1] == '@' {
			post = post[:len(post)-1]
			details = append(details, ScriptureWordDetail{
				Type: NewLine,
			})
		}

		// Add word
		newWord := ScriptureWord{word_num, text, pre, post, details, has_instant_details == 0}
		s.Blocks[lastBlock].Verses[lastVerse].Words = append(s.Blocks[lastBlock].Verses[lastVerse].Words, newWord)
	}
	s.Range.End = scripture.ScriptureRef(ref)
	s.Blocks[len(s.Blocks)-1].Range.End = scripture.ScriptureRef(ref)

	stmt.Reset()
	g.Pool <- db
	wg.Done()
}
