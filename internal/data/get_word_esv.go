package data

import (
	"fmt"
	"reflect"
	"sync"
)

func (d *DataDB) getScriptureWordEsv(wd *ScriptureWordData) {
	prepareScriptureWordEsv(wd)

	wg := new(sync.WaitGroup)
	wg.Add(2)
	go d.getScriptureWordEsv_BasicInfo(wg, wd)
	go d.getScriptureWordEsv_StrongsInfo(wg, wd)
	wg.Wait()

	d.checkScriptureWordEsv(wd)
}

func prepareScriptureWordEsv(wd *ScriptureWordData) {
	wd.Fields = make([]ScriptureWordDataField, 2)
	wd.Fields[0].Name = "EnglishCount"
	wd.Fields[1].Name = "Dictionary"
}

func (d *DataDB) checkScriptureWordEsv(wd *ScriptureWordData) {
	d.assert(len(wd.Text) > 0, fmt.Sprintf("Invalid text scanned (ref: %d, word_num: %d)", int(wd.Ref), int(wd.WordNumber)))
	d.assert(len(wd.Fields) == 2, fmt.Sprintf("Invalid fields length (ref: %d, word_num: %d, length: %d)", int(wd.Ref), int(wd.WordNumber), len(wd.Fields)))

	english_count, ok := wd.Fields[0].Data.(int)
	d.assert(ok, fmt.Sprintf("Invalid english_count type (ref: %d, word_num: %d, type: `%s`, value: %v)", int(wd.Ref), int(wd.WordNumber), reflect.TypeOf(wd.Fields[0].Data), wd.Fields[0].Data))
	d.assert(english_count >= 0, fmt.Sprintf("Invalid english_count value: below zero (ref: %d, word_num: %d, value: %d)", int(wd.Ref), int(wd.WordNumber), english_count))

	dictionary, ok := wd.Fields[1].Data.([][]ScriptureWordDataField)
	d.assert(ok, fmt.Sprintf("Invalid dictionary type (ref: %d, word_num: %d, type: `%s`, value: %v)", int(wd.Ref), int(wd.WordNumber), reflect.TypeOf(wd.Fields[1].Data), wd.Fields[1].Data))
	d.assert(len(dictionary) > 0, fmt.Sprintf("Invalid dictionary value: empty (ref: %d, word_num: %d)", int(wd.Ref), int(wd.WordNumber)))
	for i, entry := range dictionary {
		d.assert(len(entry) == 2, fmt.Sprintf("Invalid dictionary entry length (ref: %d, word_num: %d, index: %d, value: %v)", int(wd.Ref), int(wd.WordNumber), i, dictionary))

		strong, ok := entry[0].Data.(string)
		d.assert(ok, fmt.Sprintf("Invalid dictionary entry strong type (ref: %d, word_num: %d, index: %d, type: `%s`, value: %v)", int(wd.Ref), int(wd.WordNumber), i, reflect.TypeOf(entry[0].Data), entry[0].Data))
		d.assert(len(strong) > 0, fmt.Sprintf("Invalid dictionary entry strong value: empty (ref: %d, word_num: %d, index: %d)", int(wd.Ref), int(wd.WordNumber), i))

		strong_count, ok := entry[1].Data.(int)
		d.assert(ok, fmt.Sprintf("Invalid dictionary entry strong_count type (ref: %d, word_num: %d, index: %d, type: `%s`, value: %v)", int(wd.Ref), int(wd.WordNumber), i, reflect.TypeOf(entry[1].Data), entry[1].Data))
		d.assert(strong_count >= 0, fmt.Sprintf("Invalid dictionary entry strong_count value: below zero (ref: %d, word_num: %d, index: %d, value: %d)", int(wd.Ref), int(wd.WordNumber), i, strong_count))
	}
}

func (d *DataDB) getScriptureWordEsv_BasicInfo(wg *sync.WaitGroup, wd *ScriptureWordData) {
	db := <-d.pool
	stmt := db.queries.esvWordBasicInfo
	err := stmt.Bind(int(wd.Ref), int(wd.WordNumber))
	d.assert(err == nil, fmt.Sprintf("Error binding ref and word_num to sql query (ref: %d, word_num: %d) (err: %v)", int(wd.Ref), int(wd.WordNumber), err))

	has_row, err := stmt.Step()
	d.assert(err == nil, fmt.Sprintf("Error stepping sql query (ref: %d, word_num: %d) (err: %v)", int(wd.Ref), int(wd.WordNumber), err))
	d.assert(has_row, fmt.Sprintf("Error finding db row for word data (ref: %d, word_num: %d)", int(wd.Ref), int(wd.WordNumber)))

	var count int
	err = stmt.Scan(&(wd.Text), &count)
	d.assert(err == nil, fmt.Sprintf("Error scanning sql query (ref: %d, word_num: %d) (err: %v)", int(wd.Ref), int(wd.WordNumber), err))
	wd.Fields[0].Data = count

	err = stmt.Reset()
	d.assert(err == nil, fmt.Sprintf("Error reseting sql statement (err: %v)", err))
	d.pool <- db
	wg.Done()
}

const MAX_ESV_DICTIONARY_ROWS = 100

func (d *DataDB) getScriptureWordEsv_StrongsInfo(wg *sync.WaitGroup, wd *ScriptureWordData) {
	db := <-d.pool
	stmt := db.queries.esvWordStrongsInfo
	err := stmt.Bind(int(wd.Ref), int(wd.WordNumber))
	d.assert(err == nil, fmt.Sprintf("Error binding ref and word_num to sql query (ref: %d, word_num: %d) (err: %v)", int(wd.Ref), int(wd.WordNumber), err))

	dictionary := make([][]ScriptureWordDataField, 0)
	var strong string
	var count int

	i := 0
	for i < MAX_ESV_DICTIONARY_ROWS {
		has_row, err := stmt.Step()
		d.assert(err == nil, fmt.Sprintf("Error stepping sql query (ref: %d, word_num: %d, iteration: %d) (err: %v)", int(wd.Ref), int(wd.WordNumber), i, err))
		if !has_row {
			d.assert(i > 0, fmt.Sprintf("Error finding db row (ref: %d, word_num: %d)", int(wd.Ref), int(wd.WordNumber)))
			break
		}

		err = stmt.Scan(&strong, &count)
		d.assert(err == nil, fmt.Sprintf("Error scanning sql query (ref: %d, word_num: %d) (err: %v)", int(wd.Ref), int(wd.WordNumber), err))

		dictionary_entry := make([]ScriptureWordDataField, 2)
		dictionary_entry[0].Name = "Strong"
		dictionary_entry[0].Data = strong
		dictionary_entry[1].Name = "StrongCount"
		dictionary_entry[1].Data = count
		// TODO: add strongs info: form, gloss, grammar
		dictionary = append(dictionary, dictionary_entry)

		i++
		d.assert(i != MAX_ESV_DICTIONARY_ROWS, fmt.Sprintf("Looped too many times (>= %d) for returned dictionary entries (ref: %d, word_num: %d)", MAX_GNT_DICTIONARY_ROWS, int(wd.Ref), int(wd.WordNumber)))
	}

	wd.Fields[1].Data = dictionary

	err = stmt.Reset()
	d.assert(err == nil, fmt.Sprintf("Error reseting sql statement (err: %v)", err))
	d.pool <- db
	wg.Done()
}
