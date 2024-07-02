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

type SettingsValues_Shortcuts struct {
	AboutGraphe       string `json:"aboutGraphe,omitempty"`
	CheckForUpdates   string `json:"checkForUpdates,omitempty"`
	OpenSettings      string `json:"openSettings,omitempty"`
	OpenWorkspace     string `json:"openWorkspace,omitempty"`
	OpenDataDirectory string `json:"openDataDirectory,omitempty"`
	OpenLogDirectory  string `json:"openLogDirectory,omitempty"`
	PurgeLogs         string `json:"purgeLogs,omitempty"`

	PassageMode   string `json:"passageMode,omitempty"`
	SearchMode    string `json:"searchMode,omitempty"`
	OpenAnalytics string `json:"openAnalytics,omitempty"`
	OpenFunctions string `json:"openFunctions,omitempty"`
	ChooseVersion string `json:"chooseVersion,omitempty"`
	ChooseText    string `json:"chooseText,omitempty"`

	ZoomIn      string `json:"zoomIn,omitempty"`
	ZoomOut     string `json:"zoomOut,omitempty"`
	ZoomReset   string `json:"zoomReset,omitempty"`
	ChangeTheme string `json:"changeTheme,omitempty"`
}

type SettingsValues_Version struct{}

type SettingsValues_Formatting struct{}

type SettingsValues_Search struct{}

type SettingsValues_InstantDetails struct{}
