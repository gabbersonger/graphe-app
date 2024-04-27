package app

import (
	"reflect"
	rt "runtime"

	"github.com/bvinc/go-sqlite-lite/sqlite3"
)

type GrapheDB struct {
	pool_size int
	pool      chan *GrapheDBConn
}

type GrapheDBConn struct {
	conn    *sqlite3.Conn
	queries GrapheQueries
}

type GrapheQueries struct {
	GntSection            *sqlite3.Stmt
	GntWordText           *sqlite3.Stmt
	GntWordBasicInfo      *sqlite3.Stmt
	GntWordDictionary     *sqlite3.Stmt
	GntWordStrongs        *sqlite3.Stmt
	GntWordInflectedCount *sqlite3.Stmt
	GntWordLexemeCount    *sqlite3.Stmt

	LxxSection            *sqlite3.Stmt
	LxxWordText           *sqlite3.Stmt
	LxxWordBasicInfo      *sqlite3.Stmt
	LxxWordInflectedCount *sqlite3.Stmt
	LxxWordLexemeCount    *sqlite3.Stmt
}

func (a *App) setupDatabasePool() {
	dbFile := a.Env.DataDirectory + "/graphe.db"
	numDBConn := rt.NumCPU()
	a.db = GrapheDB{
		pool_size: numDBConn,
		pool:      make(chan *GrapheDBConn, numDBConn),
	}
	for i := 0; i < numDBConn; i++ {
		a.db.pool <- newGrapheDB(a, dbFile)
	}
}

func (a *App) closeDatabasePool() {
	for i := 0; i < a.db.pool_size; i++ {
		d := <-a.db.pool
		queries := reflect.ValueOf(d.queries)
		for j := 0; j < queries.NumField(); j++ {
			if queries.Field(j).Type() == reflect.TypeOf((*sqlite3.Stmt)(nil)) {
				queries.Field(j).Close()
			}
		}
		d.conn.Close()
	}
}

func newGrapheDB(a *App, dbFile string) *GrapheDBConn {
	db := &GrapheDBConn{}
	conn, err := sqlite3.Open("file:" + dbFile)
	a.check(err)
	db.conn = conn
	prepareQueries(a, db)
	return db
}

func prepareQueries(a *App, db *GrapheDBConn) {
	var err error
	db.queries = GrapheQueries{}

	db.queries.GntSection, err = db.conn.Prepare(`
        SELECT ref, word_num, text, pre, post
        FROM gnt_text
        WHERE ref >= ? AND ref <= ?;
    `)
	a.check(err)

	db.queries.GntWordText, err = db.conn.Prepare(`
        SELECT text
        FROM gnt_text
        WHERE
            ref = ?
            AND word_num = ?
        LIMIT 1;
    `)
	a.check(err)

	db.queries.GntWordBasicInfo, err = db.conn.Prepare(`
        SELECT translit, english
        FROM gnt_text_info
        WHERE
            ref = ?
            AND word_num = ?
        LIMIT 1;
    `)
	a.check(err)

	db.queries.GntWordDictionary, err = db.conn.Prepare(`
        SELECT form, gloss
        FROM gnt_text_dictionary
        WHERE
            ref = ?
            AND word_num = ?;
    `)
	a.check(err)

	db.queries.GntWordStrongs, err = db.conn.Prepare(`
        SELECT strong, grammar
        FROM gnt_text_strongs
        WHERE
            ref = ?
            AND word_num = ?;
    `)
	a.check(err)

	db.queries.GntWordInflectedCount, err = db.conn.Prepare(`
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

	db.queries.GntWordLexemeCount, err = db.conn.Prepare(`
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
	db.queries.LxxSection, err = db.conn.Prepare(`
        SELECT ref, word_num, text, pre, post
        FROM lxx_text
        WHERE ref >= ? AND ref <= ?;
    `)
	a.check(err)

	db.queries.LxxWordText, err = db.conn.Prepare(`
        SELECT text
        FROM lxx_text
        WHERE
            ref = ?
            AND word_num = ?
        LIMIT 1;
    `)
	a.check(err)

	db.queries.LxxWordBasicInfo, err = db.conn.Prepare(`
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

	db.queries.LxxWordInflectedCount, err = db.conn.Prepare(`
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

	db.queries.LxxWordLexemeCount, err = db.conn.Prepare(`
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
