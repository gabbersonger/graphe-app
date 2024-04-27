package app

type ScriptureRef int
type ScriptureVersion string

type ScriptureWord struct {
	WordNumber int    `json:"word_num"`
	Text       string `json:"text"`
	Pre        string `json:"pre"`
	Post       string `json:"post"`
}

type ScriptureVerse struct {
	Ref   ScriptureRef    `json:"ref"`
	Words []ScriptureWord `json:"words"`
}

type ScriptureRange struct {
	Version ScriptureVersion `json:"version"`
	Start   ScriptureRef     `json:"start"`
	End     ScriptureRef     `json:"end"`
}

type ScriptureBlock struct {
	Range  ScriptureRange   `json:"range"`
	Verses []ScriptureVerse `json:"verses"`
}

type ScriptureSection struct {
	Range  ScriptureRange   `json:"range"`
	Blocks []ScriptureBlock `json:"blocks"`
}

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
