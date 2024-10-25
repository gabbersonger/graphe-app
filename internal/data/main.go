package data

import (
	"fmt"
	"graphe/internal/logger"
	"graphe/internal/scripture"
	"os"
	"path/filepath"
	"reflect"
	rt "runtime"

	"github.com/bvinc/go-sqlite-lite/sqlite3"
)

type DataDB struct {
	logger            *logger.Logger
	scripture_service *scripture.ScriptureService
	pool_size         int
	pool              chan *DataDBConn
}

type DataDBConn struct {
	conn    *sqlite3.Conn
	queries DataDBQueries
}

func (d *DataDB) assert(cond bool, msg string) {
	d.logger.Assert("SettingsDB", cond, msg)
}

func (d *DataDB) log(msg string) {
	d.logger.Log("SettingsDB", msg)
}

func (d *DataDB) OnShutdown() error {
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
	return nil
}

func NewDataDB(logger *logger.Logger, scripture_service *scripture.ScriptureService) *DataDB {
	d := &DataDB{
		logger:            logger,
		scripture_service: scripture_service,
		pool_size:         rt.NumCPU(),
	}
	d.assert(d.pool_size > 0, fmt.Sprintf("Invalid db pool size (size: %d)", d.pool_size))

	home_directory, err := os.UserHomeDir()
	d.assert(err == nil, "Error getting user home directory")
	data_directory := filepath.Join(home_directory, "/Library/Application Support/Graphe")
	file_name := filepath.Join(data_directory, "/graphe.db")

	d.pool = make(chan *DataDBConn, d.pool_size)
	for i := 0; i < d.pool_size; i++ {
		d.pool <- newDataDBConn(d, file_name)
	}
	d.log(fmt.Sprintf("Created %d pooled connections (file: `%s`)", d.pool_size, file_name))
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
