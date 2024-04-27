package app

type ScriptureWordData_Dictionary struct {
	Form  string `json:"form"`
	Gloss string `json:"gloss"`
}

type ScriptureWordData_Strongs struct {
	Num     string `json:"num"`
	Grammar string `json:"grammar"`
}

type ScriptureWordData struct {
	Ref            ScriptureRef                   `json:"ref"`
	WordNumber     int                            `json:"word_number"`
	Text           string                         `json:"text"`
	Translit       string                         `json:"translit"`
	English        string                         `json:"english"`
	Dictionary     []ScriptureWordData_Dictionary `json:"dictionary"`
	Strongs        []ScriptureWordData_Strongs    `json:"strongs"`
	InflectedCount int                            `json:"inflected_count"`
	LexemeCount    int                            `json:"lexeme_count"`
}

func (a *App) GetScriptureWord(version ScriptureVersion, ref ScriptureRef, word_num int) ScriptureWordData {
	w := ScriptureWordData{Ref: ref, WordNumber: word_num}
	switch version {
	case "gnt":
		getGNTScriptureWord(a, &w)
	case "lxx":
		getLXXScriptureWord(a, &w)
	default:
		a.Throw("Unknown version (" + string(version) + ") passed to GetScriptureSections")
	}
	return w
}
