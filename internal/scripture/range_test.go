package scripture

import (
	. "graphe/internal"
	"testing"
)

func TestRangeIsValid(t *testing.T) {
	var r ScriptureRange

	Equals(t, false, r.IsValid())

	r = ScriptureRange{
		Start:   ScriptureRef(1_001_001),
		End:     ScriptureRef(1_050_026),
		Version: "esv",
	}
	Equals(t, true, r.IsValid())
	r.End = 1_001_031
	Equals(t, true, r.IsValid())

	// start equals end
	r.End = 1_001_001
	Equals(t, true, r.IsValid())

	// end before beginning
	r.Start = 1_001_002
	Equals(t, false, r.IsValid())

	// invalid start and end
	r.Start = 1_001_000
	r.End = 1_001_001
	Equals(t, false, r.IsValid())
	r.Start = 1_001_001
	r.End = 1_001_032
	Equals(t, false, r.IsValid())
	r.End = 1_051_001
	Equals(t, false, r.IsValid())
	r.Start = 1_049_999
	r.End = 1_050_026
	Equals(t, false, r.IsValid())

	// bridging book
	r.Start = 1_001_001
	r.End = 66_001_001
	Equals(t, true, r.IsValid())

	// version book order matters
	r.Version = "lxx"
	r.Start = 67_001_001
	r.End = 15_001_001
	Equals(t, true, r.IsValid())
	r.Start = 17_001_001
	r.End = 68_001_001
	Equals(t, true, r.IsValid())

	// prologue
	r.Start = 76_000_001
	r.End = 76_001_001
	Equals(t, true, r.IsValid())
	r.Start = 75_001_001
	r.End = 76_000_001
	Equals(t, true, r.IsValid())

	// superscripts
	r.Version = "esv"
	r.Start = 19_003_000
	r.End = 19_003_001
	Equals(t, true, r.IsValid())
	r.End = 19_005_000
	Equals(t, true, r.IsValid())
	r.Start = 19_002_001
	r.End = 19_003_000
	Equals(t, true, r.IsValid())

	// missing section
	r.Version = "lxx"
	r.Start = 80_001_001
	r.End = 80_001_007
	Equals(t, false, r.IsValid())
	r.Start = 80_001_001
	Equals(t, false, r.IsValid())
	r.Start = 80_001_006
	Equals(t, true, r.IsValid())
}

func TestRangeContains(t *testing.T) {
	var ran ScriptureRange
	var ref ScriptureRef

	// invalid
	Equals(t, false, ran.Contains(ref))
	ref = 0
	ran = ScriptureRange{
		Version: "esv",
		Start:   1_001_002,
		End:     1_001_004,
	}
	Equals(t, false, ran.Contains(ref))
	ref = 1_001_003
	ran.Start = 0
	Equals(t, false, ran.Contains(ref))
	ran.Start = 1_001_002
	ran.End = 0
	Equals(t, false, ran.Contains(ref))

	// between, start, end
	ran.End = 1_001_004
	Equals(t, true, ran.Contains(ref))
	ref = 1_001_002
	Equals(t, true, ran.Contains(ref))
	ref = 1_001_004
	Equals(t, true, ran.Contains(ref))

	// before, after
	ref = 1_001_001
	Equals(t, false, ran.Contains(ref))
	ref = 1_001_005
	Equals(t, false, ran.Contains(ref))

	// small range
	ran.Start = 1_001_001
	ran.End = 1_001_001
	ref = 1_001_001
	Equals(t, true, ran.Contains(ref))

	// superscript
	ran.Start = 19_003_000
	ran.End = 19_003_002
	ref = 19_003_001
	Equals(t, true, ran.Contains(ref))
	ran.Start = 19_002_001
	ref = 19_003_000
	Equals(t, true, ran.Contains(ref))
	ran.End = 19_003_000
	ref = 19_002_002
	Equals(t, true, ran.Contains(ref))

	// missing section
	ran.Version = "lxx"
	ran.Start = 80_001_008
	ran.End = 80_001_010
	ref = 80_001_009
	Equals(t, false, ran.Contains(ref))
	ran.Start = 80_001_007
	Equals(t, true, ran.Contains(ref))
	ref = 80_001_008
	Equals(t, false, ran.Contains(ref))
	ran.Start = 80_001_006
	ran.End = 80_001_008
	ref = 80_001_007
	Equals(t, false, ran.Contains(ref))

	// prologue (start, end, ref)
	ran.Start = 75_019_022
	ran.End = 76_001_001
	ref = 76_000_001
	Equals(t, true, ran.Contains(ref))
	ran.Start = 76_000_001
	ran.End = 76_001_001
	ref = 76_000_002
	Equals(t, true, ran.Contains(ref))
	ran.Start = 75_019_001
	ran.End = 76_000_001
	ref = 75_019_022
	Equals(t, true, ran.Contains(ref))

	// version book order
	ran.Start = 14_001_001
	ran.End = 15_001_001
	ref = 67_001_001
	Equals(t, true, ran.Contains(ref))
	ran.Start = 21_001_001
	ran.End = 18_001_001
	ref = 22_001_001
	Equals(t, true, ran.Contains(ref))
}
