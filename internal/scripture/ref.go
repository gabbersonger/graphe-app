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

func CreateRef(book int, chapter int, verse int) (ScriptureRef, error) {
	if book < 0 || chapter < 0 || verse < 0 {
		return 0, fmt.Errorf("Cannot pass negative values to CreateRef. args=(%d, %d, %d)", book, chapter, verse)
	}
	if book > 999 || chapter > 999 || verse > 999 {
		return 0, fmt.Errorf("Cannot pass values greater than 999 to CreateRef.  args=(%d, %d, %d)", book, chapter, verse)
	}
	ref_string := fmt.Sprint(book)
	ref_string += padLeftString(fmt.Sprint(chapter), 3)
	ref_string += padLeftString(fmt.Sprint(verse), 3)
	r, err := strconv.Atoi(ref_string)
	if err != nil {
		return 0, fmt.Errorf("Could not create a ref in CreateRef. args=(%d, %d, %d)", book, chapter, verse)
	}
	return ScriptureRef(r), nil
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

func (r ScriptureRef) IsSuperscriptChapter(v ScriptureVersion) bool {
	book_data, err := GetVersionBookData(v, r.GetBook())
	if err != nil {
		return false
	}

	chapter := r.GetChapter()
	if book_data.Superscripts != nil && len(book_data.Superscripts) > 0 {
		for _, c := range book_data.Superscripts {
			if c == chapter {
				return true
			}
		}
	}
	return false
}

func (r ScriptureRef) IsValid(v ScriptureVersion) bool {
	book := r.GetBook()
	chapter := r.GetChapter()
	verse := r.GetVerse()

	// Handle valid book
	book_data, err := GetVersionBookData(v, book)
	if err != nil {
		return false
	}

	// Handle prologue case
	if chapter == 0 && book_data.Prologue > 0 {
		return verse > 0 && verse <= book_data.Prologue
	}

	// Handle chapter invalid
	if chapter <= 0 || chapter > book_data.NumChapters {
		return false
	}

	// Handle verse invalid
	if verse == 0 {
		return r.IsSuperscriptChapter(v)
	}
	if verse < 0 || verse > book_data.NumVerses[chapter-1] {
		return false
	}

	// Handle missing sections
	for i := 0; i < len(book_data.MissingSections); i++ {
		section := book_data.MissingSections[i]
		if section.Start.GetChapter() == chapter &&
			verse >= section.Start.GetVerse() &&
			verse <= section.End.GetVerse() {
			return false
		}
	}

	return true
}

func (r ScriptureRef) IsBookStart(v ScriptureVersion) bool {
	book := r.GetBook()
	chapter := r.GetChapter()
	verse := r.GetVerse()

	book_data, err := GetVersionBookData(v, book)
	if err != nil {
		return false
	}

	if chapter == 0 && book_data.Prologue > 0 {
		return verse == 1
		// This is because no prologue is missing verse 1
	}

	if chapter != 1 {
		return false
	}

	if verse == 0 {
		return len(book_data.Superscripts) > 0 && book_data.Superscripts[0] == 1
	}

	first_verse_in_chapter := 1
	for _, s := range book_data.MissingSections {
		if s.Start.GetChapter() == 1 {
			if s.Start.GetVerse() <= first_verse_in_chapter {
				first_verse_in_chapter = s.End.GetVerse() + 1
				// Don't need to worry about reaching end of chapter
				// as not book is missing all of chapter 1
			}
		} else {
			break
		}
	}

	return verse == first_verse_in_chapter
}
