package settings

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/bvinc/go-sqlite-lite/sqlite3"
)

type SettingsDB struct {
	logger *slog.Logger
	db     *sqlite3.Conn
}

func (s *SettingsDB) assert(cond bool, msg string) {
	if !cond {
		if s.logger != nil {
			s.logger.Error(fmt.Sprintf("[SettingsDB] %s", msg))
		} else {
			panic(msg)
		}
	}
}

func (s *SettingsDB) OnShutdown() error {
	err := s.db.Close()
	s.assert(err == nil, "Error closing connection")
	s.logger.Info("Closed connection sucessfully")
	return nil
}

func NewSettingsDB(logger *slog.Logger) *SettingsDB {
	s := &SettingsDB{
		logger: logger,
	}

	home_directory, err := os.UserHomeDir()
	s.assert(err == nil, "Error getting user home directory")
	data_directory := filepath.Join(home_directory, "/Library/Application Support/Graphe")
	file_name := filepath.Join(data_directory, "/settings.db")
	s.db, err = sqlite3.Open("file:" + file_name)
	s.assert(err == nil, fmt.Sprintf("Error connecting to db (file_name: `%s`)", file_name))
	s.logger.Info(fmt.Sprintf("Connection created successfully (file_name: `%s`)", file_name))

	s.assert(s.validateDB(), "Error validating db")
	s.logger.Info("Validated settings in database")
	return s
}
