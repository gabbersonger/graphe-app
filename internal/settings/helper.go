package settings

import "reflect"

type SettingTable struct {
	name    string
	columns []SettingColumn
}

type SettingColumn struct {
	name     string
	sql_type string
}

func getSettingsTables() []SettingTable {
	r := reflect.TypeFor[SettingsValues]()
	tables := make([]SettingTable, 0, r.NumField())
	for _, field := range reflect.VisibleFields(r) {
		tables = append(tables, SettingTable{
			name:    field.Name,
			columns: getSettingsColumns(field.Type, ""),
		})
	}
	return tables
}

func getSettingsColumns(t reflect.Type, prefix string) []SettingColumn {
	columns := make([]SettingColumn, 0, t.NumField())
	for _, field := range reflect.VisibleFields(t) {
		n := field.Name
		if len(prefix) > 0 {
			n = prefix + "_" + n
		}

		column := SettingColumn{name: n}
		switch field.Type.Kind() {
		case reflect.String:
			column.sql_type = "varchar"
		case reflect.Int:
			column.sql_type = "integer"
		case reflect.Bool:
			column.sql_type = "integer"
		case reflect.Struct:
			columns = append(columns, getSettingsColumns(field.Type, n)...)
			continue
		}
		columns = append(columns, column)
	}
	return columns
}
