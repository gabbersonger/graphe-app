package database

import (
	"fmt"
	"sync"
)

func getScriptureWord_GNT(g *GrapheDB, d *ScriptureWordData) {
	prepareScriptureWord_GNT(d)
	wg := new(sync.WaitGroup)
	wg.Add(4)
	go getScriptureWordText_GNT(g, wg, d)
	go getScriptureWordBasicInfo_GNT(g, wg, d)
	go getScriptureWordDictionaryValues_GNT(g, wg, d)
	go getScriptureWordInflectedCount_GNT(g, wg, d)
	wg.Wait()
}

func prepareScriptureWord_GNT(d *ScriptureWordData) {
	d.Fields = make([]ScriptureWordDataField, 3)
	d.Fields[0].Name = "Translit"
	d.Fields[1].Name = "English"
	d.Fields[2].Name = "InflectedCount[int]"
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
	d.Fields[0].Data = translit
	d.Fields[1].Data = english

	stmt.Reset()
	g.Pool <- db
	wg.Done()
}

func getScriptureWordDictionaryValues_GNT(g *GrapheDB, wg *sync.WaitGroup, d *ScriptureWordData) {
	db := <-g.Pool
	stmt := db.queries.gntWordDictionary
	err := stmt.Bind(int(d.Ref), int(d.WordNumber))
	g.check(err)

	d.Collections = make([][]ScriptureWordDataField, 0)
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

		fields := make([]ScriptureWordDataField, 5)
		fields[0].Name = "Form"
		fields[0].Data = form
		fields[1].Name = "Gloss"
		fields[1].Data = gloss
		fields[2].Name = "Strong"
		fields[2].Data = strong
		fields[3].Name = "Grammar"
		fields[3].Data = grammar
		fields[4].Name = "FormCount[int]"
		fields[4].Data = fmt.Sprint(count)
		d.Collections = append(d.Collections, fields)
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
	d.Fields[2].Data = fmt.Sprint(count)

	stmt.Reset()
	g.Pool <- db
	wg.Done()
}
