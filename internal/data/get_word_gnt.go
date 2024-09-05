package data

import (
	"fmt"
	"reflect"
	"sync"
)

func (d *DataDB) getScriptureWordGnt(wd *ScriptureWordData) {
	prepareScriptureWordGnt(wd)

	wg := new(sync.WaitGroup)
	wg.Add(4)
	go d.getScriptureWordGnt_Text(wg, wd)
	go d.getScriptureWordGnt_BasicInfo(wg, wd)
	go d.getScriptureWordGnt_InflectedCount(wg, wd)
	go d.getScriptureWordGnt_Dictionary(wg, wd)
	wg.Wait()

	d.checkScriptureWordGnt(wd)
}

func prepareScriptureWordGnt(wd *ScriptureWordData) {
	wd.Fields = make([]ScriptureWordDataField, 4)
	wd.Fields[0].Name = "Translit"
	wd.Fields[1].Name = "English"
	wd.Fields[2].Name = "InflectedCount"
	wd.Fields[3].Name = "Dictionary"
}

func (d *DataDB) checkScriptureWordGnt(wd *ScriptureWordData) {
	d.assert(len(wd.Text) > 0, fmt.Sprintf("Invalid text scanned (ref: %d, word_num: %d)", int(wd.Ref), int(wd.WordNumber)))
	d.assert(len(wd.Fields) == 4, fmt.Sprintf("Invalid fields length (ref: %d, word_num: %d, length: %d)", int(wd.Ref), int(wd.WordNumber), len(wd.Fields)))

	translit, ok := wd.Fields[0].Data.(string)
	d.assert(ok, fmt.Sprintf("Invalid translit type (ref: %d, word_num: %d, type: `%s`, value: %v)", int(wd.Ref), int(wd.WordNumber), reflect.TypeOf(wd.Fields[0].Data), wd.Fields[0].Data))
	d.assert(len(translit) > 0, fmt.Sprintf("Invalid translit value: empty (ref: %d, word_num: %d)", int(wd.Ref), int(wd.WordNumber)))

	english, ok := wd.Fields[1].Data.(string)
	d.assert(ok, fmt.Sprintf("Invalid english type (ref: %d, word_num: %d, type: `%s`, value: %v)", int(wd.Ref), int(wd.WordNumber), reflect.TypeOf(wd.Fields[1].Data), wd.Fields[1].Data))
	d.assert(len(english) > 0, fmt.Sprintf("Invalid english value: empty (ref: %d, word_num: %d)", int(wd.Ref), int(wd.WordNumber)))

	inflected_count, ok := wd.Fields[2].Data.(int)
	d.assert(ok, fmt.Sprintf("Invalid inflected_count type (ref: %d, word_num: %d, type: `%s`, value: %v)", int(wd.Ref), int(wd.WordNumber), reflect.TypeOf(wd.Fields[2].Data), wd.Fields[2].Data))
	d.assert(inflected_count >= 0, fmt.Sprintf("Invalid inflected_count value: below zero (ref: %d, word_num: %d, value: %d)", int(wd.Ref), int(wd.WordNumber), inflected_count))

	dictionary, ok := wd.Fields[3].Data.([][]ScriptureWordDataField)
	d.assert(ok, fmt.Sprintf("Invalid dictionary type (ref: %d, word_num: %d, type: `%s`, value: %v)", int(wd.Ref), int(wd.WordNumber), reflect.TypeOf(wd.Fields[3].Data), wd.Fields[3].Data))
	d.assert(len(dictionary) > 0, fmt.Sprintf("Invalid dictionary value: empty (ref: %d, word_num: %d)", int(wd.Ref), int(wd.WordNumber)))
	for i, entry := range dictionary {
		d.assert(len(entry) == 5, fmt.Sprintf("Invalid dictionary entry length (ref: %d, word_num: %d, index: %d, value: %v)", int(wd.Ref), int(wd.WordNumber), i, dictionary))

		form, ok := entry[0].Data.(string)
		d.assert(ok, fmt.Sprintf("Invalid dictionary entry form type (ref: %d, word_num: %d, index: %d, type: `%s`, value: %v)", int(wd.Ref), int(wd.WordNumber), i, reflect.TypeOf(entry[0].Data), entry[0].Data))
		d.assert(len(form) > 0, fmt.Sprintf("Invalid dictionary entry form value: empty (ref: %d, word_num: %d, index: %d)", int(wd.Ref), int(wd.WordNumber), i))

		gloss, ok := entry[1].Data.(string)
		d.assert(ok, fmt.Sprintf("Invalid dictionary entry gloss type (ref: %d, word_num: %d, index: %d, type: `%s`, value: %v)", int(wd.Ref), int(wd.WordNumber), i, reflect.TypeOf(entry[1].Data), entry[1].Data))
		d.assert(len(gloss) > 0, fmt.Sprintf("Invalid dictionary entry gloss value: empty (ref: %d, word_num: %d, index: %d)", int(wd.Ref), int(wd.WordNumber), i))

		strong, ok := entry[2].Data.(string)
		d.assert(ok, fmt.Sprintf("Invalid dictionary entry strong type (ref: %d, word_num: %d, index: %d, type: `%s`, value: %v)", int(wd.Ref), int(wd.WordNumber), i, reflect.TypeOf(entry[2].Data), entry[2].Data))
		d.assert(len(strong) > 0, fmt.Sprintf("Invalid dictionary entry strong value: empty (ref: %d, word_num: %d, index: %d)", int(wd.Ref), int(wd.WordNumber), i))

		grammar, ok := entry[3].Data.(string)
		d.assert(ok, fmt.Sprintf("Invalid dictionary entry grammar type (ref: %d, word_num: %d, index: %d, type: `%s`, value: %v)", int(wd.Ref), int(wd.WordNumber), i, reflect.TypeOf(entry[3].Data), entry[3].Data))
		d.assert(len(grammar) > 0, fmt.Sprintf("Invalid dictionary entry grammar value: empty (ref: %d, word_num: %d, index: %d)", int(wd.Ref), int(wd.WordNumber), i))

		form_count, ok := entry[4].Data.(int)
		d.assert(ok, fmt.Sprintf("Invalid dictionary entry form_count type (ref: %d, word_num: %d, index: %d, type: `%s`, value: %v)", int(wd.Ref), int(wd.WordNumber), i, reflect.TypeOf(entry[4].Data), entry[4].Data))
		d.assert(form_count >= 0, fmt.Sprintf("Invalid dictionary entry form_count value: below zero (ref: %d, word_num: %d, index: %d, value: %d)", int(wd.Ref), int(wd.WordNumber), i, form_count))
	}
}

