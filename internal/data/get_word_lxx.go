package data

import (
	"fmt"
	"reflect"
	"sync"
)

func (d *DataDB) getScriptureWordLxx(wd *ScriptureWordData) {
	prepareScriptureWordLxx(wd)

	wg := new(sync.WaitGroup)
	wg.Add(3)
	go d.getScriptureWordLxx_Text(wg, wd)
	go d.getScriptureWordLxx_Info(wg, wd)
	go d.getScriptureWordLxx_InflectedCount(wg, wd)
	wg.Wait()

	d.checkScriptureWordLxx(wd)
}

func prepareScriptureWordLxx(wd *ScriptureWordData) {
	wd.Fields = make([]ScriptureWordDataField, 8)
	wd.Fields[0].Name = "Translit"
	wd.Fields[1].Name = "English"
	wd.Fields[2].Name = "Strong"
	wd.Fields[3].Name = "Grammar"
	wd.Fields[4].Name = "Form"
	wd.Fields[5].Name = "Gloss"
	wd.Fields[6].Name = "FormCount"
	wd.Fields[7].Name = "InflectedCount"
}

func (d *DataDB) checkScriptureWordLxx(wd *ScriptureWordData) {
	d.assert(len(wd.Text) > 0, fmt.Sprintf("Invalid text scanned (ref: %d, word_num: %d)", int(wd.Ref), int(wd.WordNumber)))
	d.assert(len(wd.Fields) == 8, fmt.Sprintf("Invalid fields length (ref: %d, word_num: %d, length: %d)", int(wd.Ref), int(wd.WordNumber), len(wd.Fields)))

	translit, ok := wd.Fields[0].Data.(string)
	d.assert(ok, fmt.Sprintf("Invalid translit type (ref: %d, word_num: %d, type: `%s`, value: %v)", int(wd.Ref), int(wd.WordNumber), reflect.TypeOf(wd.Fields[0].Data), wd.Fields[0].Data))
	d.assert(len(translit) > 0, fmt.Sprintf("Invalid translit value: empty (ref: %d, word_num: %d)", int(wd.Ref), int(wd.WordNumber)))

	english, ok := wd.Fields[1].Data.(string)
	d.assert(ok, fmt.Sprintf("Invalid english type (ref: %d, word_num: %d, type: `%s`, value: %v)", int(wd.Ref), int(wd.WordNumber), reflect.TypeOf(wd.Fields[1].Data), wd.Fields[1].Data))
	d.assert(len(english) > 0, fmt.Sprintf("Invalid english value: empty (ref: %d, word_num: %d)", int(wd.Ref), int(wd.WordNumber)))

	strong, ok := wd.Fields[2].Data.(string)
	d.assert(ok, fmt.Sprintf("Invalid strong type (ref: %d, word_num: %d, type: `%s`, value: %v)", int(wd.Ref), int(wd.WordNumber), reflect.TypeOf(wd.Fields[2].Data), wd.Fields[2].Data))
	d.assert(len(strong) > 0, fmt.Sprintf("Invalid strong value: empty (ref: %d, word_num: %d)", int(wd.Ref), int(wd.WordNumber)))

	grammar, ok := wd.Fields[3].Data.(string)
	d.assert(ok, fmt.Sprintf("Invalid grammar type (ref: %d, word_num: %d, type: `%s`, value: %v)", int(wd.Ref), int(wd.WordNumber), reflect.TypeOf(wd.Fields[3].Data), wd.Fields[3].Data))
	d.assert(len(grammar) > 0, fmt.Sprintf("Invalid grammar value: empty (ref: %d, word_num: %d)", int(wd.Ref), int(wd.WordNumber)))

	form, ok := wd.Fields[4].Data.(string)
	d.assert(ok, fmt.Sprintf("Invalid form type (ref: %d, word_num: %d, type: `%s`, value: %v)", int(wd.Ref), int(wd.WordNumber), reflect.TypeOf(wd.Fields[4].Data), wd.Fields[4].Data))
	d.assert(len(form) > 0, fmt.Sprintf("Invalid form value: empty (ref: %d, word_num: %d)", int(wd.Ref), int(wd.WordNumber)))

	gloss, ok := wd.Fields[5].Data.(string)
	d.assert(ok, fmt.Sprintf("Invalid gloss type (ref: %d, word_num: %d, type: `%s`, value: %v)", int(wd.Ref), int(wd.WordNumber), reflect.TypeOf(wd.Fields[5].Data), wd.Fields[5].Data))
	d.assert(len(gloss) > 0, fmt.Sprintf("Invalid gloss value: empty (ref: %d, word_num: %d)", int(wd.Ref), int(wd.WordNumber)))

	form_count, ok := wd.Fields[6].Data.(int)
	d.assert(ok, fmt.Sprintf("Invalid form_count type (ref: %d, word_num: %d, type: `%s`, value: %v)", int(wd.Ref), int(wd.WordNumber), reflect.TypeOf(wd.Fields[6].Data), wd.Fields[6].Data))
	d.assert(form_count >= 0, fmt.Sprintf("Invalid form_count value: below zero (ref: %d, word_num: %d, value: %d)", int(wd.Ref), int(wd.WordNumber), form_count))

	inflected_count, ok := wd.Fields[7].Data.(int)
	d.assert(ok, fmt.Sprintf("Invalid inflected_count type (ref: %d, word_num: %d, type: `%s`, value: %v)", int(wd.Ref), int(wd.WordNumber), reflect.TypeOf(wd.Fields[7].Data), wd.Fields[7].Data))
	d.assert(inflected_count >= 0, fmt.Sprintf("Invalid inflected_count value: below zero (ref: %d, word_num: %d, value: %d)", int(wd.Ref), int(wd.WordNumber), inflected_count))
}

