package settings

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/bvinc/go-sqlite-lite/sqlite3"
)

func (s *Settings) setupDB(dbFile string) {
	var err error
	s.db, err = sqlite3.Open("file:" + dbFile)
	s.check(err)
	ensureSettingsValuesExist(s)
}

func ensureSettingsValuesExist(s *Settings) {
	fields := reflect.VisibleFields(reflect.TypeFor[SettingsValues]())
	for _, field := range fields {
		table := strings.ToLower(field.Name)
		ensureTableExists(s, table)

		columns := createSettingColumns(field.Type, "")
		for _, column := range columns {
			ensureColumnExists(s, table, column)
		}
	}
}

type SettingColumn struct {
	Name string
	Type string
}

func createSettingColumns(t reflect.Type, prefix string) []SettingColumn {
	arr := make([]SettingColumn, 0)

	if prefix != "" {
		prefix += "_"
	}

	fields := reflect.VisibleFields(t)
	for _, field := range fields {
		n := strings.ToLower(field.Name)
		t := ""
		switch field.Type.Kind() {
		case reflect.String:
			t = "varchar"
		case reflect.Int:
			t = "integer"
		case reflect.Bool:
			t = "integer"
		case reflect.Struct:
			arr = append(arr, createSettingColumns(field.Type, n)...)
			continue
		}

		arr = append(arr, SettingColumn{
			Name: prefix + n,
			Type: t,
		})
	}

	return arr
}

func ensureTableExists(s *Settings, table string) {
	stmt, err := s.db.Prepare(fmt.Sprintf(`
		SELECT 1
		FROM sqlite_master
		WHERE
			type = 'table'
			AND name = '%s';
	`, table))
	s.check(err)

	hasRow, err := stmt.Step()
	s.check(err)
	if !hasRow {
		err := s.db.Exec(fmt.Sprintf(`
			CREATE TABLE %s (
				id integer PRIMARY KEY
			);
		`, table))
		s.check(err)
	}
	stmt.Close()
}

func ensureColumnExists(s *Settings, table string, column SettingColumn) {
	stmt, err := s.db.Prepare(fmt.Sprintf(`
		SELECT 1
		FROM PRAGMA_TABLE_INFO('%s')
		WHERE
			name = '%s'
			AND lower(type) = '%s';
	`, table, column.Name, column.Type))
	s.check(err)

	hasRow, err := stmt.Step()
	s.check(err)
	if !hasRow {
		err := s.db.Exec(fmt.Sprintf(`
			ALTER TABLE %s
				DROP COLUMN %s;
		`, table, column.Name))
		s.check(err)

		err = s.db.Exec(fmt.Sprintf(`
			ALTER TABLE %s
				ADD %s %s;
		`, table, column.Name, column.Type))
		s.check(err)
	}
	stmt.Close()
}
