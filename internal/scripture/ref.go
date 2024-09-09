package scripture

import (
	"fmt"
	"strconv"
	"strings"
)

func padLeftString(text string, length int) string {
	if len(text) >= length {
		return text
	}
	return strings.Repeat("0", length-len(text)) + text
}

func (s *ScriptureService) CreateRef(book, chapter, verse int) ScriptureRef {
	s.assert(book >= 0 && book <= 999, fmt.Sprintf("Invalid book (book: %d)", book))
	s.assert(chapter >= 0 && chapter <= 999, fmt.Sprintf("Invalid chapter (chapter: %d)", chapter))
	s.assert(verse >= 0 && verse <= 999, fmt.Sprintf("Invalid verse (verse: %d)", verse))

	ref_string := fmt.Sprint(book)
	ref_string += padLeftString(fmt.Sprint(chapter), 3)
	ref_string += padLeftString(fmt.Sprint(verse), 3)
	r, err := strconv.Atoi(ref_string)
	s.assert(err == nil, fmt.Sprintf("Error creating ref (book: %d, chapter: %d, verse: %d)", book, chapter, verse))
	return ScriptureRef(r)
}

func (s *ScriptureService) CreateFirstValidRef(version ScriptureVersion, book int) ScriptureRef {
	version_index := s.getVersionIndex(version)
	s.assert(version_index >= 0 && version_index < len(VersionsData), fmt.Sprintf("Error getting version index (version: `%s`)", version))
	vb_index := s.getVersionBookIndex(version, book)
	s.assert(vb_index >= 0 && vb_index < len(VersionsData[version_index].Books), fmt.Sprintf("Error getting version book index (version: `%s`, book: %d, index: %d)", version, book, vb_index))
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
		if s.GetRefChapter(section.Start) == chapter &&
			verse >= s.GetRefVerse(section.Start) &&
			verse <= s.GetRefVerse(section.End) {
			verse = s.GetRefVerse(section.End) + 1
			// Assumed: no book missing all of chapter 0 and 1
		}
	}

	return s.CreateRef(book, chapter, verse)
}

func (s *ScriptureService) GetRefBook(r ScriptureRef) int {
	return (int(r) - int(r)%1_000_000) / 1_000_000
}

func (s *ScriptureService) GetRefChapter(r ScriptureRef) int {
	return ((int(r) - int(r)%1_000) / 1_000) % 1_000
}

func (s *ScriptureService) GetRefVerse(r ScriptureRef) int {
	return int(r) % 1000
}

func (s *ScriptureService) BCV(r ScriptureRef) (int, int, int) {
	return s.GetRefBook(r), s.GetRefChapter(r), s.GetRefVerse(r)
}

func (s *ScriptureService) IsRefBookStart(ref ScriptureRef, version ScriptureVersion) bool {
	s.assert(s.IsVersionValid(version), fmt.Sprintf("Invalid version (version: `%s`)", version))
	if !s.IsRefValid(ref, version) {
		return false
	}

	book, chapter, verse := s.BCV(ref)
	vb_data := s.getVersionBookData(version, book)
	s.assert(vb_data != nil, fmt.Sprintf("Error getting version book data (version: `%s`, ref: %d)", version, int(ref)))

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
		for _, ms := range vb_data.MissingSections {
			if s.GetRefChapter(ms.Start) == 1 {
				if s.GetRefVerse(ms.Start) <= first_verse_in_chapter {
					first_verse_in_chapter = s.GetRefVerse(ms.End) + 1
					// Assumed: no book missing all of chapter 1
				}
			} else {
				break
			}
		}
	}
	return verse == first_verse_in_chapter
}

func (s *ScriptureService) isRefSuperscriptChapter(ref ScriptureRef, version ScriptureVersion) bool {
	s.assert(s.IsVersionValid(version), fmt.Sprintf("Invalid version (version: `%s`)", version))
	vb_data := s.getVersionBookData(version, s.GetRefBook(ref))
	s.assert(vb_data != nil, fmt.Sprintf("Error getting version book data (version: `%s`, ref: %d)", version, int(ref)))

	chapter := s.GetRefChapter(ref)
	if vb_data.Superscripts != nil && len(vb_data.Superscripts) > 0 {
		for _, c := range vb_data.Superscripts {
			if c == chapter {
				return true
			}
		}
	}
	return false
}

func (s *ScriptureService) IsRefValid(ref ScriptureRef, version ScriptureVersion) bool {
	s.assert(s.IsVersionValid(version), fmt.Sprintf("Invalid version (version: `%s`)", version))

	book, chapter, verse := s.BCV(ref)
	vb_data := s.getVersionBookData(version, book)
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
		return s.isRefSuperscriptChapter(ref, version)
	}
	if verse < 0 || verse > vb_data.NumVerses[chapter-1] {
		return false
	}

	// Handle missing sections
	for i := 0; i < len(vb_data.MissingSections); i++ {
		section := vb_data.MissingSections[i]
		if s.GetRefChapter(section.Start) == chapter &&
			verse >= s.GetRefVerse(section.Start) &&
			verse <= s.GetRefVerse(section.End) {
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

func (s *ScriptureService) RefToString(ref ScriptureRef, version ScriptureVersion, format ScriptureRefStringType) string {
	s.assert(s.IsRefValid(ref, version), fmt.Sprintf("Invalid ref and version pair (ref: %d, version: `%s`)", ref, version))

	book, chapter, verse := s.BCV(ref)
	vb_data := s.getVersionBookData(version, book)
	s.assert(vb_data != nil, fmt.Sprintf("Error getting version book data (version: `%s`, book: %d)", version, book))

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
		s.assert(false, fmt.Sprintf("Invalid format (format: %d)", format))
	}
	return "" // Unreachable
}
