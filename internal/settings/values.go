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

func (s *SettingsDB) GetSettings() SettingsValues {
	v := getDefaultValues()

	r := reflect.ValueOf(&v).Elem()
	for _, t := range getSettingsTables() {
		t_item := r.FieldByName(t.name)
		for _, c := range t.columns {
			// Get item for column (e.g. Appearence>Theme for "appearence_theme")
			item := t_item
			for _, part := range strings.Split(c.name, "_") {
				item = item.FieldByName(part)
				s.assert(item.IsValid(), "Error matchings struct to column name from getSettingsTables()")
			}

			// Get item from db
			stmt, err := s.db.Prepare(fmt.Sprintf(`
				SELECT %s
				FROM %s
				WHERE id = 1
				LIMIT 1;
			`, c.name, t.name))
			s.assert(err == nil, fmt.Sprintf("Error preparing query for getting setting value (table: %s, column: %s)", t.name, c.name))
			has_row, err := stmt.Step()
			s.assert(err == nil, fmt.Sprintf("Error executing query for getting setting value (table: %s, column: %s)", t.name, c.name))
			s.assert(has_row, fmt.Sprintf("Error getting row for setting (table: %s, column: %s)", t.name, c.name))

			// Place item into the struct
			switch item.Kind() {
			case reflect.Int:
				val, ok, err := stmt.ColumnInt64(0)
				s.assert(err == nil, fmt.Sprintf("Error getting setting as int (table: %s, column: %s)", t.name, c.name))
				if ok {
					item.SetInt(val)
				}
			case reflect.String:
				val, ok, err := stmt.ColumnText(0)
				s.assert(err == nil, fmt.Sprintf("Error getting setting as string (table: %s, column: %s)", t.name, c.name))
				if ok {
					item.SetString(val)
				}
			case reflect.Bool:
				val, ok, err := stmt.ColumnInt(0)
				s.assert(err == nil, fmt.Sprintf("Error getting setting as string (table: %s, column: %s)", t.name, c.name))
				if ok {
					item.SetBool(val == 1)
				}
			default:
				s.assert(false, fmt.Sprintf("Invalid variable format (table: %s, column: %s)", t.name, c.name))
			}

			stmt.Close()
		}
	}

	return v
}

func (s *SettingsDB) ResetSetting(key []string) interface{} {
	item := reflect.ValueOf(getDefaultValues())
	for i, k := range key {
		k_name := capitalise(k)
		item = reflect.Indirect(item).FieldByName(k_name)
		s.assert(item.IsValid(), fmt.Sprintf("Invalid key (key: %v, index: %d)", key, i))
	}
	resetValue(s, item, key)
	return nil
}

func capitalise(s string) string {
	return cases.Title(language.English, cases.Compact).String(string(s[0])) + s[1:]
}

func resetValue(s *SettingsDB, item reflect.Value, key []string) {
	s.assert(item.IsValid(), fmt.Sprintf("Invalid item (item: %v, key: %v)", item, key))
	s.assert(len(key) > 0, "Invalid key (length = 0)")
	switch item.Kind() {
	case reflect.Struct:
		fields := reflect.VisibleFields(item.Type())
		new_key := append(key, "")
		for _, f := range fields {
			child_item := item.FieldByName(f.Name)
			new_key[len(new_key)-1] = f.Name
			resetValue(s, child_item, new_key)
		}
	default:
		s.UpdateSetting(key, item.Interface())
	}
}

func (s *SettingsDB) UpdateSetting(key []string, val interface{}) bool {
	s.assert(len(key) >= 2, fmt.Sprintf("Not enough values in key provided (key: %v, val: %v)", key, val))

	// Get the table
	table_name := capitalise(key[0])
	r := reflect.TypeOf(getDefaultValues())
	item, found := r.FieldByName(table_name)
	s.assert(found, fmt.Sprintf("Invalid first value (table name) in key (key: %v, val: %v)", key, val))

	// Get the specific field's column name
	column_name := ""
	for i, k := range key[1:] {
		k_name := capitalise(k)
		item, found = item.Type.FieldByName(k_name)
		s.assert(found, fmt.Sprintf("Invalid value in key - did not match struct (key: %v [index: %d], val: %v)", key, i, val))
		if i > 0 {
			column_name += "_"
		}
		column_name += k_name
	}

	// Update the value in the db
	switch val.(type) {
	case float64:
		val := int(val.(float64))
		execUpdate(s, fmt.Sprintf(`
			UPDATE %s SET %s = %d WHERE id = 1;
		`, table_name, column_name, val))
	case string:
		if val == "DEFAULT" {
			execUpdate(s, fmt.Sprintf(`
				UPDATE %s SET %s = NULL WHERE id = 1;
			`, table_name, column_name))
		} else {
			execUpdate(s, fmt.Sprintf(`
				UPDATE %s SET %s = '%s' WHERE id = 1;
			`, table_name, column_name, val))
		}
	case bool:
		int_val := 0
		if val.(bool) {
			int_val = 1
		}
		execUpdate(s, fmt.Sprintf(`
			UPDATE %s SET %s = %d WHERE id = 1;
		`, table_name, column_name, int_val))
	default:
		s.assert(false, fmt.Sprintf("Invalid value format (key: %v, val: %v)", key, val))
	}

	s.logger.Info(fmt.Sprintf("Updated setting (key: %v, val: %v)", key, val))
	return true
}

func execUpdate(s *SettingsDB, query string) {
	err := s.db.Begin()
	s.assert(err == nil, "Error beginning transaction")
	err = s.db.Exec(query)
	s.assert(err == nil, fmt.Sprintf("Error executing query (query: `%s`)", strings.TrimSpace(query)))
	num_updates := s.db.Changes()
	if num_updates > 1 {
		s.db.Rollback()
		s.assert(false, fmt.Sprintf("More than one row was updated for query (query: `%s`, num_updates: %d)", strings.TrimSpace(query), num_updates))
	}
	s.db.Commit()
}
