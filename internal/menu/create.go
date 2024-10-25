package menu

import "github.com/wailsapp/wails/v3/pkg/application"

func (m *MenuManager) menuItemEventCallback(eventName string, data ...any) func(*application.Context) {
	return func(ctx *application.Context) {
		m.app.EmitEvent(eventName, data)
	}
}

func (m *MenuManager) addMenuItem(submenu *application.Menu, label string, shortcut string, callback func(*application.Context)) *application.MenuItem {
	menu_item := submenu.Add(label)
	menu_item.SetAccelerator(shortcut)
	menu_item.OnClick(callback)
	return menu_item
}

func (m *MenuManager) addGrapheSubmenu(menu *application.Menu) {
	grapheMenu := menu.AddSubmenu("Graphe")
	m.addMenuItem(grapheMenu, "One thing", "Ctrl+Option+A", m.menuItemEventCallback("graphe:one"))
	m.addMenuItem(grapheMenu, "Something else", "Ctrl+Option+B", m.menuItemEventCallback("graphe:two"))
	grapheMenu.AddSeparator()
	m.addMenuItem(grapheMenu, "Another thing", "Ctrl+Option+C", m.menuItemEventCallback("graphe:three"))
	menu.Append(grapheMenu)
}

func (m *MenuManager) addEditSubmenu(menu *application.Menu) {
	menu.AddRole(application.EditMenu)
	// TODO
}

func (m *MenuManager) addWorkspaceSubmenu(menu *application.Menu) {
	// TODO
}

func (m *MenuManager) addSearchSubmenu(menu *application.Menu) {
	// TODO
}

func (m *MenuManager) addFunctionsSubmenu(menu *application.Menu) {
	// TODO
}

func (m *MenuManager) addViewSubmenu(menu *application.Menu) {
	// TODO
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
