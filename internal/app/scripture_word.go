package app

func (a *App) GetScriptureWord(version ScriptureVersion, ref ScriptureRef, word_num int) ScriptureWordData {
	w := ScriptureWordData{Ref: ref, WordNumber: word_num}
	switch version {
	case "gnt":
		getGNTScriptureWord(a, &w)
	case "lxx":
		getLXXScriptureWord(a, &w)
	case "esv":
		getESVScriptureWord(a, &w)
	default:
		a.Throw("Unknown version (" + string(version) + ") passed to GetScriptureSections")
	}
	return w
}
