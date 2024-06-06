package scripture

type ScriptureRange struct {
	Version ScriptureVersion `json:"version"`
	Start   ScriptureRef     `json:"start"`
	End     ScriptureRef     `json:"end"`
}

func (r ScriptureRange) IsValid() bool {
	// Handle invalid start or end ref
	if !r.Start.IsValid(r.Version) || !r.End.IsValid(r.Version) {
		return false
	}

	start_index, err_1 := GetVersionBookIndex(r.Version, r.Start.GetBook())
	end_index, err_2 := GetVersionBookIndex(r.Version, r.End.GetBook())
	if err_1 != nil || err_2 != nil {
		return false
	}

	// Handle start and end relative positioning in version
	if start_index < end_index {
		return true
	}
	return start_index == end_index && r.Start <= r.End
}

func (r ScriptureRange) Contains(ref ScriptureRef) bool {
	if !ref.IsValid(r.Version) || !r.IsValid() {
		return false
	}

	start_book := r.Start.GetBook()
	end_book := r.End.GetBook()
	ref_book := ref.GetBook()

	if start_book == end_book {
		return ref >= r.Start && ref <= r.End
	} else if ref_book == start_book {
		return ref >= r.Start
	} else if ref_book == end_book {
		return ref <= r.End
	}

	start_book_index, err_1 := GetVersionBookIndex(r.Version, start_book)
	end_book_index, err_2 := GetVersionBookIndex(r.Version, end_book)
	ref_book_index, err_3 := GetVersionBookIndex(r.Version, ref_book)
	if err_1 != nil || err_2 != nil || err_3 != nil {
		return false
	}

	return ref_book_index >= start_book_index && ref_book_index <= end_book_index
}
