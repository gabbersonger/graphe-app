package scripture_test

import (
	. "graphe/internal"
	. "graphe/internal/scripture"
	"testing"
)

func TestRefCreateBook(t *testing.T) {
	var r ScriptureRef
	var err error

	r, err = CreateRef(1, 1, 1)
	Ok(t, err)
	Equals(t, ScriptureRef(1_001_001), r)

	r, err = CreateRef(-1, 1, 1)
	Assert(t, err != nil, "negative value for book didn't error")
	r, err = CreateRef(1, -1, 1)
	Assert(t, err != nil, "negative value for chapter didn't error")
	r, err = CreateRef(1, 1, -1)
	Assert(t, err != nil, "negative for verse didn't error")

	r, err = CreateRef(1000, 1, 1)
	Assert(t, err != nil, "value over 1000 for book didn't error")
	r, err = CreateRef(1, 1000, 1)
	Assert(t, err != nil, "value over 1000 for chapter didn't error")
	r, err = CreateRef(1, 1, 1000)
	Assert(t, err != nil, "value over 1000 for verse didn't error")

	r, err = CreateRef(100, 101, 102)
	Ok(t, err)
	Equals(t, ScriptureRef(100_101_102), r)
}

func TestRefGetBook(t *testing.T) {
	var r ScriptureRef
	Equals(t, 0, r.GetBook())

	r = 0
	Equals(t, 0, r.GetBook())

	r = 1_000_000
	Equals(t, 1, r.GetBook())

	r = 10_000_000
	Equals(t, 10, r.GetBook())

	r = 1_001_000
	Equals(t, 1, r.GetBook())

	r = 1_000_001
	Equals(t, 1, r.GetBook())

	r = 1_000_001
	Equals(t, 1, r.GetBook())
}

func TestRefGetChapter(t *testing.T) {
	var r ScriptureRef
	Equals(t, 0, r.GetChapter())

	r = 0
	Equals(t, 0, r.GetChapter())

	r = 1_000_000
	Equals(t, 0, r.GetChapter())

	r = 1_000_055
	Equals(t, 0, r.GetChapter())

	r = 1_055_000
	Equals(t, 55, r.GetChapter())

	r = 1_001_000
	Equals(t, 1, r.GetChapter())

	r = 1_055_001
	Equals(t, 55, r.GetChapter())

	r = 1_055_999
	Equals(t, 55, r.GetChapter())

	r = 1_001_001
	Equals(t, 1, r.GetChapter())
}

func TestRefGetVerse(t *testing.T) {
	var r ScriptureRef
	Equals(t, 0, r.GetVerse())

	r = 0
	Equals(t, 0, r.GetVerse())

	r = 1_055_000
	Equals(t, 0, r.GetVerse())

	r = 1_001_001
	Equals(t, 1, r.GetVerse())

	r = 1_001_055
	Equals(t, 55, r.GetVerse())

	r = 1_001_999
	Equals(t, 999, r.GetVerse())

	r = 1_999_001
	Equals(t, 1, r.GetVerse())
}

func TestRefIsSuperscriptChapter(t *testing.T) {
	var r ScriptureRef

	r, _ = CreateRef(19, 3, 1)
	Equals(t, true, r.IsSuperscriptChapter("esv"))
	r, _ = CreateRef(19, 3, 1)
	Equals(t, true, r.IsSuperscriptChapter("lxx"))
	r, _ = CreateRef(19, 25, 0)
	Equals(t, true, r.IsSuperscriptChapter("esv"))
	r, _ = CreateRef(19, 25, 0)
	Equals(t, true, r.IsSuperscriptChapter("lxx"))
	r, _ = CreateRef(19, 50, 4)
	Equals(t, true, r.IsSuperscriptChapter("esv"))
	r, _ = CreateRef(19, 50, 4)
	Equals(t, true, r.IsSuperscriptChapter("lxx"))
	r, _ = CreateRef(19, 145, 999)
	Equals(t, true, r.IsSuperscriptChapter("esv"))
	r, _ = CreateRef(19, 145, 999)
	Equals(t, true, r.IsSuperscriptChapter("lxx"))

	r, _ = CreateRef(1, 1, 1)
	Equals(t, false, r.IsSuperscriptChapter("asd"))
	r, _ = CreateRef(19, 3, 1)
	Equals(t, false, r.IsSuperscriptChapter("asd"))
	r, _ = CreateRef(1, 1, 1)
	Equals(t, false, r.IsSuperscriptChapter("esv"))
	r, _ = CreateRef(19, 3, 0)
	Equals(t, false, r.IsSuperscriptChapter("gnt"))
	r, _ = CreateRef(19, 3, 1)
	Equals(t, false, r.IsSuperscriptChapter("gnt"))
	r, _ = CreateRef(999, 999, 999)
	Equals(t, false, r.IsSuperscriptChapter("esv"))
	r, _ = CreateRef(19, 999, 999)
	Equals(t, false, r.IsSuperscriptChapter("esv"))
	r, _ = CreateRef(-1, -1, -1)
	Equals(t, false, r.IsSuperscriptChapter("esv"))
}

