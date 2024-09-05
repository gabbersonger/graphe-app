package scripture

import (
	"fmt"
	"strconv"
	"strings"
)

func padLeftString(text string, length int) string {
	assert(length >= 0, "Length was less than 0")
	if len(text) >= length {
		return text
	}
	return strings.Repeat("0", length-len(text)) + text
}

func CreateRef(book, chapter, verse int) ScriptureRef {
	assert(book >= 0 && book <= 999, fmt.Sprintf("Invalid book (book: %d)", book))
	assert(chapter >= 0 && chapter <= 999, fmt.Sprintf("Invalid chapter (chapter: %d)", chapter))
	assert(verse >= 0 && verse <= 999, fmt.Sprintf("Invalid verse (verse: %d)", verse))

	ref_string := fmt.Sprint(book)
	ref_string += padLeftString(fmt.Sprint(chapter), 3)
	ref_string += padLeftString(fmt.Sprint(verse), 3)
	r, err := strconv.Atoi(ref_string)
	assert(err == nil, fmt.Sprintf("Error creating ref (book: %d, chapter: %d, verse: %d)", book, chapter, verse))
	return ScriptureRef(r)
}

func CreateFirstValidRef(version ScriptureVersion, book int) ScriptureRef {
	version_index := getVersionIndex(version)
	assert(version_index >= 0 && version_index < len(VersionsData), fmt.Sprintf("Error getting version index (version: `%s`)", version))
	vb_index := getVersionBookIndex(version, book)
	assert(vb_index >= 0 && vb_index < len(VersionsData[version_index].Books), fmt.Sprintf("Error getting version book index (version: `%s`, book: %d, index: %d)", version, book, vb_index))
	vb_data := VersionsData[version_index].Books[vb_index]

	chapter := 1
	verse := 1

	// Handle prologues and superscripts
	if vb_data.Prologue > 0 {
		chapter = 0
	}

	// Handle superscripts
	if len(vb_data.Superscripts) > 0 && vb_data.Superscripts[0] == chapter {
		verse = 0
	}

	// Handle missing sections
	for i := 0; i < len(vb_data.MissingSections); i++ {
		section := vb_data.MissingSections[i]
		if section.Start.GetChapter() == chapter &&
			verse >= section.Start.GetVerse() &&
			verse <= section.End.GetVerse() {
			verse = section.End.GetVerse() + 1
			// Assumed: no book missing all of chapter 0 and 1
		}
	}

	return CreateRef(book, chapter, verse)
}

func (r ScriptureRef) GetBook() int {
	return (int(r) - int(r)%1_000_000) / 1_000_000
}

func (r ScriptureRef) GetChapter() int {
	return ((int(r) - int(r)%1_000) / 1_000) % 1_000
}

func (r ScriptureRef) GetVerse() int {
	return int(r) % 1000
}

func (r ScriptureRef) BCV() (int, int, int) {
	return r.GetBook(), r.GetChapter(), r.GetVerse()
}

func (r ScriptureRef) IsBookStart(version ScriptureVersion) bool {
	assert(version.IsValid(), fmt.Sprintf("Invalid version (version: `%s`)", version))
	if !r.IsValid(version) {
		return false
	}

	book, chapter, verse := r.BCV()
	vb_data := getVersionBookData(version, book)
	assert(vb_data != nil, fmt.Sprintf("Error getting version book data (version: `%s`, ref: %d)", version, int(r)))

	if vb_data.Prologue > 0 {
		return chapter == 0 && verse == 1
		// Assumed: All prologues start at verse 1
	}

	if chapter != 1 {
		return false
	}

	if verse == 0 {
		return len(vb_data.Superscripts) > 0 && vb_data.Superscripts[0] == 1
	}

	first_verse_in_chapter := 1
	if len(vb_data.Superscripts) > 0 && vb_data.Superscripts[0] == 1 {
		first_verse_in_chapter = 0
	} else {
		for _, s := range vb_data.MissingSections {
			if s.Start.GetChapter() == 1 {
				if s.Start.GetVerse() <= first_verse_in_chapter {
					first_verse_in_chapter = s.End.GetVerse() + 1
					// Assumed: no book missing all of chapter 1
				}
			} else {
				break
			}
		}
	}
	return verse == first_verse_in_chapter
}

func (r ScriptureRef) isSuperscriptChapter(version ScriptureVersion) bool {
	assert(version.IsValid(), fmt.Sprintf("Invalid version (version: `%s`)", version))
	vb_data := getVersionBookData(version, r.GetBook())
	assert(vb_data != nil, fmt.Sprintf("Error getting version book data (version: `%s`, ref: %d)", version, int(r)))

	chapter := r.GetChapter()
	if vb_data.Superscripts != nil && len(vb_data.Superscripts) > 0 {
		for _, c := range vb_data.Superscripts {
			if c == chapter {
				return true
			}
		}
	}
	return false
}

func (r ScriptureRef) IsValid(version ScriptureVersion) bool {
	assert(version.IsValid(), fmt.Sprintf("Invalid version (version: `%s`)", version))

	book, chapter, verse := r.BCV()
	vb_data := getVersionBookData(version, book)
	if vb_data == nil {
		return false
	}

	// Handle prologue case
	if chapter == 0 && vb_data.Prologue > 0 {
		return verse > 0 && verse <= vb_data.Prologue
	}

	// Handle chapter invalid
	if chapter <= 0 || chapter > vb_data.NumChapters {
		return false
	}

	// Handle verse invalid
	if verse == 0 {
		return r.isSuperscriptChapter(version)
	}
	if verse < 0 || verse > vb_data.NumVerses[chapter-1] {
		return false
	}

	// Handle missing sections
	for i := 0; i < len(vb_data.MissingSections); i++ {
		section := vb_data.MissingSections[i]
		if section.Start.GetChapter() == chapter &&
			verse >= section.Start.GetVerse() &&
			verse <= section.End.GetVerse() {
			return false
		}
	}

	return true
}

type ScriptureRefStringType int

const (
	StringShort ScriptureRefStringType = iota
	StringLong
	StringChapter
	StringBook
	StringVersionBook
)

func (r ScriptureRef) ToString(version ScriptureVersion, format ScriptureRefStringType) string {
	assert(r.IsValid(version), fmt.Sprintf("Invalid ref and version pair (ref: %d, version: `%s`)", r, version))

	book, chapter, verse := r.BCV()
	vb_data := getVersionBookData(version, book)
	assert(vb_data != nil, fmt.Sprintf("Error getting version book data (version: `%s`, book: %d)", version, book))

	is_single_chapter_book := vb_data.NumChapters == 1

	switch format {
	case StringShort:
		if is_single_chapter_book {
			return fmt.Sprintf("%s %d", BibleData[book-1].Short, verse)
		} else {
			return fmt.Sprintf("%s %d:%d", BibleData[book-1].Short, chapter, verse)
		}
	case StringLong:
		if is_single_chapter_book {
			return fmt.Sprintf("%s %d", BibleData[book-1].Name, verse)
		} else {
			return fmt.Sprintf("%s %d:%d", BibleData[book-1].Name, chapter, verse)
		}
	case StringChapter:
		if is_single_chapter_book {
			return BibleData[book-1].Short
		} else {
			return fmt.Sprintf("%s %d", BibleData[book-1].Short, chapter)
		}
	case StringBook:
		return BibleData[book-1].Name
	case StringVersionBook:
		return vb_data.DisplayName
	default:
		assert(false, fmt.Sprintf("Invalid format (format: %d)", format))
	}
	return "" // Unreachable
}
