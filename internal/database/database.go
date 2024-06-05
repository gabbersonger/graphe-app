package database

import (
	"context"
	"reflect"
	rt "runtime"

	"github.com/bvinc/go-sqlite-lite/sqlite3"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type GrapheDB struct {
	ctx       context.Context
	pool_size int
	Pool      chan *GrapheDBConn
}

type GrapheDBConn struct {
	conn    *sqlite3.Conn
	queries GrapheQueries
}

func NewDB(ctx context.Context, dbFile string) *GrapheDB {
	g := &GrapheDB{}
	g.ctx = ctx
	g.pool_size = rt.NumCPU()
	g.Pool = make(chan *GrapheDBConn, g.pool_size)
	for i := 0; i < g.pool_size; i++ {
		g.Pool <- newGrapheDBConn(g, dbFile)
	}
	return g
}

func (g *GrapheDB) check(e error) {
	if e != nil {
		runtime.LogFatal(g.ctx, e.Error())
	}
}

func (g *GrapheDB) throw(s string) {
	runtime.LogFatal(g.ctx, s)
}

func (g *GrapheDB) ClosePool() {
	for i := 0; i < g.pool_size; i++ {
		db := <-g.Pool
		queries := reflect.ValueOf(db.queries)
		for j := 0; j < queries.NumField(); j++ {
			if queries.Field(j).Type() == reflect.TypeOf((*sqlite3.Stmt)(nil)) {
				queries.Field(j).MethodByName("Close")
			}
		}
		db.conn.Close()
	}
}

func newGrapheDBConn(g *GrapheDB, dbFile string) *GrapheDBConn {
	db := &GrapheDBConn{}
	conn, err := sqlite3.Open("file:" + dbFile)
	g.check(err)
	db.conn = conn
	prepareQueries(g, db)
	return db
}
