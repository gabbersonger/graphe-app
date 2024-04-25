package app

import "sync"

func getLXXScriptureSection(a *App, wg *sync.WaitGroup, s *ScriptureSection) {
	s.Blocks = make([]ScriptureBlock, 0, 1)

	db := <-a.db.pool
	stmt, err := db.getQuery("GetLXXScriptureSection")
	a.check(err)
	err = stmt.Bind(int(s.Range.Start), int(s.Range.End))
	a.check(err)

	var ref, word_num int
	var text, pre, post string
	createNextBlock := true
	lastRef := 0

	for {
		hasRow, err := stmt.Step()
		a.check(err)
		if !hasRow {
			break
		}
		err = stmt.Scan(&ref, &word_num, &text, &pre, &post)
		a.check(err)

		// TODO: temporary fix (breaks at chapters) until paragraph breaks included in data
		// if len(s.Blocks)-1 >= 0 {
		// 	last_block_start := s.Blocks[len(s.Blocks)-1].Range.Start
		// 	last_block_chapter := last_block_start - last_block_start%1000
		// 	this_chapter := ref - ref%1000

		// 	if this_chapter != int(last_block_chapter) {
		// 		createNextBlock = true
		// 	}
		// }

		// Add block if needed
		if createNextBlock {
			createNextBlock = false
			if len(s.Blocks) > 0 {
				s.Blocks[len(s.Blocks)-1].Range.End = ScriptureRef(lastRef)
			}

			newBlock := ScriptureBlock{}
			newBlock.Range.Version = s.Range.Version
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

		// Check for paragraph break
		n := -1
		runes := []rune(post)
		for i, rune := range runes {
			if rune == '¶' {
				createNextBlock = true
				n = i
				break
			}
		}
		if n >= 0 {
			runes = append(runes[:n], runes[n+1:]...)
			post = string(runes)
		}

		// Add word
		newWord := ScriptureWord{word_num, text, pre, post}
		s.Blocks[lastBlock].Verses[lastVerse].Words = append(s.Blocks[lastBlock].Verses[lastVerse].Words, newWord)
	}
	s.Range.End = ScriptureRef(ref)
	s.Blocks[len(s.Blocks)-1].Range.End = ScriptureRef(ref)

	stmt.Reset()
	a.db.pool <- db
	wg.Done()
}

func getLXXScriptureWord(a *App, w *ScriptureWordData) {
	// TODO
}
