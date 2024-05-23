package app

type ScriptureRef int
type ScriptureVersion string

type ScriptureWordDetailType int

const (
	Footnote ScriptureWordDetailType = iota
	Crossref ScriptureWordDetailType = iota
)

type ScriptureWordDetail struct {
	Position bool                    `json:"position"`
	Type     ScriptureWordDetailType `json:"type"`
	Data     string                  `json:"data"`
}

type ScriptureVerseDetailType int

const (
	Title   ScriptureVerseDetailType = iota
	Heading ScriptureVerseDetailType = iota
)

type ScriptureVerseDetail struct {
	Type ScriptureVerseDetailType `json:"type"`
	Data string                   `json:"data"`
}

type ScriptureWord struct {
	WordNumber       int                   `json:"word_num"`
	Text             string                `json:"text"`
	Pre              string                `json:"pre"`
	Post             string                `json:"post"`
	Details          []ScriptureWordDetail `json:"details,omitempty"`
	NoInstantDetails bool                  `json:"no_instant_details,omitempty"`
}

type ScriptureVerse struct {
	Ref          ScriptureRef           `json:"ref"`
	Words        []ScriptureWord        `json:"words"`
	Details      []ScriptureVerseDetail `json:"details,omitempty"`
	Continuation bool                   `json:"continuation,omitempty"`
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
	Form    string `json:"form"`
	Gloss   string `json:"gloss"`
	Strong  string `json:"strong"`
	Grammar string `json:"grammar"`
	Count   int    `json:"count"`
}

type ScriptureWordData struct {
	Ref            ScriptureRef                   `json:"ref"`
	WordNumber     int                            `json:"word_number"`
	Text           string                         `json:"text"`
	Translit       string                         `json:"translit"`
	English        string                         `json:"english"`
	Dictionary     []ScriptureWordData_Dictionary `json:"dictionary"`
	InflectedCount int                            `json:"inflected_count"`
}
