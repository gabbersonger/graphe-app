package settings

import (
	"fmt"
	"strings"
)

func (s *SettingsDB) validateDB() bool {
	for _, t := range getSettingsTables() {
		ensureTableExists(s, t.name)
		for _, c := range t.columns {
			ensureColumnExists(s, t.name, c)
		}
		ensureRowExists(s, t.name)
	}
	return true
}

func getRowInDB(s *SettingsDB, q string, args ...interface{}) bool {
	query := strings.TrimSpace(fmt.Sprintf(q, args...))
	err := s.db.Begin()
	s.assert(err == nil, "Error beginning transaction")
	stmt, err := s.db.Prepare(query)
	s.assert(err == nil, fmt.Sprintf("Error preparing query (query: `%s`)", query))
	has_row, err := stmt.Step()
	s.assert(err == nil, fmt.Sprintf("Error getting first row (query: `%s`)", query))
	if s.db.TotalChanges() > 0 {
		s.db.Rollback()
		s.assert(false, fmt.Sprintf("No rows should be changed (query: `%s`", query))
	}
	err = stmt.Close()
	s.assert(err == nil, "Error closing statement")
	err = s.db.Commit()
	s.assert(err == nil, "Error commiting")
	return has_row
}

func execInDB(s *SettingsDB, q string, args ...interface{}) {
	query := strings.TrimSpace(fmt.Sprintf(q, args...))
	err := s.db.Begin()
	s.assert(err == nil, "Error beginning transaction")
	err = s.db.Exec(query)
	s.assert(err == nil, fmt.Sprintf("Error executing query (query: `%s`)", query))
	s.db.Commit()
}

func checkTableExists(s *SettingsDB, t string) bool {
	return getRowInDB(s, `
		SELECT 1
		FROM sqlite_master
		WHERE type = 'table' AND name = '%s';
	`, t)
}

func ensureTableExists(s *SettingsDB, t string) {
	if !checkTableExists(s, t) {
		execInDB(s, `
			CREATE TABLE %s (
				id INTEGER PRIMARY KEY
			);
		`, t)
		s.assert(checkTableExists(s, t), fmt.Sprintf("Error creating table (table_name: `%s`", t))
	}
}

func checkColumnExists(s *SettingsDB, t string, c SettingColumn) bool {
	return getRowInDB(s, `
		SELECT 1
		FROM PRAGMA_TABLE_INFO('%s')
		WHERE name = '%s' AND lower(type) = '%s'
		LIMIT 1;
	`, t, c.name, c.sql_type)
}

func ensureColumnExists(s *SettingsDB, t string, c SettingColumn) {
	if !checkColumnExists(s, t, c) {
		execInDB(s, `
			ALTER TABLE %s ADD %s %s;
		`, t, c.name, c.sql_type)
		s.assert(checkColumnExists(s, t, c), fmt.Sprintf("Error creating column (table_name: `%s`, column_name: `%s`", t, c))
	}
}

func checkRowExists(s *SettingsDB, t string) bool {
	return getRowInDB(s, `
		SELECT 1
		FROM %s
		WHERE id = 1
		LIMIT 1;
	`, t)
}

func ensureRowExists(s *SettingsDB, t string) {
	if !checkRowExists(s, t) {
		execInDB(s, `
			INSERT INTO %s (id) VALUES (1);
		`, t)
		s.assert(checkRowExists(s, t), fmt.Sprintf("Error creating row (table_name: `%s`", t))
	}
}
