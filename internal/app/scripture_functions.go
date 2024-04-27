package app

import (
	"fmt"
	"strconv"
	"strings"
)

func (r ScriptureRef) getBook() int {
	return (int(r) - int(r)%1_000_000) / 1_000_000
}

func (r ScriptureRef) getChapter() int {
	return ((int(r) - int(r)%1_000) / 1_000) % 1_000
}

func (r ScriptureRef) getVerse() int {
	return int(r) % 1000
}

func getVersionIndex(version ScriptureVersion) int {
	for i, v := range versionData {
		if v.name == string(version) {
			return i
		}
	}
	return -1
}

func getVersionBookIndex(version ScriptureVersion, book int) int {
	version_index := getVersionIndex(version)
	if version_index == -1 {
		return -1
	}
	for i, b := range versionData[version_index].books {
		if b.book_number == book {
			return i
		}
	}
	return -1
}

func isSuperscriptChapter(version ScriptureVersion, book int, chapter int) bool {
	version_index := getVersionIndex(version)
	book_index := getVersionBookIndex(version, book)
	if version_index == -1 || book_index == -1 {
		return false
	}
	book_data := versionData[version_index].books[book_index]
	if book_data.superscripts != nil && len(book_data.superscripts) > 0 {
		for _, c := range book_data.superscripts {
			if c == chapter {
				return true
			}
		}
	}
	return false
}

func (r ScriptureRef) isValidRef(version ScriptureVersion) bool {
	version_index := getVersionIndex(version)
	if version_index == -1 {
		return false
	}

	book := r.getBook()
	chapter := r.getChapter()
	verse := r.getVerse()

	// Handle book invalid
	book_index := getVersionBookIndex(version, book)
	if book_index == -1 {
		return false
	}
	book_data := versionData[version_index].books[book_index]

	// Handle prologue case
	if chapter == 0 && book_data.prologue > 0 {
		return verse > 0 && verse <= book_data.prologue
	}

	// Handle chapter invalid
	if chapter <= 0 || chapter > book_data.num_chapters {
		return false
	}

	// Handle verse invalid
	if verse == 0 {
		return isSuperscriptChapter(version, book, chapter)
	}
	if verse < 0 || verse > book_data.num_verses[chapter-1] {
		return false
	}

	// Handle missing sections
	for i := 0; i < len(book_data.missing_sections); i++ {
		section := book_data.missing_sections[i]
		if section.start.getChapter() == chapter &&
			verse >= section.start.getVerse() &&
			verse <= section.end.getVerse() {
			return false
		}
	}

	return true
}

func (r ScriptureRange) isValidRange() bool {
	// Handle invalid start or end ref
	if !r.Start.isValidRef(r.Version) || !r.End.isValidRef(r.Version) {
		return false
	}

	// Handle start and end relative positioning in version
	start_book_index := getVersionBookIndex(r.Version, r.Start.getBook())
	end_book_index := getVersionBookIndex(r.Version, r.End.getBook())
	if start_book_index < end_book_index {
		return true
	}
	return start_book_index == end_book_index && r.Start <= r.End
}

func padLeftString(text string, length int) string {
	if len(text) >= length {
		return text
	}
	return strings.Repeat("0", length-len(text)) + text
}

func createRef(book int, chapter int, verse int) (int, error) {
	ref_string := fmt.Sprint(book)
	ref_string += padLeftString(fmt.Sprint(chapter), 3)
	ref_string += padLeftString(fmt.Sprint(verse), 3)
	return strconv.Atoi(ref_string)
}
