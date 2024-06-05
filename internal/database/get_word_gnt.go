package database

import (
	"fmt"
	"sync"
)

func getScriptureWord_GNT(g *GrapheDB, d *ScriptureWordData) {
	wg := new(sync.WaitGroup)
	wg.Add(4)
	go getScriptureWordText_GNT(g, wg, d)
	go getScriptureWordBasicInfo_GNT(g, wg, d)
	go getScriptureWordDictionaryValues_GNT(g, wg, d)
	go getScriptureWordInflectedCount_GNT(g, wg, d)
	wg.Wait()
}

func getScriptureWordText_GNT(g *GrapheDB, wg *sync.WaitGroup, d *ScriptureWordData) {
	db := <-g.Pool
	stmt := db.queries.gntWordText
	err := stmt.Bind(int(d.Ref), int(d.WordNumber))
	g.check(err)

	hasRow, err := stmt.Step()
	g.check(err)
	if !hasRow {
		g.throw(fmt.Sprintf("Could not find row using query `gntWordText` for (ref=%d, word_num=%d)", int(d.Ref), d.WordNumber))
	}
	err = stmt.Scan(&(d.Text))
	g.check(err)

	stmt.Reset()
	g.Pool <- db
	wg.Done()
}

func getScriptureWordBasicInfo_GNT(g *GrapheDB, wg *sync.WaitGroup, d *ScriptureWordData) {
	db := <-g.Pool
	stmt := db.queries.gntWordBasicInfo
	err := stmt.Bind(int(d.Ref), int(d.WordNumber))
	g.check(err)

	hasRow, err := stmt.Step()
	g.check(err)
	if !hasRow {
		g.throw(fmt.Sprintf("Could not find row using query `gntWordBasicInfo` for (ref=%d, word_num=%d)", int(d.Ref), d.WordNumber))
	}
	var translit, english string
	err = stmt.Scan(&translit, &english)
	g.check(err)
	d.Fields["Translit"] = translit
	d.Fields["English"] = english

	stmt.Reset()
	g.Pool <- db
	wg.Done()
}

func getScriptureWordDictionaryValues_GNT(g *GrapheDB, wg *sync.WaitGroup, d *ScriptureWordData) {
	db := <-g.Pool
	stmt := db.queries.gntWordDictionary
	err := stmt.Bind(int(d.Ref), int(d.WordNumber))
	g.check(err)

	d.Collections = make([]ScriptureWordDataFields, 0, 1)
	var form, gloss, strong, grammar string
	var count int
	for {
		hasRow, err := stmt.Step()
		g.check(err)
		if !hasRow {
			break
		}
		err = stmt.Scan(&form, &gloss, &strong, &grammar, &count)
		g.check(err)

		c := make(ScriptureWordDataFields)
		c["Form"] = form
		c["Gloss"] = gloss
		c["Strong"] = strong
		c["Grammar"] = grammar
		c["FormCount"] = fmt.Sprint(count)
		d.Collections = append(d.Collections, c)
	}

	stmt.Reset()
	g.Pool <- db
	wg.Done()
}

func getScriptureWordInflectedCount_GNT(g *GrapheDB, wg *sync.WaitGroup, d *ScriptureWordData) {
	db := <-g.Pool
	stmt := db.queries.gntWordInflectedCount
	err := stmt.Bind(int(d.Ref), int(d.WordNumber))
	g.check(err)

	hasRow, err := stmt.Step()
	g.check(err)
	if !hasRow {
		g.throw(fmt.Sprintf("Could not find row using query `gntWordInflectedCount` for (ref=%d, word_num=%d)", int(d.Ref), d.WordNumber))
	}
	var count int
	err = stmt.Scan(&count)
	g.check(err)
	d.Fields["InflectedCount"] = fmt.Sprint(count)

	stmt.Reset()
	g.Pool <- db
	wg.Done()
}
