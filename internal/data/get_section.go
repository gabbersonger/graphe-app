package data

import (
	"fmt"
	"graphe/internal/scripture"
	"strings"
	"sync"

	"github.com/bvinc/go-sqlite-lite/sqlite3"
)

// TYPES
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
	RefString    string                 `json:"ref_string"`
	Words        []ScriptureWord        `json:"words"`
	Details      []ScriptureVerseDetail `json:"details,omitempty"`
	Continuation bool                   `json:"continuation,omitempty"`
}

type ScriptureVerseDetail struct {
	Type ScriptureVerseDetailType `json:"type"`
	Data string                   `json:"data,omitempty"`
}

type ScriptureWord struct {
	WordNumber        int                   `json:"word_num"`
	Text              string                `json:"text"`
	Pre               string                `json:"pre"`
	Post              string                `json:"post"`
	Details           []ScriptureWordDetail `json:"details,omitempty"`
	HasInstantDetails bool                  `json:"has_instant_details"`
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
	Indent   ScriptureWordDetailType = iota
	Footnote ScriptureWordDetailType = iota
	Crossref ScriptureWordDetailType = iota
)

// FUNCTIONS

func (d *DataDB) GetScriptureSection(r scripture.ScriptureRange) []ScriptureSection {
	d.assert(d.scripture_service.IsRangeValid(r), fmt.Sprintf("Invalid range (version: `%s`, start: %d, end: %d)", r.Version, int(r.Start), int(r.End)))

	data := createEmptyBookSections(d, r)
	d.assert(len(data) > 0, "Error creating empty sections for range")

	wg := new(sync.WaitGroup)
	wg.Add(len(data))
	for i := 0; i < len(data); i++ {
		go getScriptureSection(d, wg, &data[i])
	}
	wg.Wait()

	return data
}

func createEmptyBookSections(d *DataDB, r scripture.ScriptureRange) []ScriptureSection {
	ranges := d.scripture_service.DivideIntoBookRanges(r)
	d.assert(len(ranges) > 0, fmt.Sprintf("Error dividing range into books (version: `%s`, start: %d, end: %d)", r.Version, int(r.Start), int(r.End)))
	d.assert(ranges[0].Start == r.Start, fmt.Sprintf("Error in created range: start not included (version: `%s`, start: %d, end: %d)", r.Version, int(r.Start), int(r.End)))
	d.assert(ranges[len(ranges)-1].End == r.End, fmt.Sprintf("Error in created range: end not included (version: `%s`, start: %d, end: %d)", r.Version, int(r.Start), int(r.End)))

	sections := make([]ScriptureSection, 0, len(ranges))
	for _, r := range ranges {
		sections = append(sections, ScriptureSection{
			Range: r,
		})
	}
	return sections
}

const ESTIMATED_BLOCK_COUNT = 1
const ESTIMATED_VERSE_COUNT = 35
const ESTIMATED_WORD_COUNT = 30

