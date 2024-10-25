package settings

import (
	"fmt"
	"graphe/internal/logger"
	"graphe/internal/menu"
	"os"
	"path/filepath"

	"github.com/bvinc/go-sqlite-lite/sqlite3"
)

type SettingsDB struct {
	logger       *logger.Logger
	menu_manager *menu.MenuManager
	db           *sqlite3.Conn
}

func (s *SettingsDB) assert(cond bool, msg string) {
	s.logger.Assert("SettingsDB", cond, msg)
}

func (s *SettingsDB) log(msg string) {
	s.logger.Log("SettingsDB", msg)
}

func (s *SettingsDB) Name() string {
	return "SettingsDB"
}

func (s *SettingsDB) OnShutdown() error {
	err := s.db.Close()
	s.assert(err == nil, "Error closing connection")
	s.log("Closed connection sucessfully")
	return nil
}

func NewSettingsDB(logger *logger.Logger, menu_manager *menu.MenuManager) *SettingsDB {
	s := &SettingsDB{
		logger:       logger,
		menu_manager: menu_manager,
	}

	home_directory, err := os.UserHomeDir()
	s.assert(err == nil, "Error getting user home directory")
	data_directory := filepath.Join(home_directory, "/Library/Application Support/Graphe")
	file_name := filepath.Join(data_directory, "/settings.db")
	s.db, err = sqlite3.Open("file:" + file_name)
	s.assert(err == nil, fmt.Sprintf("Error connecting to db (file_name: `%s`)", file_name))
	s.log(fmt.Sprintf("Connection created successfully (file_name: `%s`)", file_name))

	s.assert(s.validateDB(), "Error validating db")
	s.log("Validated settings in database")

	shortcuts := s.GetSettings().Shortcuts
	s.menu_manager.SetShortcuts(&shortcuts)

	return s
}
