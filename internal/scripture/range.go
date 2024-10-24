package scripture

import "fmt"

type ScriptureRange struct {
	Version ScriptureVersion `json:"version"`
	Start   ScriptureRef     `json:"start"`
	End     ScriptureRef     `json:"end"`
}

func (s *ScriptureService) IsRangeValid(rang ScriptureRange) bool {
	s.assert(s.IsRefValid(rang.Start, rang.Version), fmt.Sprintf("Invalid start to range (version: `%s`, start: %d, end: %d)", rang.Version, int(rang.Start), int(rang.End)))
	s.assert(s.IsRefValid(rang.End, rang.Version), fmt.Sprintf("Invalid end to range (version: `%s`, start: %d, end: %d)", rang.Version, int(rang.Start), int(rang.End)))

	start_vb_index := s.getVersionBookIndex(rang.Version, s.GetRefBook(rang.Start))
	end_vb_index := s.getVersionBookIndex(rang.Version, s.GetRefBook(rang.End))
	s.assert(start_vb_index >= 0, fmt.Sprintf("Error getting start's version book index (ref: %d, val: %d)", int(rang.Start), start_vb_index))
	s.assert(end_vb_index >= 0, fmt.Sprintf("Error getting start's version book index (ref: %d, val: %d)", int(rang.End), end_vb_index))

	if start_vb_index < end_vb_index {
		return true
	}
	return start_vb_index == end_vb_index && rang.Start <= rang.End
}

func (s *ScriptureService) RangeContains(rang ScriptureRange, ref ScriptureRef) bool {
	s.assert(s.IsRangeValid(rang), fmt.Sprintf("Invalid range (version: `%s`, start: %d, end: %d)", rang.Version, int(rang.Start), int(rang.End)))
	s.assert(s.IsRefValid(ref, rang.Version), fmt.Sprintf("Invalid ref for version (version: `%s`, ref: %d)", rang.Version, int(ref)))

	start_vb_index := s.getVersionBookIndex(rang.Version, s.GetRefBook(rang.Start))
	end_vb_index := s.getVersionBookIndex(rang.Version, s.GetRefBook(rang.End))
	ref_vb_index := s.getVersionBookIndex(rang.Version, s.GetRefBook(ref))
	s.assert(start_vb_index >= 0, fmt.Sprintf("Error getting start's version book index (version: `%s`, ref: %d, val: %d)", rang.Version, int(rang.Start), start_vb_index))
	s.assert(end_vb_index >= 0, fmt.Sprintf("Error getting end's version book index (version: `%s`, ref: %d, val: %d)", rang.Version, int(rang.End), end_vb_index))
	s.assert(ref_vb_index >= 0, fmt.Sprintf("Error getting ref's version book index (version: `%s`, ref: %d, val: %d)", rang.Version, int(ref), ref_vb_index))

	if start_vb_index == end_vb_index { // start = end (must be between)
		return ref >= rang.Start && ref <= rang.End
	} else if ref_vb_index == start_vb_index { // start = ref (must be after)
		return ref >= rang.Start
	} else if ref_vb_index == end_vb_index { // ref = end (must be before)
		return ref <= rang.End
	}
	return ref_vb_index > start_vb_index && ref_vb_index < end_vb_index // start =/= end (must be between)
}

func (s *ScriptureService) DivideIntoBookRanges(rang ScriptureRange) []ScriptureRange {
	s.assert(s.IsRangeValid(rang), fmt.Sprintf("Invalid range (version: `%s`, start: %d, end: %d)", rang.Version, int(rang.Start), int(rang.End)))

	version_index := s.getVersionIndex(rang.Version)
	s.assert(version_index >= 0 && version_index < len(VersionsData), fmt.Sprintf("Error getting version index (version: `%s`, val: %d)", rang.Version, version_index))

	start_vb_index := s.getVersionBookIndex(rang.Version, s.GetRefBook(rang.Start))
	end_vb_index := s.getVersionBookIndex(rang.Version, s.GetRefBook(rang.End))
	s.assert(start_vb_index >= 0, fmt.Sprintf("Error getting start's version book index (version: `%s`, ref: %d, val: %d)", rang.Version, int(rang.Start), start_vb_index))
	s.assert(end_vb_index >= 0, fmt.Sprintf("Error getting end's version book index (version: `%s`, ref: %d, val: %d)", rang.Version, int(rang.End), end_vb_index))
	s.assert(start_vb_index <= end_vb_index, fmt.Sprintf("Version book index misordering: start before end (start: %d, end %d)", start_vb_index, end_vb_index))

	ranges := make([]ScriptureRange, 0, end_vb_index-start_vb_index+1)
	for i := start_vb_index; i <= end_vb_index; i++ {
		s.assert(i >= 0 && i < len(VersionsData[version_index].Books), fmt.Sprintf("Invalid version book index (version: `%s`, book index: %d)", rang.Version, i))
		book_data := VersionsData[version_index].Books[i]

		book_start := s.CreateFirstValidRef(rang.Version, book_data.BookNumber)
		book_end := s.CreateLastValidRef(rang.Version, book_data.BookNumber)
		s.assert(s.IsRefValid(book_start, rang.Version), fmt.Sprintf("Invalid start ref for book created (version: `%s`, version book index: %d, created: %d)", rang.Version, int(i), int(book_start)))
		s.assert(s.IsRefValid(book_end, rang.Version), fmt.Sprintf("Invalid end ref for book created (version: `%s`, version book index: %d, created: %d)", rang.Version, int(i), int(book_end)))

		if i == start_vb_index {
			book_start = rang.Start
		}
		if i == end_vb_index {
			book_end = rang.End
		}

		ranges = append(ranges, ScriptureRange{
			Version: rang.Version,
			Start:   book_start,
			End:     book_end,
		})
	}
	return ranges
}
