package database

import (
	"fmt"
	"sync"
)

func getScriptureWord_LXX(g *GrapheDB, d *ScriptureWordData) {
	wg := new(sync.WaitGroup)
	wg.Add(3)
	go getScriptureWordText_LXX(g, wg, d)
	go getScriptureWordInfo_LXX(g, wg, d)
	go getScriptureWordInflectedCount_LXX(g, wg, d)
	wg.Wait()
}

func getScriptureWordText_LXX(g *GrapheDB, wg *sync.WaitGroup, d *ScriptureWordData) {
	db := <-g.Pool
	stmt := db.queries.lxxWordText
	err := stmt.Bind(int(d.Ref), int(d.WordNumber))
	g.check(err)

	hasRow, err := stmt.Step()
	g.check(err)
	if !hasRow {
		g.throw(fmt.Sprintf("Could not find row using query `lxxWordText` for (ref=%d, word_num=%d)", int(d.Ref), d.WordNumber))
	}
	err = stmt.Scan(&(d.Text))
	g.check(err)

	stmt.Reset()
	g.Pool <- db
	wg.Done()
}

func getScriptureWordInfo_LXX(g *GrapheDB, wg *sync.WaitGroup, d *ScriptureWordData) {
	db := <-g.Pool
	stmt := db.queries.lxxWordBasicInfo
	err := stmt.Bind(int(d.Ref), int(d.WordNumber))
	g.check(err)

	hasRow, err := stmt.Step()
	g.check(err)
	if !hasRow {
		g.throw(fmt.Sprintf("Could not find row using query `lxxWordBasicInfo` for (ref=%d, word_num=%d)", int(d.Ref), d.WordNumber))
	}
	var translit, english, strong, grammar, form, gloss string
	var count int
	err = stmt.Scan(&translit, &english, &strong, &grammar, &form, &gloss, &count)
	g.check(err)
	d.Fields["Translit"] = translit
	d.Fields["English"] = english
	d.Fields["Strong"] = strong
	d.Fields["Grammar"] = grammar
	d.Fields["Form"] = form
	d.Fields["Gloss"] = gloss
	d.Fields["FormCount"] = fmt.Sprint(count)

	stmt.Reset()
	g.Pool <- db
	wg.Done()
}

func getScriptureWordInflectedCount_LXX(g *GrapheDB, wg *sync.WaitGroup, d *ScriptureWordData) {
	db := <-g.Pool
	stmt := db.queries.lxxWordInflectedCount
	err := stmt.Bind(int(d.Ref), int(d.WordNumber))
	g.check(err)

	hasRow, err := stmt.Step()
	g.check(err)
	if !hasRow {
		g.throw(fmt.Sprintf("Could not find row using query `lxxWordInflectedCount` for (ref=%d, word_num=%d)", int(d.Ref), d.WordNumber))
	}
	var count int
	err = stmt.Scan(&count)
	g.check(err)
	d.Fields["InflectedCount"] = fmt.Sprint(count)

	stmt.Reset()
	g.Pool <- db
	wg.Done()
}
