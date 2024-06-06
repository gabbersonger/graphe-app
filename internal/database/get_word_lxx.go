package database

import (
	"fmt"
	"sync"
)

func getScriptureWord_LXX(g *GrapheDB, d *ScriptureWordData) {
	prepareScriptureWord_LXX(d)
	wg := new(sync.WaitGroup)
	wg.Add(3)
	go getScriptureWordText_LXX(g, wg, d)
	go getScriptureWordInfo_LXX(g, wg, d)
	go getScriptureWordInflectedCount_LXX(g, wg, d)
	wg.Wait()
}

func prepareScriptureWord_LXX(d *ScriptureWordData) {
	d.Fields = make([]ScriptureWordDataField, 8)
	d.Fields[0].Name = "Translit"
	d.Fields[1].Name = "English"
	d.Fields[2].Name = "Strong"
	d.Fields[3].Name = "Grammar"
	d.Fields[4].Name = "Form"
	d.Fields[5].Name = "Gloss"
	d.Fields[6].Name = "FormCount[int]"
	d.Fields[7].Name = "InflectedCount[int]"
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
	d.Fields[0].Data = translit
	d.Fields[1].Data = english
	d.Fields[2].Data = strong
	d.Fields[3].Data = grammar
	d.Fields[4].Data = form
	d.Fields[5].Data = gloss
	d.Fields[6].Data = fmt.Sprint(count)

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
	d.Fields[7].Data = fmt.Sprint(count)

	stmt.Reset()
	g.Pool <- db
	wg.Done()
}
