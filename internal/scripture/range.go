package scripture

import "fmt"

type ScriptureRange struct {
	Version ScriptureVersion `json:"version"`
	Start   ScriptureRef     `json:"start"`
	End     ScriptureRef     `json:"end"`
}

func (r ScriptureRange) IsValid() bool {
	assert(r.Start.IsValid(r.Version), fmt.Sprintf("Invalid start to range (version: `%s`, start: %d, end: %d)", r.Version, int(r.Start), int(r.End)))
	assert(r.End.IsValid(r.Version), fmt.Sprintf("Invalid end to range (version: `%s`, start: %d, end: %d)", r.Version, int(r.Start), int(r.End)))

	start_vb_index := getVersionBookIndex(r.Version, r.Start.GetBook())
	end_vb_index := getVersionBookIndex(r.Version, r.End.GetBook())
	assert(start_vb_index >= 0, fmt.Sprintf("Error getting start's version book index (ref: %d, val: %d)", int(r.Start), start_vb_index))
	assert(end_vb_index >= 0, fmt.Sprintf("Error getting start's version book index (ref: %d, val: %d)", int(r.End), end_vb_index))

	if start_vb_index < end_vb_index {
		return true
	}
	return start_vb_index == end_vb_index && r.Start <= r.End
}

func (r ScriptureRange) Contains(ref ScriptureRef) bool {
	assert(r.IsValid(), fmt.Sprintf("Invalid range (version: `%s`, start: %d, end: %d)", r.Version, int(r.Start), int(r.End)))
	assert(ref.IsValid(r.Version), fmt.Sprintf("Invalid ref for version (version: `%s`, ref: %d)", r.Version, int(ref)))

	start_vb_index := getVersionBookIndex(r.Version, r.Start.GetBook())
	end_vb_index := getVersionBookIndex(r.Version, r.End.GetBook())
	ref_vb_index := getVersionBookIndex(r.Version, ref.GetBook())
	assert(start_vb_index >= 0, fmt.Sprintf("Error getting start's version book index (version: `%s`, ref: %d, val: %d)", r.Version, int(r.Start), start_vb_index))
	assert(end_vb_index >= 0, fmt.Sprintf("Error getting end's version book index (version: `%s`, ref: %d, val: %d)", r.Version, int(r.End), end_vb_index))
	assert(ref_vb_index >= 0, fmt.Sprintf("Error getting ref's version book index (version: `%s`, ref: %d, val: %d)", r.Version, int(ref), ref_vb_index))

	if start_vb_index == end_vb_index { // start = end (must be between)
		return ref >= r.Start && ref <= r.End
	} else if ref_vb_index == start_vb_index { // start = ref (must be after)
		return ref >= r.Start
	} else if ref_vb_index == end_vb_index { // ref = end (must be before)
		return ref <= r.End
	}
	return ref_vb_index > start_vb_index && ref_vb_index < end_vb_index // start =/= end (must be between)
}

func (r ScriptureRange) DivideIntoBookRanges() []ScriptureRange {
	assert(r.IsValid(), fmt.Sprintf("Invalid range (version: `%s`, start: %d, end: %d)", r.Version, int(r.Start), int(r.End)))

	version_index := getVersionIndex(r.Version)
	assert(version_index >= 0 && version_index < len(VersionsData), fmt.Sprintf("Error getting version index (version: `%s`, val: %d)", r.Version, version_index))

	start_vb_index := getVersionBookIndex(r.Version, r.Start.GetBook())
	end_vb_index := getVersionBookIndex(r.Version, r.End.GetBook())
	assert(start_vb_index >= 0, fmt.Sprintf("Error getting start's version book index (version: `%s`, ref: %d, val: %d)", r.Version, int(r.Start), start_vb_index))
	assert(end_vb_index >= 0, fmt.Sprintf("Error getting end's version book index (version: `%s`, ref: %d, val: %d)", r.Version, int(r.End), end_vb_index))
	assert(start_vb_index <= end_vb_index, fmt.Sprintf("Version book index misordering: start before end (start: %d, end %d)", start_vb_index, end_vb_index))

	ranges := make([]ScriptureRange, 0, end_vb_index-start_vb_index+1)
	for i := start_vb_index; i <= end_vb_index; i++ {
		assert(i >= 0 && i < len(VersionsData[version_index].Books), fmt.Sprintf("Invalid version book index (version: `%s`, book index: %d)", r.Version, i))
		book_data := VersionsData[version_index].Books[i]

		book_start := CreateFirstValidRef(r.Version, book_data.BookNumber)
		book_end := CreateRef(book_data.BookNumber, book_data.NumChapters, book_data.NumVerses[len(book_data.NumVerses)-1])
		assert(book_start.IsValid(r.Version), fmt.Sprintf("Invalid start ref for book created (version: `%s`, version book index: %d, created: %d)", r.Version, int(i), int(book_start)))
		assert(book_end.IsValid(r.Version), fmt.Sprintf("Invalid end ref for book created (version: `%s`, version book index: %d, created: %d)", r.Version, int(i), int(book_end)))

		ranges = append(ranges, ScriptureRange{
			Version: r.Version,
			Start:   book_start,
			End:     book_end,
		})
	}
	return ranges
}
