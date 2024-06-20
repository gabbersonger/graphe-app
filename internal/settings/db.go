package settings

import (
	"fmt"
	"reflect"
	"strings"
)

func validateDB(s *Settings) {
	tables := getSettingsTables()
	fmt.Println(tables)
	for _, t := range tables {
		ensureTableExists(s, t.name)
		for _, c := range t.columns {
			ensureColumnExists(s, t.name, c)
		}
		ensureRowExists(s, t.name)
	}
}

type SettingTable struct {
	name    string
	columns []SettingColumn
}

type SettingColumn struct {
	name     string
	sql_type string
}

func getSettingsTables() []SettingTable {
	t := reflect.TypeFor[SettingsValues]()
	tables := make([]SettingTable, 0, t.NumField())
	fields := reflect.VisibleFields(t)
	for _, field := range fields {
		tables = append(tables, SettingTable{
			name:    strings.ToLower(field.Name),
			columns: getSettingsColumns(field.Type, ""),
		})
	}
	return tables
}

func getSettingsColumns(t reflect.Type, prefix string) []SettingColumn {
	columns := make([]SettingColumn, 0, t.NumField())
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
			columns = append(columns, getSettingsColumns(field.Type, n)...)
			continue
		}
		if prefix != "" {
			n = prefix + "_" + n
		}
		columns = append(columns, SettingColumn{
			name:     n,
			sql_type: t,
		})
	}
	return columns
}

func getRowInDB(s *Settings, q string, args ...interface{}) bool {
	s.db.Begin()
	stmt, err := s.db.Prepare(fmt.Sprintf(q, args...))
	s.check(err)
	hasRow, err := stmt.Step()
	s.check(err)
	stmt.Close()
	s.db.Commit()
	return hasRow
}

func execInDB(s *Settings, q string, args ...interface{}) {
	s.db.Begin()
	fmt.Println(fmt.Sprintf(q, args...))
	err := s.db.Exec(fmt.Sprintf(q, args...))
	s.check(err)
	s.db.Commit()
}

func ensureTableExists(s *Settings, t string) {
	table_exists := getRowInDB(s, `
		SELECT 1
		FROM sqlite_master
		WHERE type = 'table' AND name = '%s';
	`, t)
	if !table_exists {
		execInDB(s, `
			CREATE TABLE %s (
				id INTEGER PRIMARY KEY
			);`, t)
	}
}

func ensureColumnExists(s *Settings, t string, c SettingColumn) {
	column_exists := getRowInDB(s, `
		SELECT 1
		FROM PRAGMA_TABLE_INFO('%s')
		WHERE name = '%s' AND lower(type) = '%s'
		LIMIT 1;
	`, t, c.name, c.sql_type)
	if !column_exists {
		execInDB(s, `ALTER TABLE %s DROP COLUMN %s;`, t, c.name)
		execInDB(s, `ALTER TABLE %s ADD %s %s;`, t, c.name, c.sql_type)
	}
}

func ensureRowExists(s *Settings, t string) {
	row_exists := getRowInDB(s, `
		SELECT 1
		FROM %s
		WHERE id = 1
		LIMIT 1;
	`, t)
	if !row_exists {
		execInDB(s, `INSERT INTO %s (id) VALUES (1);`, t)
	}
}
