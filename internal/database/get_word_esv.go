package database

import (
	"fmt"
	"sync"
)

func getScriptureWord_ESV(g *GrapheDB, d *ScriptureWordData) {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go getScriptureWordText_ESV(g, wg, d)
	wg.Wait()
}

func getScriptureWordText_ESV(g *GrapheDB, wg *sync.WaitGroup, d *ScriptureWordData) {
	db := <-g.Pool
	stmt := db.queries.esvWordText
	err := stmt.Bind(int(d.Ref), int(d.WordNumber))
	g.check(err)

	hasRow, err := stmt.Step()
	g.check(err)
	if !hasRow {
		g.throw(fmt.Sprintf("Could not find row using query `esvWordText` for (ref=%d, word_num=%d)", int(d.Ref), d.WordNumber))
	}
	err = stmt.Scan(&(d.Text))
	g.check(err)

	stmt.Reset()
	g.Pool <- db
	wg.Done()
}