func TestRefIsValid(t *testing.T) {
	var r ScriptureRef

	r, _ = CreateRef(1, 1, 1)
	Equals(t, true, r.IsValid("esv"))
	r, _ = CreateRef(999, 999, 999)
	Equals(t, false, r.IsValid("esv"))
	r, _ = CreateRef(-1, -1, -1)
	Equals(t, false, r.IsValid("esv"))

	// valid extremes
	r, _ = CreateRef(1, 1, 31)
	Equals(t, true, r.IsValid("esv"))
	r, _ = CreateRef(1, 50, 1)
	Equals(t, true, r.IsValid("esv"))
	r, _ = CreateRef(66, 1, 1)
	Equals(t, true, r.IsValid("esv"))
	r, _ = CreateRef(1, 1, 31)
	Equals(t, true, r.IsValid("lxx"))
	r, _ = CreateRef(1, 50, 1)
	Equals(t, true, r.IsValid("lxx"))
	r, _ = CreateRef(81, 1, 1)
	Equals(t, true, r.IsValid("lxx"))

	// just out of bounds
	r, _ = CreateRef(1, 1, 0)
	Equals(t, false, r.IsValid("esv"))
	r, _ = CreateRef(1, 1, 32)
	Equals(t, false, r.IsValid("esv"))
	r, _ = CreateRef(1, 0, 1)
	Equals(t, false, r.IsValid("esv"))
	r, _ = CreateRef(1, 51, 1)
	Equals(t, false, r.IsValid("esv"))
	r, _ = CreateRef(0, 1, 1)
	Equals(t, false, r.IsValid("esv"))
	r, _ = CreateRef(67, 1, 1)
	Equals(t, false, r.IsValid("esv"))

	// superscripts
	r, _ = CreateRef(19, 3, 0)
	Equals(t, true, r.IsValid("esv"))
	r, _ = CreateRef(74, 1, 0)
	Equals(t, true, r.IsValid("lxx"))

	// valid in other versions
	r, _ = CreateRef(1, 1, 1)
	Equals(t, false, r.IsValid("gnt"))
	r, _ = CreateRef(40, 1, 1)
	Equals(t, false, r.IsValid("lxx"))

	// missing sections
	r, _ = CreateRef(40, 12, 46)
	Equals(t, true, r.IsValid("esv"))
	r, _ = CreateRef(40, 12, 47)
	Equals(t, false, r.IsValid("esv"))
	r, _ = CreateRef(40, 12, 48)
	Equals(t, true, r.IsValid("esv"))

	r, _ = CreateRef(2, 35, 16)
	Equals(t, true, r.IsValid("lxx"))
	r, _ = CreateRef(2, 35, 17)
	Equals(t, false, r.IsValid("lxx"))
	r, _ = CreateRef(2, 35, 18)
	Equals(t, false, r.IsValid("lxx"))
	r, _ = CreateRef(2, 35, 19)
	Equals(t, true, r.IsValid("lxx"))

	r, _ = CreateRef(9, 17, 11)
	Equals(t, true, r.IsValid("lxx"))
	r, _ = CreateRef(9, 17, 12)
	Equals(t, false, r.IsValid("lxx"))
	r, _ = CreateRef(9, 17, 31)
	Equals(t, false, r.IsValid("lxx"))
	r, _ = CreateRef(9, 17, 32)
	Equals(t, true, r.IsValid("lxx"))

	// missing in other versions
	r, _ = CreateRef(2, 35, 18)
	Equals(t, true, r.IsValid("esv"))
	r, _ = CreateRef(9, 17, 31)
	Equals(t, true, r.IsValid("esv"))

	// prologues
	r, _ = CreateRef(76, 0, 0)
	Equals(t, false, r.IsValid("lxx"))
	r, _ = CreateRef(76, 0, 1)
	Equals(t, true, r.IsValid("lxx"))
	r, _ = CreateRef(76, 0, 36)
	Equals(t, true, r.IsValid("lxx"))
	r, _ = CreateRef(76, 0, 37)
	Equals(t, false, r.IsValid("lxx"))
	r, _ = CreateRef(76, 0, 1)
	Equals(t, false, r.IsValid("esv"))
}