func getScriptureSection(d *DataDB, wg *sync.WaitGroup, section *ScriptureSection) {
	section.Blocks = make([]ScriptureBlock, 0, ESTIMATED_BLOCK_COUNT)

	db := <-d.pool
	var stmt *sqlite3.Stmt
	switch section.Range.Version {
	case "gnt":
		stmt = db.queries.gntSection
	case "lxx":
		stmt = db.queries.lxxSection
	case "esv":
		stmt = db.queries.esvSection
	default:
		d.assert(false, fmt.Sprintf("Invalid version (version:`%s`)", section.Range.Version))
	}

	err := stmt.Bind(int(section.Range.Start), int(section.Range.End))
	d.assert(err == nil, fmt.Sprintf("Error binding range to sql statement (version: `%s`, start: %d, end: %d) (err: %v)", section.Range.Version, int(section.Range.Start), int(section.Range.End), err))

	var ref, word_num int
	var text, pre, post string
	has_instant_details := 1

	create_next_block := true
	last_ref := 0
	for {
		has_row, err := stmt.Step()
		d.assert(err == nil, fmt.Sprintf("Error stepping through sql statement (version: `%s`, start: %d, end: %d) (err: %v)", section.Range.Version, int(section.Range.Start), int(section.Range.End), err))
		if !has_row {
			break
		}

		err = stmt.Scan(&ref, &word_num, &text, &pre, &post, &has_instant_details)
		d.assert(err == nil, fmt.Sprintf("Error scanning row (version: `%s`, start: %d, end: %d) (err: %v)", section.Range.Version, int(section.Range.Start), int(section.Range.End), err))

		// Add block if necessary
		if create_next_block {
			create_next_block = false
			if len(section.Blocks) > 0 {
				section.Blocks[len(section.Blocks)-1].Range.End = scripture.ScriptureRef(last_ref)
			}
			addScriptureBlock(section, scripture.ScriptureRef(ref))
		}
		last_block := len(section.Blocks) - 1

		// Add verse if necessary
		if ref != last_ref || len(section.Blocks[last_block].Verses) == 0 {
			addScriptureVerse(d, section, scripture.ScriptureRef(ref), last_ref == ref)
			last_ref = ref
		}
		last_verse := len(section.Blocks[last_block].Verses) - 1

		// Check for paragraph breaks
		runes := []rune(post)
		for i, rune := range runes {
			if rune == '¶' {
				create_next_block = true
				post = string(runes[:i]) + string(runes[i+1:])
				break
			}
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

		// TODO: footnotes and cross-references

		// Check for poetry line-breaks
		if len(post) > 0 && post[len(post)-1] == '@' {
			post = post[:len(post)-1]
			details = append(details, ScriptureWordDetail{Type: NewLine})
		}

		// Check for poetry line-indents
		if len(pre) > 0 && pre[0] == '~' {
			if len(pre) > 1 {
				pre = pre[1:]
			} else {
				pre = ""
			}
			details = append(details, ScriptureWordDetail{Type: Indent})
		}

		// Add word
		new_word := ScriptureWord{word_num, text, pre, post, details, has_instant_details == 1}
		section.Blocks[last_block].Verses[last_verse].Words = append(section.Blocks[last_block].Verses[last_verse].Words, new_word)
	}
	section.Range.End = scripture.ScriptureRef(ref)
	section.Blocks[len(section.Blocks)-1].Range.End = scripture.ScriptureRef(ref)

	stmt.Reset()
	d.pool <- db
	wg.Done()
}

func addScriptureBlock(section *ScriptureSection, ref scripture.ScriptureRef) {
	new_block := ScriptureBlock{}
	new_block.Range.Version = section.Range.Version
	new_block.Range.Start = ref
	new_block.Verses = make([]ScriptureVerse, 0, ESTIMATED_VERSE_COUNT)
	section.Blocks = append(section.Blocks, new_block)
}

func addScriptureVerse(d *DataDB, section *ScriptureSection, ref scripture.ScriptureRef, continuation bool) {
	last_block := len(section.Blocks) - 1

	new_verse := ScriptureVerse{}
	new_verse.Ref = ref
	if len(section.Blocks[last_block].Verses) == 0 {
		new_verse.RefString = d.scripture_service.RefToString(ref, section.Range.Version, scripture.StringShort)
	} else {
		new_verse.RefString = fmt.Sprintf("%d", d.scripture_service.GetRefVerse(ref))
	}
	new_verse.Words = make([]ScriptureWord, 0, ESTIMATED_WORD_COUNT)
	new_verse.Continuation = continuation
	new_verse.Details = make([]ScriptureVerseDetail, 0)

	// Add book titles
	if len(section.Blocks[last_block].Verses) == 0 &&
		!continuation &&
		d.scripture_service.IsRefBookStart(ref, section.Range.Version) {
		new_verse.Details = append(new_verse.Details, ScriptureVerseDetail{
			Type: Title,
			Data: d.scripture_service.RefToString(ref, section.Range.Version, scripture.StringVersionBook),
		})
	}

	// TODO: Add section headings

	section.Blocks[last_block].Verses = append(section.Blocks[last_block].Verses, new_verse)
}
