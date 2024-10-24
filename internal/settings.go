package internal

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
	Theme string                         `json:"theme"`
	Font  SettingsValues_Appearence_Font `json:"font"`
	Zoom  int                            `json:"zoom"`
}

type SettingsValues_Appearence_Font struct {
	System  string `json:"system"`
	Greek   string `json:"greek"`
	Hebrew  string `json:"hebrew"`
	English string `json:"english"`
}

type SettingsValues_Shortcuts struct {
	AboutGraphe       string `json:"aboutGraphe"`
	CheckForUpdates   string `json:"checkForUpdates"`
	OpenSettings      string `json:"openSettings"`
	OpenWorkspace     string `json:"openWorkspace"`
	OpenDataDirectory string `json:"openDataDirectory"`
	OpenLogDirectory  string `json:"openLogDirectory"`
	PurgeLogs         string `json:"purgeLogs"`

	PassageMode   string `json:"passageMode"`
	SearchMode    string `json:"searchMode"`
	OpenAnalytics string `json:"openAnalytics"`
	OpenFunctions string `json:"openFunctions"`
	ChooseVersion string `json:"chooseVersion"`
	ChooseText    string `json:"chooseText"`

	ZoomIn      string `json:"zoomIn"`
	ZoomOut     string `json:"zoomOut"`
	ZoomReset   string `json:"zoomReset"`
	ChangeTheme string `json:"changeTheme"`
}

type SettingsValues_Version struct{}

type SettingsValues_Formatting struct{}

type SettingsValues_Search struct{}

type SettingsValues_InstantDetails struct{}
