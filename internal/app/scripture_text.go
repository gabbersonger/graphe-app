package app

type ScriptureRef int
type ScriptureWord string
type ScriptureVersion string

type ScriptureVerse struct {
	Ref   ScriptureRef    `json:"ref"`
	Words []ScriptureWord `json:"words"`
}

type ScriptureRange struct {
	Start ScriptureRef `json:"start"`
	End   ScriptureRef `json:"end"`
}

type ScriptureBlock struct {
	Range  ScriptureRange   `json:"range"`
	Verses []ScriptureVerse `json:"verses"`
}

type ScriptureSection struct {
	Version ScriptureVersion `json:"version"`
	Range   ScriptureRange   `json:"range"`
	Blocks  []ScriptureBlock `json:"blocks"`
}

func getScriptureSection(a *App, out chan *ScriptureSection, ver ScriptureVersion, ran ScriptureRange) {
	s := ScriptureSection{Version: ver, Range: ran}
	s.Blocks = make([]ScriptureBlock, 0, 1)

	db := <-a.db_pool
	stmt, err := db.getQuery("GetScriptureSection")
	a.check(err)
	err = stmt.Bind(int(ran.Start), int(ran.End))
	a.check(err)

	var ref, word_num int
	var text string
	createNextBlock := true
	lastRef := 0

	for {
		hasRow, err := stmt.Step()
		a.check(err)
		if !hasRow {
			break
		}
		err = stmt.Scan(&ref, &word_num, &text)
		a.check(err)

		// Add block if needed
		if createNextBlock {
			createNextBlock = false
			if len(s.Blocks) > 0 {
				s.Blocks[len(s.Blocks)-1].Range.End = ScriptureRef(lastRef)
			}

			newBlock := ScriptureBlock{}
			newBlock.Range.Start = ScriptureRef(ref)
			newBlock.Verses = make([]ScriptureVerse, 0, 20) // TODO: pick the right value
			s.Blocks = append(s.Blocks, newBlock)
		}
		lastBlock := len(s.Blocks) - 1

		// Add verse if needed
		if lastRef != ref || len(s.Blocks[lastBlock].Verses) == 0 {
			lastRef = ref

			newVerse := ScriptureVerse{}
			newVerse.Ref = ScriptureRef(ref)
			newVerse.Words = make([]ScriptureWord, 0, 40) // TODO: pick the right value
			s.Blocks[lastBlock].Verses = append(s.Blocks[lastBlock].Verses, newVerse)
		}
		lastVerse := len(s.Blocks[lastBlock].Verses) - 1

		// Add word
		newWord := ScriptureWord(text)
		s.Blocks[lastBlock].Verses[lastVerse].Words = append(s.Blocks[lastBlock].Verses[lastVerse].Words, newWord)

		runes := []rune(text)
		max := len(runes) - 1
		for i := range runes {
			if runes[max-i] == '¶' {
				createNextBlock = true
			}
		}
	}
	s.Range.End = ScriptureRef(ref)
	s.Blocks[len(s.Blocks)-1].Range.End = ScriptureRef(ref)

	stmt.Reset()
	a.db_pool <- db
	out <- &s
}

func (a *App) GetScriptureSections(ver ScriptureVersion, ran []ScriptureRange) []ScriptureSection {
	feed := make(chan *ScriptureSection, len(ran))
	for _, r := range ran {
		go getScriptureSection(a, feed, ver, r)
	}

	data := make([]ScriptureSection, 0, len(ran))
	for i := 0; i < len(ran); i++ {
		data = append(data, *(<-feed))
	}
	return data
}
