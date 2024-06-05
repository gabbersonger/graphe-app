package database

import (
	"fmt"
	"sync"
)

func getScriptureWord_ESV(g *GrapheDB, d *ScriptureWordData) {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go getScriptureWordBasicInfo_ESV(g, wg, d)
	go getScriptureWordStrongsInfo_ESV(g, wg, d)
	wg.Wait()
}

func getScriptureWordBasicInfo_ESV(g *GrapheDB, wg *sync.WaitGroup, d *ScriptureWordData) {
	db := <-g.Pool
	stmt := db.queries.esvWordBasicInfo
	err := stmt.Bind(int(d.Ref), int(d.WordNumber))
	g.check(err)

	hasRow, err := stmt.Step()
	g.check(err)
	if !hasRow {
		g.throw(fmt.Sprintf("Could not find row using query `esvWordBasicInfo` for (ref=%d, word_num=%d)", int(d.Ref), d.WordNumber))
	}
	var count int
	err = stmt.Scan(&(d.Text), &count)
	g.check(err)
	d.Fields["EnglishCount"] = fmt.Sprint(count)

	stmt.Reset()
	g.Pool <- db
	wg.Done()
}

func getScriptureWordStrongsInfo_ESV(g *GrapheDB, wg *sync.WaitGroup, d *ScriptureWordData) {
	db := <-g.Pool
	stmt := db.queries.esvWordStrongsInfo
	err := stmt.Bind(int(d.Ref), int(d.WordNumber))
	g.check(err)

	d.Collections = make([]ScriptureWordDataFields, 0, 1)
	var strong string
	var count int
	for {
		hasRow, err := stmt.Step()
		g.check(err)
		if !hasRow {
			break
		}
		err = stmt.Scan(&strong, &count)
		g.check(err)

		c := make(ScriptureWordDataFields)
		c["Strong"] = strong
		c["StrongCount"] = fmt.Sprint(count)
		d.Collections = append(d.Collections, c)
	}

	stmt.Reset()
	g.Pool <- db
	wg.Done()
}
