package settings

import (
	"fmt"
	"reflect"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func getDefaultValues() SettingsValues {
	return SettingsValues{
		Appearence: SettingsValues_Appearence{
			Theme: "hanok",
			Font: SettingsValues_Appearence_Font{
				System:  "System",
				Greek:   "SBL Greek",
				Hebrew:  "SBL Hebrew",
				English: "Neuton",
			},
			Zoom: 100,
		},
		Shortcuts: SettingsValues_Shortcuts{
			AboutGraphe:       "",
			CheckForUpdates:   "",
			OpenSettings:      "cmdorctrl+,",
			OpenWorkspace:     "cmdorctrl+shift+,",
			OpenDataDirectory: "",
			OpenLogDirectory:  "",
			PurgeLogs:         "",

			PassageMode:   "cmdorctrl+P",
			SearchMode:    "cmdorctrl+F",
			OpenAnalytics: "cmdorctrl+\\",
			OpenFunctions: "cmdorctrl+]",
			ChooseVersion: "cmdorctrl+D",
			ChooseText:    "cmdorctrl+T",

			ZoomIn:      "cmdorctrl+plus",
			ZoomOut:     "cmdorctrl+-",
			ZoomReset:   "cmdorctrl+0",
			ChangeTheme: "",
		},
	}
}

func (s *Settings) GetSettings() SettingsValues {
	v := getDefaultValues()
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

func capitalise(s string) string {
	return cases.Title(language.English, cases.Compact).String(string(s[0])) + s[1:]
}

func (s *Settings) UpdateSetting(field []string, value interface{}) bool {
	if len(field) < 2 {
		s.throw(fmt.Sprintf("UpdateSetting has less than 2 values in `field`: (field: %v, value:...)", field))
	}

	r := reflect.TypeOf(SettingsValues{})

	table := capitalise(field[0])
	item, found := r.FieldByName(table)
	if !found {
		s.throw(fmt.Sprintf("UpdateSetting had invalid first-value for `field`: (field: %v, value:...)", field))
	}

	column := ""
	for i, f := range field[1:] {
		field_name := capitalise(f)
		item, found = item.Type.FieldByName(field_name)
		if !found {
			s.throw(fmt.Sprintf("UpdateSetting had invalid value (index=%d) for `field`: (field: %v, value:...)", i+1, field))
		}
		if i > 0 {
			column += "_"
		}
		column += field_name
	}

	var err error
	switch value.(type) {
	case float64:
		value = int(value.(float64))
		err = s.db.Exec(fmt.Sprintf(`
			UPDATE %s
			SET %s = %d
			WHERE id = 1;
		`, table, column, value))
	case string:
		if value == "DEFAULT" {
			err = s.db.Exec(fmt.Sprintf(`
				UPDATE %s
				SET %s = NULL
				WHERE id = 1;
			`, table, column))

		} else {
			err = s.db.Exec(fmt.Sprintf(`
				UPDATE %s
				SET %s = '%s'
				WHERE id = 1;
			`, table, column, value))
		}
	case bool:
		int_value := 0
		if value.(bool) {
			int_value = 1
		}
		err = s.db.Exec(fmt.Sprintf(`
			UPDATE %s
			SET %s = %d
			WHERE id = 1;
		`, table, column, int_value))
	default:
		s.throw(fmt.Sprintf("UpdateSetting had invalid `value` type (type: %s)", reflect.TypeOf(value).String()))
	}
	s.check(err)
	return true
}
