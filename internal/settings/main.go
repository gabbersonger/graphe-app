package settings

import (
	"context"

	"github.com/bvinc/go-sqlite-lite/sqlite3"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Settings struct {
	ctx context.Context
	db  *sqlite3.Conn
}

func (s *Settings) check(e error) {
	if e != nil {
		runtime.LogFatal(s.ctx, e.Error())
	}
}

func (s *Settings) throw(st string) {
	runtime.LogFatal(s.ctx, st)
}

func Startup(ctx context.Context, dbFile string) *Settings {
	s := &Settings{}
	s.ctx = ctx
	var err error
	s.db, err = sqlite3.Open("file:" + dbFile)
	s.check(err)
	validateDB(s)
	return s
}

func (s *Settings) Shutdown() {
	s.db.Close()
}