func (d *DataDB) getScriptureWordLxx_Text(wg *sync.WaitGroup, wd *ScriptureWordData) {
	db := <-d.pool
	stmt := db.queries.lxxWordText
	err := stmt.Bind(int(wd.Ref), int(wd.WordNumber))
	d.assert(err == nil, fmt.Sprintf("Error binding ref and word_num to sql query (ref: %d, word_num: %d) (err: %v)", int(wd.Ref), int(wd.WordNumber), err))

	has_row, err := stmt.Step()
	d.assert(err == nil, fmt.Sprintf("Error stepping sql query (ref: %d, word_num: %d) (err: %v)", int(wd.Ref), int(wd.WordNumber), err))
	d.assert(has_row, fmt.Sprintf("Error finding db row for word data (ref: %d, word_num: %d)", int(wd.Ref), int(wd.WordNumber)))

	err = stmt.Scan(&(wd.Text))
	d.assert(err == nil, fmt.Sprintf("Error scanning sql query (ref: %d, word_num: %d) (err: %v)", int(wd.Ref), int(wd.WordNumber), err))

	err = stmt.Reset()
	d.assert(err == nil, fmt.Sprintf("Error reseting sql statement (err: %v)", err))
	d.pool <- db
	wg.Done()
}

func (d *DataDB) getScriptureWordLxx_Info(wg *sync.WaitGroup, wd *ScriptureWordData) {
	db := <-d.pool
	stmt := db.queries.lxxWordBasicInfo
	err := stmt.Bind(int(wd.Ref), int(wd.WordNumber))
	d.assert(err == nil, fmt.Sprintf("Error binding ref and word_num to sql query (ref: %d, word_num: %d) (err: %v)", int(wd.Ref), int(wd.WordNumber), err))

	has_row, err := stmt.Step()
	d.assert(err == nil, fmt.Sprintf("Error stepping sql query (ref: %d, word_num: %d) (err: %v)", int(wd.Ref), int(wd.WordNumber), err))
	d.assert(has_row, fmt.Sprintf("Error finding db row for word data (ref: %d, word_num: %d)", int(wd.Ref), int(wd.WordNumber)))

	var translit, english, strong, grammar, form, gloss string
	var count int
	err = stmt.Scan(&translit, &english, &strong, &grammar, &form, &gloss, &count)
	d.assert(err == nil, fmt.Sprintf("Error scanning sql query (ref: %d, word_num: %d) (err: %v)", int(wd.Ref), int(wd.WordNumber), err))
	wd.Fields[0].Data = translit
	wd.Fields[1].Data = english
	wd.Fields[2].Data = strong
	wd.Fields[3].Data = grammar
	wd.Fields[4].Data = form
	wd.Fields[5].Data = gloss
	wd.Fields[6].Data = count

	err = stmt.Reset()
	d.assert(err == nil, fmt.Sprintf("Error reseting sql statement (err: %v)", err))
	d.pool <- db
	wg.Done()
}

func (d *DataDB) getScriptureWordLxx_InflectedCount(wg *sync.WaitGroup, wd *ScriptureWordData) {
	db := <-d.pool
	stmt := db.queries.lxxWordInflectedCount
	err := stmt.Bind(int(wd.Ref), int(wd.WordNumber))
	d.assert(err == nil, fmt.Sprintf("Error binding ref and word_num to sql query (ref: %d, word_num: %d) (err: %v)", int(wd.Ref), int(wd.WordNumber), err))

	has_row, err := stmt.Step()
	d.assert(err == nil, fmt.Sprintf("Error stepping sql query (ref: %d, word_num: %d) (err: %v)", int(wd.Ref), int(wd.WordNumber), err))
	d.assert(has_row, fmt.Sprintf("Error finding db row for word data (ref: %d, word_num: %d)", int(wd.Ref), int(wd.WordNumber)))

	var count int
	err = stmt.Scan(&count)
	d.assert(err == nil, fmt.Sprintf("Error scanning sql query (ref: %d, word_num: %d) (err: %v)", int(wd.Ref), int(wd.WordNumber), err))
	wd.Fields[7].Data = count

	err = stmt.Reset()
	d.assert(err == nil, fmt.Sprintf("Error reseting sql statement (err: %v)", err))
	d.pool <- db
	wg.Done()
}
