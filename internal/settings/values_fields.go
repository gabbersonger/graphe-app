package settings

type SettingsValues struct {
	General        SettingsValues_General        `json:"general"`
	Appearence     SettingsValues_Appearence     `json:"appearence"`
	Shortcuts      SettingsValues_Shortcuts      `json:"shortcuts"`
	Version        SettingsValues_Version        `json:"version"`
	Formatting     SettingsValues_Formatting     `json:"formatting"`
	Search         SettingsValues_Search         `json:"search"`
	InstantDetails SettingsValues_InstantDetails `json:"instantDetails"`
}

type SettingsValues_General struct{}

type SettingsValues_Appearence struct {
	Theme string                         `json:"theme,omitempty"`
	Font  SettingsValues_Appearence_Font `json:"font"`
	Zoom  int                            `json:"zoom,omitempty"`
}

type SettingsValues_Appearence_Font struct {
	System  string `json:"system,omitempty"`
	Greek   string `json:"greek,omitempty"`
	Hebrew  string `json:"hebrew,omitempty"`
	English string `json:"english,omitempty"`
}

type SettingsValues_Shortcuts struct{}

type SettingsValues_Version struct{}

type SettingsValues_Formatting struct{}

type SettingsValues_Search struct{}

type SettingsValues_InstantDetails struct{}
