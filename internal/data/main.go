package data

import (
	"context"
	"fmt"
	"reflect"
	rt "runtime"

	"github.com/bvinc/go-sqlite-lite/sqlite3"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type DataDB struct {
	ctx       context.Context
	pool_size int
	pool      chan *DataDBConn
}

type DataDBConn struct {
	conn    *sqlite3.Conn
	queries DataDBQueries
}

func (d *DataDB) assert(cond bool, msg string) {
	if !cond {
		if d.ctx != nil {
			runtime.LogFatal(d.ctx, fmt.Sprintf("[DataDB] %s", msg))
		} else {
			panic(msg)
		}
	}
}

func (d *DataDB) log(msg string) {
	runtime.LogInfo(d.ctx, fmt.Sprintf("[DataDB] %s", msg))
}

func CreateDB(ctx context.Context, dbFile string) *DataDB {
	d := &DataDB{}

	d.ctx = ctx
	d.assert(d.ctx != nil, "Invalid context")

	d.pool_size = rt.NumCPU()
	d.assert(d.pool_size > 0, fmt.Sprintf("Invalid db pool size (size: %d)", d.pool_size))

	d.pool = make(chan *DataDBConn, d.pool_size)
	for i := 0; i < d.pool_size; i++ {
		d.pool <- newDataDBConn(d, dbFile)
	}
	d.log(fmt.Sprintf("Created %d pooled connections successfully (file: `%s`)", d.pool_size, dbFile))
	return d
}

func newDataDBConn(d *DataDB, dbFile string) *DataDBConn {
	d.assert(len(dbFile) > 0, "Invalid dbFile (0 length)")
	db := &DataDBConn{}
	conn, err := sqlite3.Open("file:" + dbFile)
	d.assert(err == nil, fmt.Sprintf("Error connecting to data db (file name: `%s`)", dbFile))
	db.conn = conn
	d.prepareQueries(db)
	return db
}

func (d *DataDB) Ping() bool {
	db := <-d.pool
	stmt, err := db.conn.Prepare(`SELECT 1 FROM sqlite_master WHERE type = 'table' LIMIT 1;`)
	d.assert(err == nil, "Error preparing ping query ")
	has_row, err := stmt.Step()
	d.assert(err == nil, "Error getting ping query value")
	stmt.Close()
	d.pool <- db
	d.log("Connection pinged successfully")
	return has_row
}

func (d *DataDB) Shutdown() {
	for i := 0; i < d.pool_size; i++ {
		db := <-d.pool
		queries := reflect.ValueOf(db.queries)
		for j := 0; j < queries.NumField(); j++ {
			if queries.Field(j).Type() == reflect.TypeOf((*sqlite3.Stmt)(nil)) {
				queries.Field(j).MethodByName("Close")
			}
		}
		db.conn.Close()
	}
	d.log(fmt.Sprintf("Closed %d pooled connections", d.pool_size))
}