func TestRefIsBookStart(t *testing.T) {
	var r ScriptureRef

	r, _ = CreateRef(1, 1, 1)
	Equals(t, true, r.IsBookStart("esv"))
	Equals(t, true, r.IsBookStart("lxx"))
	Equals(t, false, r.IsBookStart("gnt"))

	r, _ = CreateRef(40, 1, 1)
	Equals(t, true, r.IsBookStart("esv"))
	Equals(t, false, r.IsBookStart("lxx"))
	Equals(t, true, r.IsBookStart("gnt"))

	r, _ = CreateRef(1, 2, 1)
	Equals(t, false, r.IsBookStart("esv"))
	r, _ = CreateRef(1, 50, 1)
	Equals(t, false, r.IsBookStart("esv"))
	r, _ = CreateRef(999, 999, 999)
	Equals(t, false, r.IsBookStart("esv"))
	r, _ = CreateRef(-1, -1, -1)
	Equals(t, false, r.IsBookStart("esv"))

	// missing start of book
	r, _ = CreateRef(80, 1, 1)
	Equals(t, false, r.IsBookStart("lxx"))
	r, _ = CreateRef(80, 1, 2)
	Equals(t, false, r.IsBookStart("lxx"))
	r, _ = CreateRef(80, 1, 5)
	Equals(t, false, r.IsBookStart("lxx"))
	r, _ = CreateRef(80, 1, 6)
	Equals(t, true, r.IsBookStart("lxx"))

	// superscripts
	r, _ = CreateRef(25, 1, 0)
	Equals(t, true, r.IsBookStart("lxx"))
	r, _ = CreateRef(25, 1, 1)
	Equals(t, false, r.IsBookStart("lxx"))
	r, _ = CreateRef(1, 1, 0)
	Equals(t, false, r.IsBookStart("lxx"))

	// prologues
	r, _ = CreateRef(76, 0, 1)
	Equals(t, true, r.IsBookStart("lxx"))
	r, _ = CreateRef(76, 1, 1)
	Equals(t, false, r.IsBookStart("lxx"))
	r, _ = CreateRef(1, 0, 1)
	Equals(t, false, r.IsBookStart("lxx"))
}

func TestRefToString(t *testing.T) {
	var r ScriptureRef
	var s string
	var err error

	// valid ref
	r, _ = CreateRef(1, 1, 1)

	s, err = r.ToString("esv", 0)
	Ok(t, err)
	Equals(t, "Gen 1:1", s)
	s, err = r.ToString("lxx", 0)
	Ok(t, err)
	Equals(t, "Gen 1:1", s)
	_, err = r.ToString("gnt", 0)
	Assert(t, err != nil, "making short string for gen 1:1 (gnt) didn't error")

	s, err = r.ToString("esv", 1)
	Ok(t, err)
	Equals(t, "Genesis 1:1", s)
	s, err = r.ToString("lxx", 1)
	Ok(t, err)
	Equals(t, "Genesis 1:1", s)
	s, err = r.ToString("gnt", 1)
	Assert(t, err != nil, "making long string for gen 1:1 (gnt) didn't error")

	s, err = r.ToString("esv", 2)
	Ok(t, err)
	Equals(t, "Gen 1", s)
	s, err = r.ToString("lxx", 2)
	Ok(t, err)
	Equals(t, "Gen 1", s)
	s, err = r.ToString("gnt", 2)
	Assert(t, err != nil, "making chapter string for gen 1:1 (gnt) didn't error")

	s, err = r.ToString("esv", 3)
	Ok(t, err)
	Equals(t, "Genesis", s)
	s, err = r.ToString("lxx", 3)
	Ok(t, err)
	Equals(t, "Genesis", s)
	s, err = r.ToString("gnt", 3)
	Assert(t, err != nil, "making book string for gen 1:1 (gnt) didn't error")

	// invalid ref
	r, _ = CreateRef(1, 51, 1)
	s, err = r.ToString("esv", 0)
	Assert(t, err != nil, "invalid ref didn't error for short string")
	s, err = r.ToString("esv", 1)
	Assert(t, err != nil, "invalid ref didn't error for long string")
	s, err = r.ToString("esv", 2)
	Assert(t, err != nil, "invalid ref didn't error for chapter string")
	s, err = r.ToString("esv", 3)
	Assert(t, err != nil, "invalid ref didn't error for book string")

	// single chapter book
	r, _ = CreateRef(64, 1, 1)
	s, err = r.ToString("esv", 0)
	Ok(t, err)
	Equals(t, "3Jo 1", s)
	s, err = r.ToString("esv", 1)
	Ok(t, err)
	Equals(t, "3 John 1", s)
	s, err = r.ToString("esv", 2)
	Ok(t, err)
	Equals(t, "3Jo", s)
	s, err = r.ToString("esv", 3)
	Ok(t, err)
	Equals(t, "3 John", s)

	r, _ = CreateRef(31, 1, 21)
	s, err = r.ToString("esv", 0)
	Ok(t, err)
	Equals(t, "Obad 21", s)
	s, err = r.ToString("esv", 1)
	Ok(t, err)
	Equals(t, "Obadiah 21", s)
	s, err = r.ToString("esv", 2)
	Ok(t, err)
	Equals(t, "Obad", s)
	s, err = r.ToString("esv", 3)
	Ok(t, err)
	Equals(t, "Obadiah", s)
}
