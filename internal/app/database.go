package app

import (
	"errors"
	rt "runtime"

	"github.com/bvinc/go-sqlite-lite/sqlite3"
)

type GrapheDB struct {
	pool_size int
	pool      chan *GrapheDBConn
}

type GrapheDBConn struct {
	conn    *sqlite3.Conn
	queries map[string]*sqlite3.Stmt
}

func (db *GrapheDBConn) getQuery(key string) (*sqlite3.Stmt, error) {
	stmt, ok := db.queries[key]
	if !ok {
		return nil, errors.New("DB does not have query: '" + key + "'")
	}
	return stmt, nil
}

func prepareQuery(db *GrapheDBConn, key string, sql string) error {
	stmt, err := db.conn.Prepare(sql)
	if err == nil {
		db.queries[key] = stmt
	}
	return err
}

func prepareQueries(a *App, db *GrapheDBConn) {
	db.queries = make(map[string]*sqlite3.Stmt)

	var err error

	// GNT
	err = prepareQuery(db, "GetGNTScriptureSection", `
        SELECT ref, word_num, text, pre, post
        FROM gnt_text
        WHERE ref >= ? AND ref <= ?;
    `)
	a.check(err)

	err = prepareQuery(db, "GetGNTScriptureWordText", `
        SELECT text
        FROM gnt_text
        WHERE
            ref = ?
            AND word_num = ?
        LIMIT 1;
    `)
	a.check(err)

	err = prepareQuery(db, "GetGNTScriptureWordBasicInfo", `
        SELECT translit, english
        FROM gnt_text_info
        WHERE
            ref = ?
            AND word_num = ?
        LIMIT 1;
    `)
	a.check(err)

	err = prepareQuery(db, "GetGNTScriptureWordDictionaryInfo", `
        SELECT form, gloss
        FROM gnt_text_dictionary
        WHERE
            ref = ?
            AND word_num = ?;
    `)
	a.check(err)

	err = prepareQuery(db, "GetGNTScriptureWordStrongsInfo", `
        SELECT strong, grammar
        FROM gnt_text_strongs
        WHERE
            ref = ?
            AND word_num = ?;
    `)
	a.check(err)

	err = prepareQuery(db, "GetGNTScriptureWordInflectedCount", `
        SELECT count(*)
        FROM gnt_text
        WHERE text = (
        	SELECT text
         	FROM gnt_text
          	WHERE
           		ref = ?
             	AND word_num = ?
            LIMIT 1
        )
        LIMIT 1;
    `)
	a.check(err)

	err = prepareQuery(db, "GetGNTScriptureWordLexemeCount", `
        SELECT count(*)
        FROM gnt_text_dictionary
        WHERE form = (
        	SELECT form
         	FROM gnt_text_dictionary
          	WHERE
           		ref = ?
		        AND word_num = ?
			LIMIT 1
        )
        LIMIT 1;
    `)
	a.check(err)

	// LXX
	err = prepareQuery(db, "GetLXXScriptureSection", `
        SELECT ref, word_num, text, pre, post
        FROM lxx_text
        WHERE ref >= ? AND ref <= ?;
    `)
	a.check(err)

	err = prepareQuery(db, "GetLXXScriptureWordText", `
        SELECT text
        FROM lxx_text
        WHERE
            ref = ?
            AND word_num = ?
        LIMIT 1;
    `)
	a.check(err)

	err = prepareQuery(db, "GetLXXScriptureWordBasicInfo", `
        SELECT
        	translit, english, strongs, grammar,
         	dictionary_form, dictionary_gloss
        FROM lxx_text_info
        WHERE
            ref = ?
            AND word_num = ?
        LIMIT 1;
    `)
	a.check(err)

	err = prepareQuery(db, "GetLXXScriptureWordInflectedCount", `
        SELECT count(*)
        FROM lxx_text
        WHERE text = (
        	SELECT text
         	FROM lxx_text
          	WHERE
           		ref = ?
             	AND word_num = ?
            LIMIT 1
        )
        LIMIT 1;
    `)
	a.check(err)

	err = prepareQuery(db, "GetLXXScriptureWordLexemeCount", `
        SELECT count(*)
        FROM lxx_text_info
        WHERE dictionary_form = (
        	SELECT dictionary_form
         	FROM lxx_text_info
          	WHERE
           		ref = ?
	            AND word_num = ?
			LIMIT 1
        )
        LIMIT 1;
    `)
	a.check(err)
}

func newGrapheDB(a *App, dbFile string) *GrapheDBConn {
	db := &GrapheDBConn{}
	conn, err := sqlite3.Open("file:" + dbFile)
	a.check(err)
	db.conn = conn
	prepareQueries(a, db)
	return db
}

func (a *App) setupDatabasePool() {
	dbFile := a.Env.DataDirectory + "/graphe.db"

	num_db_conn := rt.NumCPU()
	a.db = GrapheDB{
		pool_size: num_db_conn,
		pool:      make(chan *GrapheDBConn, num_db_conn),
	}

	for i := 0; i < a.db.pool_size; i++ {
		a.db.pool <- newGrapheDB(a, dbFile)
	}
}

func (a *App) closeDatabasePool() {
	for i := 0; i < a.db.pool_size; i++ {
		d := <-a.db.pool
		for _, stmt := range d.queries {
			stmt.Close()
		}
		d.conn.Close()
	}
}
