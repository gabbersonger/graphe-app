package scripture

import "testing"

func TestGetBook(t *testing.T) {
	var r ScriptureRef
	var b int

	b = r.GetBook()
	if b != 0 {
		t.Errorf("empty ref breaks for ScriptureRef.getBook()")
	}

	r = 0
	b = r.GetBook()
	if b != 0 {
		t.Errorf("ref with value 0 breaks for ScriptureRef.getBook()")
	}

	r = 1_000_000
	b = r.GetBook()
	if b != 1 {
		t.Errorf("basic case for ScriptureRef.getBook() does not work")
	}

	r = 10_000_000
	b = r.GetBook()
	if b != 1 {
		t.Errorf("double digit case for ScriptureRef.getBook() does not work")
	}

	r = 1_001_000
	b = r.GetBook()
	if b != 1 {
		t.Errorf("including chapter value breaks ScriptureRef.getBook()")
	}

	r = 1_000_001
	b = r.GetBook()
	if b != 1 {
		t.Errorf("including verse value breaks ScriptureRef.getBook()")
	}

	r = 1_000_001
	b = r.GetBook()
	if b != 1 {
		t.Errorf("complex case for ScriptureRef.getBook() does not work")
	}
}
