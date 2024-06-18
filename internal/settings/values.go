package settings

type SettingsValues struct {
	General    struct{}
	Appearence struct {
		Theme string
		Font  struct {
			System  string
			Greek   string
			Hebrew  string
			English string
		}
	}
	Shortcuts      struct{}
	Version        struct{}
	Formatting     struct{}
	Search         struct{}
	InstantDetails struct{}
}

func (s *Settings) setupValues() {
	s.values = &SettingsValues{}
}

func (s *Settings) GetSettings() *SettingsValues {
	return s.values
}

func (s *Settings) UpdateSetting(group string, field string, value string) bool {
	// TODO
	return true
}