func (d *DataDB) getScriptureWordGnt_Text(wg *sync.WaitGroup, wd *ScriptureWordData) {
	db := <-d.pool
	stmt := db.queries.gntWordText
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

func (d *DataDB) getScriptureWordGnt_BasicInfo(wg *sync.WaitGroup, wd *ScriptureWordData) {
	db := <-d.pool
	stmt := db.queries.gntWordBasicInfo
	err := stmt.Bind(int(wd.Ref), int(wd.WordNumber))
	d.assert(err == nil, fmt.Sprintf("Error binding ref and word_num to sql query (ref: %d, word_num: %d) (err: %v)", int(wd.Ref), int(wd.WordNumber), err))

	has_row, err := stmt.Step()
	d.assert(err == nil, fmt.Sprintf("Error stepping sql query (ref: %d, word_num: %d) (err: %v)", int(wd.Ref), int(wd.WordNumber), err))
	d.assert(has_row, fmt.Sprintf("Error finding db row for word data (ref: %d, word_num: %d)", int(wd.Ref), int(wd.WordNumber)))

	var translit, english string
	err = stmt.Scan(&translit, &english)
	d.assert(err == nil, fmt.Sprintf("Error scanning sql query (ref: %d, word_num: %d) (err: %v)", int(wd.Ref), int(wd.WordNumber), err))
	wd.Fields[0].Data = translit
	wd.Fields[1].Data = english

	err = stmt.Reset()
	d.assert(err == nil, fmt.Sprintf("Error reseting sql statement (err: %v)", err))
	d.pool <- db
	wg.Done()
}

func (d *DataDB) getScriptureWordGnt_InflectedCount(wg *sync.WaitGroup, wd *ScriptureWordData) {
	db := <-d.pool
	stmt := db.queries.gntWordInflectedCount
	err := stmt.Bind(int(wd.Ref), int(wd.WordNumber))
	d.assert(err == nil, fmt.Sprintf("Error binding ref and word_num to sql query (ref: %d, word_num: %d) (err: %v)", int(wd.Ref), int(wd.WordNumber), err))

	has_row, err := stmt.Step()
	d.assert(err == nil, fmt.Sprintf("Error stepping sql query (ref: %d, word_num: %d) (err: %v)", int(wd.Ref), int(wd.WordNumber), err))
	d.assert(has_row, fmt.Sprintf("Error finding db row for word data (ref: %d, word_num: %d)", int(wd.Ref), int(wd.WordNumber)))

	var count int
	err = stmt.Scan(&count)
	d.assert(err == nil, fmt.Sprintf("Error scanning sql query (ref: %d, word_num: %d) (err: %v)", int(wd.Ref), int(wd.WordNumber), err))
	wd.Fields[2].Data = count

	err = stmt.Reset()
	d.assert(err == nil, fmt.Sprintf("Error reseting sql statement (err: %v)", err))
	d.pool <- db
	wg.Done()
}

const MAX_GNT_DICTIONARY_ROWS = 100

func (d *DataDB) getScriptureWordGnt_Dictionary(wg *sync.WaitGroup, wd *ScriptureWordData) {
	db := <-d.pool
	stmt := db.queries.gntWordDictionary
	err := stmt.Bind(int(wd.Ref), int(wd.WordNumber))
	d.assert(err == nil, fmt.Sprintf("Error binding ref and word_num to sql query (ref: %d, word_num: %d) (err: %v)", int(wd.Ref), int(wd.WordNumber), err))

	dictionary := make([][]ScriptureWordDataField, 0)
	var form, gloss, strong, grammar string
	var count int

	i := 0
	for i < MAX_GNT_DICTIONARY_ROWS {
		has_row, err := stmt.Step()
		d.assert(err == nil, fmt.Sprintf("Error stepping sql query (ref: %d, word_num: %d, iteration: %d) (err: %v)", int(wd.Ref), int(wd.WordNumber), i, err))
		if !has_row {
			d.assert(i > 0, fmt.Sprintf("Error finding db row (ref: %d, word_num: %d)", int(wd.Ref), int(wd.WordNumber)))
			break
		}

		err = stmt.Scan(&form, &gloss, &strong, &grammar, &count)
		d.assert(err == nil, fmt.Sprintf("Error scanning sql query (ref: %d, word_num: %d) (err: %v)", int(wd.Ref), int(wd.WordNumber), err))

		dictionary_entry := make([]ScriptureWordDataField, 5)
		dictionary_entry[0].Name = "Form"
		dictionary_entry[0].Data = form
		dictionary_entry[1].Name = "Gloss"
		dictionary_entry[1].Data = gloss
		dictionary_entry[2].Name = "Strong"
		dictionary_entry[2].Data = strong
		dictionary_entry[3].Name = "Grammar"
		dictionary_entry[3].Data = grammar
		dictionary_entry[4].Name = "FormCount"
		dictionary_entry[4].Data = count
		dictionary = append(dictionary, dictionary_entry)

		i++
		d.assert(i != MAX_GNT_DICTIONARY_ROWS, fmt.Sprintf("Looped too many times (>= %d) for returned dictionary entries (ref: %d, word_num: %d)", MAX_GNT_DICTIONARY_ROWS, int(wd.Ref), int(wd.WordNumber)))
	}

	wd.Fields[3].Data = dictionary

	err = stmt.Reset()
	d.assert(err == nil, fmt.Sprintf("Error reseting sql statement (err: %v)", err))
	d.pool <- db
	wg.Done()
}
