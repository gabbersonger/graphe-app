package menu

import (
	"github.com/wailsapp/wails/v3/pkg/application"
)

type SettingValue struct {
	Setting []string `json:"setting"`
	Value   any      `json:"value"`
}

func (m *MenuManager) menuItemEventCallback(eventName string, data ...any) func(*application.Context) {
	return func(ctx *application.Context) {
		m.app.EmitEvent(eventName, data)
	}
}

func (m *MenuManager) addMenuItem(submenu *application.Menu, label string, shortcut string, callback func(*application.Context)) *application.MenuItem {
	menu_item := submenu.Add(label)
	if len(shortcut) > 0 {
		menu_item.SetAccelerator(shortcut)
	}
	menu_item.OnClick(callback)
	return menu_item
}

func (m *MenuManager) addGrapheSubmenu(menu *application.Menu) {
	grapheMenu := menu.AddSubmenu("Graphe")
	m.addMenuItem(grapheMenu, "About Graphe", m.shortcuts.AboutGraphe, func(*application.Context) {
		m.app.ShowAboutDialog()
	})
	grapheMenu.AddSeparator()
	m.addMenuItem(grapheMenu, "Check for Updates", m.shortcuts.CheckForUpdates, func(*application.Context) {
		m.app.EmitEvent("window:settings:section", "version")
	})
	grapheMenu.AddSeparator()
	m.addMenuItem(grapheMenu, "Settings", m.shortcuts.OpenSettings, func(*application.Context) {
		m.app.EmitEvent("graphe:mode", "settings")
	})
	m.addMenuItem(grapheMenu, "Workspace", m.shortcuts.OpenWorkspace, func(*application.Context) {
		m.app.EmitEvent("graphe:mode", "workspace")
	})
	grapheMenu.AddSeparator()
	grapheMenu.AddRole(application.Hide)
	grapheMenu.AddRole(application.CloseWindow)
	grapheMenu.AddRole(application.Quit)
	menu.Append(grapheMenu)
}

func (m *MenuManager) addEditSubmenu(menu *application.Menu) {
	editMenu := menu.AddSubmenu("Edit")
	editMenu.AddRole(application.Undo)
	editMenu.AddRole(application.Redo)
	editMenu.AddSeparator()
	editMenu.AddRole(application.Cut)
	editMenu.AddRole(application.Copy)
	editMenu.AddRole(application.Paste)
	editMenu.AddRole(application.PasteAndMatchStyle)
	editMenu.AddRole(application.Delete)
	editMenu.AddRole(application.SelectAll)
	menu.Append(editMenu)
}

func (m *MenuManager) addWorkspaceSubmenu(menu *application.Menu) {
	workspaceMenu := menu.AddSubmenu("Workspace")
	m.addMenuItem(workspaceMenu, "Passage Mode", m.shortcuts.PassageMode, func(*application.Context) {
		m.app.EmitEvent("window:workspace:mode", "passage")
	})
	m.addMenuItem(workspaceMenu, "Search Mode", m.shortcuts.SearchMode, func(*application.Context) {
		m.app.EmitEvent("window:workspace:mode", "search")
	})
	workspaceMenu.AddSeparator()
	m.addMenuItem(workspaceMenu, "Analytics", m.shortcuts.OpenAnalytics, func(*application.Context) {
		m.app.EmitEvent("window:workspace:sidebar", "toggle")
	})
	m.addMenuItem(workspaceMenu, "Functions", m.shortcuts.OpenFunctions, func(*application.Context) {
		m.app.EmitEvent("window:workspace:modal", "functions")
	})
	workspaceMenu.AddSeparator()
	m.addMenuItem(workspaceMenu, "Choose Version...", m.shortcuts.ChooseVersion, func(*application.Context) {
		m.app.EmitEvent("window:workspace:modal", "version")
	})
	m.addMenuItem(workspaceMenu, "Choose Text...", m.shortcuts.ChooseText, func(*application.Context) {
		m.app.EmitEvent("window:workspace:modal", "text")
	})
	menu.Append(workspaceMenu)
}

func (m *MenuManager) addSearchSubmenu(menu *application.Menu) {
	searchMenu := menu.AddSubmenu("Search")
	searchMenu.Add("TODO") // TODO
	menu.Append(searchMenu)
}

func (m *MenuManager) addFunctionsSubmenu(menu *application.Menu) {
	functionsMenu := menu.AddSubmenu("Functions")
	functionsMenu.Add("TODO") // TODO
	menu.Append(functionsMenu)
}

func (m *MenuManager) addViewSubmenu(menu *application.Menu) {
	viewMenu := menu.AddSubmenu("View")
	m.addMenuItem(viewMenu, "Zoom In", m.shortcuts.ZoomIn, func(*application.Context) {
		m.app.EmitEvent("graphe:setting", SettingValue{
			Setting: []string{"appearence", "zoom"},
			Value:   "in",
		})
	})
	m.addMenuItem(viewMenu, "Zoom Out", m.shortcuts.ZoomOut, func(*application.Context) {
		m.app.EmitEvent("graphe:setting", SettingValue{
			Setting: []string{"appearence", "zoom"},
			Value:   "out",
		})
	})
	m.addMenuItem(viewMenu, "Reset Zoom", m.shortcuts.ZoomReset, func(*application.Context) {
		m.app.EmitEvent("graphe:setting", SettingValue{
			Setting: []string{"appearence", "zoom"},
			Value:   "reset",
		})
	})
	viewMenu.AddSeparator()
	viewMenu.AddRole(application.Minimize)
	viewMenu.AddSeparator()
	m.addMenuItem(viewMenu, "Change Theme...", m.shortcuts.ChangeTheme, func(*application.Context) {
		m.app.EmitEvent("window:settings:section", "appearence")
	})
	viewMenu.AddSeparator()
	menu.Append(viewMenu)
}

func (m *MenuManager) updateMenu() {
	if m.app == nil || m.shortcuts == nil {
		m.log(`Tried updating menu prior to having valid app or shortcut values`)
		return
	}

	menu := m.app.NewMenu()
	m.addGrapheSubmenu(menu)
	m.addEditSubmenu(menu)
	m.addWorkspaceSubmenu(menu)
	m.addSearchSubmenu(menu)
	m.addFunctionsSubmenu(menu)
	m.addViewSubmenu(menu)
	m.app.SetMenu(menu)
}
