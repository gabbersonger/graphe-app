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
