package app

import (
	"context"
	"errors"

	"github.com/bvinc/go-sqlite-lite/sqlite3"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type GrapheDB struct {
	conn    *sqlite3.Conn
	queries map[string]*sqlite3.Stmt
}

func (db *GrapheDB) getQuery(key string) (*sqlite3.Stmt, error) {
	stmt, ok := db.queries[key]
	if !ok {
		return nil, errors.New("DB does not have query: '" + key + "'")
	}
	return stmt, nil
}

func (db *GrapheDB) prepareQuery(key string, sql string) error {
	stmt, err := db.conn.Prepare(sql)
	if err == nil {
		db.queries[key] = stmt
	}
	return err
}

func (db *GrapheDB) prepareQueries(ctx context.Context) {
	db.queries = make(map[string]*sqlite3.Stmt)

	var err error

	err = db.prepareQuery("GetScriptureSection", `
        SELECT ref, word_num, text
        FROM gnt_text 
        WHERE ref >= ? AND ref <= ?;
    `)
	check(ctx, err)

	err = db.prepareQuery("GetScriptureWordBasicInfo", `
        SELECT translit, english, conjoin_word, sub_meaning
        FROM gnt_text_info
        WHERE 
            ref = ? 
            AND word_num = ? 
        LIMIT 1;
    `)
	check(ctx, err)

	err = db.prepareQuery("GetScriptureWordDictionaryInfo", `
        SELECT form, gloss
        FROM gnt_text_dictionary
        WHERE 
            ref = ? 
            AND word_num = ?;
    `)
	check(ctx, err)

	err = db.prepareQuery("GetScriptureWordStrongsInfo", `
        SELECT strong, grammar
        FROM gnt_text_strongs
        WHERE 
            ref = ? 
            AND word_num = ?;
    `)
	check(ctx, err)
}

func newGrapheDB(ctx context.Context, dbFile string) *GrapheDB {
	db := &GrapheDB{}
	conn, err := sqlite3.Open("file:" + dbFile)
	check(ctx, err)
	db.conn = conn
	db.prepareQueries(ctx)
	return db
}

func (a *App) setupDatabasePool() {
	dbFile := a.Env.DataDirectory + "/graphe.db"

	runtime.LogWarning(a.ctx, dbFile)

	a.db_pool = make(chan *GrapheDB, max_db_conn)
	for i := 0; i < max_db_conn; i++ {
		a.db_pool <- newGrapheDB(a.ctx, dbFile)
	}
}

func (a *App) closeDatabasePool() {
	for i := 0; i < len(a.db_pool); i++ {
		d := <-a.db_pool
		for _, stmt := range d.queries {
			stmt.Close()
		}
		d.conn.Close()
	}
}
