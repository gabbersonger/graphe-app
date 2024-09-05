package settings

import (
	"context"
	"fmt"

	"github.com/bvinc/go-sqlite-lite/sqlite3"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type SettingsDB struct {
	ctx context.Context
	db  *sqlite3.Conn
}

func (s *SettingsDB) assert(cond bool, msg string) {
	if !cond {
		if s.ctx != nil {
			runtime.LogFatal(s.ctx, fmt.Sprintf("[SettingsDB] %s", msg))
		} else {
			panic(msg)
		}
	}
}

func (d *SettingsDB) log(msg string) {
	runtime.LogInfo(d.ctx, fmt.Sprintf("[SettingsDB] %s", msg))
}

func CreateDB(ctx context.Context, dbFile string) *SettingsDB {
	var err error
	s := &SettingsDB{}

	s.ctx = ctx
	s.assert(s.ctx != nil, "Invalid context")

	s.assert(len(dbFile) > 0, "Invalid dbFile (0 length)")
	s.db, err = sqlite3.Open("file:" + dbFile)
	s.assert(err == nil, fmt.Sprintf("Error connecting to settings db (file name: `%s`)", dbFile))
	s.log(fmt.Sprintf("Connection created successfully (file: `%s`)", dbFile))

	s.assert(s.validateDB(), "Error validating settings db")
	s.log("Validated settings in database")
	return s
}

func (s *SettingsDB) Ping() bool {
	stmt, err := s.db.Prepare(`SELECT 1 FROM sqlite_master WHERE type = 'table' LIMIT 1;`)
	s.assert(err == nil, "Error preparing ping query ")
	has_row, err := stmt.Step()
	s.assert(err == nil, "Error getting ping query value from database")
	stmt.Close()
	s.log("Connection pinged successfully")
	return has_row
}

func (s *SettingsDB) Shutdown() {
	s.db.Close()
	s.log("Closed connection sucessfully")
}
