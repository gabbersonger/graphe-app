package settings

import (
	"fmt"
	"reflect"
	"strings"
)

func getDefaultSettingsValues() SettingsValues {
	v := SettingsValues{}
	v.Appearence.Theme = "hanok"
	v.Appearence.Font.Greek = "default"
	v.Appearence.Font.System = "default"
	v.Appearence.Font.Hebrew = "default"
	v.Appearence.Font.English = "default"
	return v
}

func (s *Settings) GetSettings() SettingsValues {
	v := getDefaultSettingsValues()
	tables := getSettingsTables()

	r := reflect.ValueOf(&v).Elem()
	for _, table := range tables {
		item := r.FieldByName(table.name)
		for _, column := range table.columns {
			cur_item := item
			parts := strings.Split(column.name, "_")
			for _, part := range parts {
				cur_item = cur_item.FieldByName(part)
			}

			stmt, err := s.db.Prepare(fmt.Sprintf(`
				SELECT %s
				FROM %s
				WHERE id = 1
				LIMIT 1;
			`, column.name, table.name))
			s.check(err)
			hasRow, err := stmt.Step()
			s.check(err)
			if !hasRow {
				s.throw(fmt.Sprintf("Could not find row for setting (table=%s, column=%s)", table.name, column.name))
			}
			switch cur_item.Kind() {
			case reflect.Int:
				val, ok, err := stmt.ColumnInt64(0)
				s.check(err)
				if ok {
					cur_item.SetInt(val)
				}
			case reflect.String:
				val, ok, err := stmt.ColumnText(0)
				s.check(err)
				if ok {
					cur_item.SetString(val)
				}
			case reflect.Bool:
				val, ok, err := stmt.ColumnInt(0)
				s.check(err)
				if ok {
					cur_item.SetBool(val == 1)
				}
			}
		}
	}
	return v
}

func (s *Settings) UpdateSetting(field []string, value interface{}) bool {
	// TODO: update the database with the users setting
	return true
}
