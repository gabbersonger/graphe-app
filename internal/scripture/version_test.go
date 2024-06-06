package scripture_test

import (
	. "graphe/internal"
	. "graphe/internal/scripture"
	"testing"
)

func TestVersionIsValid(t *testing.T) {
	Assert(t, ScriptureVersion("esv").IsValid(), "'esv' is an invalid version")
	Assert(t, ScriptureVersion("gnt").IsValid(), "'gnt' is an invalid version")
	Assert(t, ScriptureVersion("lxx").IsValid(), "'lxx' is an invalid version")
	Assert(t, ScriptureVersion("hot").IsValid(), "'hot' is an invalid version")
	Assert(t, !ScriptureVersion("asd").IsValid(), "'asd' is a valid version")
	Assert(t, !ScriptureVersion("").IsValid(), "'asd' is a valid version")
	Assert(t, !ScriptureVersion("1").IsValid(), "'1' is a valid version")
	Assert(t, !ScriptureVersion("esv-").IsValid(), "'esv-' is a valid version")
	Assert(t, !ScriptureVersion("esvgnt").IsValid(), "'esvgnt' is a valid version")
}

func TestGetVersionIndex(t *testing.T) {
	var v int
	var err error

	v, err = GetVersionIndex(ScriptureVersion("esv"))
	Ok(t, err)
	Equals(t, 3, v)
	v, err = GetVersionIndex(ScriptureVersion("gnt"))
	Ok(t, err)
	Equals(t, 1, v)
	v, err = GetVersionIndex(ScriptureVersion("lxx"))
	Ok(t, err)
	Equals(t, 0, v)
	v, err = GetVersionIndex(ScriptureVersion("hot"))
	Ok(t, err)
	Equals(t, 2, v)
	v, err = GetVersionIndex(ScriptureVersion("asd"))
	Assert(t, err != nil, "'asd' was used as version and didn't return error")
	v, err = GetVersionIndex(ScriptureVersion(""))
	Assert(t, err != nil, "'' was used as version and didn't return error")
	v, err = GetVersionIndex(ScriptureVersion("1"))
	Assert(t, err != nil, "'1' was used as version and didn't return error")
	v, err = GetVersionIndex(ScriptureVersion("esv-"))
	Assert(t, err != nil, "'esv-' was used as version and didn't return error")
	v, err = GetVersionIndex(ScriptureVersion("esvgnt"))
	Assert(t, err != nil, "'esvgnt' was used as version and didn't return error")
}

func TestGetVersionBookIndex(t *testing.T) {
	var i int
	var err error

	esv := ScriptureVersion("esv")
	lxx := ScriptureVersion("lxx")

	i, err = GetVersionBookIndex(esv, 1)
	Ok(t, err)
	Equals(t, 0, i)
	i, err = GetVersionBookIndex(lxx, 1)
	Ok(t, err)
	Equals(t, 0, i)

	i, err = GetVersionBookIndex(esv, 15)
	Ok(t, err)
	Equals(t, 14, i)
	i, err = GetVersionBookIndex(lxx, 67)
	Ok(t, err)
	Equals(t, 14, i)

	i, err = GetVersionBookIndex(esv, 16)
	Ok(t, err)
	Equals(t, 15, i)
	i, err = GetVersionBookIndex(lxx, 15)
	Ok(t, err)
	Equals(t, 15, i)
}
