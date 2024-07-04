package app

import (
	"fmt"
	"graphe/internal/settings"
	"os/exec"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func menuCallbackEmit(a *App, eventName string, data ...interface{}) func(cd *menu.CallbackData) {
	return func(cd *menu.CallbackData) {
		runtime.EventsEmit(a.ctx, eventName, data...)
	}
}

func menuCallbackOpenFolder(a *App, fname string) func(cd *menu.CallbackData) {
	return func(cd *menu.CallbackData) {
		switch a.env.Platform {
		case "darwin":
			exec.Command("open", "--reveal", fname).Run()
			break
		case "linux":
			exec.Command("xdg-open", fname).Run()
			break
		case "windows":
			exec.Command("explorer", fname).Run()
			break
		default:
			a.Throw("Unknown operating system for revealing folder")
		}
	}
}

func (a *App) shortcutToKeyCode(s string) *keys.Accelerator {
	if len(s) == 0 {
		return nil
	}
	shortcut, err := keys.Parse(s)
	if err != nil {
		fmt.Println(err.Error())
		a.Throw(fmt.Sprintf("Error parsing shortcut: %s", s))
	}
	return shortcut
}

func (a *App) newMenu(s settings.SettingsValues_Shortcuts) *menu.Menu {
	appMenu := menu.NewMenu()

	grapheMenu := appMenu.AddSubmenu("Graphe")
	grapheMenu.AddText("About Graphe", a.shortcutToKeyCode(s.AboutGraphe), menuCallbackEmit(a, "global:about"))
	grapheMenu.AddText("Check for Updates", a.shortcutToKeyCode(s.CheckForUpdates), menuCallbackEmit(a, "window:settings:section", "version"))
	grapheMenu.AddSeparator()
	grapheMenu.AddText("Settings", a.shortcutToKeyCode(s.OpenSettings), menuCallbackEmit(a, "graphe:mode", "settings"))
	grapheMenu.AddText("Workspace", a.shortcutToKeyCode(s.OpenWorkspace), menuCallbackEmit(a, "graphe:mode", "workspace"))
	grapheMenu.AddSeparator()
	grapheMenu.AddText("Open Data Directory", a.shortcutToKeyCode(s.OpenDataDirectory), menuCallbackOpenFolder(a, a.env.DataDirectory))
	grapheMenu.AddText("Open Log Directory", a.shortcutToKeyCode(s.OpenLogDirectory), menuCallbackOpenFolder(a, a.env.LogDirectory))
	grapheMenu.AddText("Purge logs", a.shortcutToKeyCode(s.PurgeLogs), menuCallbackEmit(a, "global:purge_logs"))
	grapheMenu.AddSeparator()
	grapheMenu.AddText("Hide Graphe", keys.CmdOrCtrl("H"), func(cd *menu.CallbackData) { runtime.Hide(a.ctx) })
	grapheMenu.AddText("Close Window", keys.CmdOrCtrl("W"), func(cd *menu.CallbackData) { runtime.Quit(a.ctx) })
	grapheMenu.AddText("Quit Graphe", keys.CmdOrCtrl("Q"), func(cd *menu.CallbackData) { runtime.Quit(a.ctx) })

	appMenu.Append(menu.EditMenu())

	workspaceMenu := appMenu.AddSubmenu("Workspace")
	workspaceMenu.AddText("Passage Mode", a.shortcutToKeyCode(s.PassageMode), menuCallbackEmit(a, "window:workspace:mode", "passage"))
	workspaceMenu.AddText("Search Mode", a.shortcutToKeyCode(s.SearchMode), menuCallbackEmit(a, "window:workspace:mode", "search"))
	workspaceMenu.AddSeparator()
	workspaceMenu.AddText("Analytics", a.shortcutToKeyCode(s.OpenAnalytics), menuCallbackEmit(a, "window:workspace:sidebar", "toggle"))
	workspaceMenu.AddText("Functions", a.shortcutToKeyCode(s.OpenFunctions), menuCallbackEmit(a, "window:workspace:modal", "functions"))
	workspaceMenu.AddSeparator()
	workspaceMenu.AddText("Choose Version...", a.shortcutToKeyCode(s.ChooseVersion), menuCallbackEmit(a, "window:workspace:modal", "version"))
	workspaceMenu.AddText("Choose Text...", a.shortcutToKeyCode(s.ChooseText), menuCallbackEmit(a, "window:workspace:modal", "text"))

	searchMenu := appMenu.AddSubmenu("Search")
	searchMenu.AddText("TODO", nil, nil)

	functionsMenu := appMenu.AddSubmenu("Functions")
	functionsMenu.AddText("TODO", nil, nil)

	viewMenu := appMenu.AddSubmenu("View")
	viewMenu.AddText("Zoom In", a.shortcutToKeyCode(s.ZoomIn), menuCallbackEmit(a, "graphe:setting", []string{"appearence", "zoom"}, "in"))
	viewMenu.AddText("Zoom Out", a.shortcutToKeyCode(s.ZoomOut), menuCallbackEmit(a, "graphe:setting", []string{"appearence", "zoom"}, "out"))
	viewMenu.AddText("Reset Zoom", a.shortcutToKeyCode(s.ZoomReset), menuCallbackEmit(a, "graphe:setting", []string{"appearence", "zoom"}, "reset"))
	viewMenu.AddSeparator()
	viewMenu.AddText("Minimise Window", keys.CmdOrCtrl("M"), func(cd *menu.CallbackData) { runtime.WindowMinimise(a.ctx) })
	viewMenu.AddSeparator()
	viewMenu.AddText("Change Theme...", a.shortcutToKeyCode(s.ChangeTheme), menuCallbackEmit(a, "window:settings:section", "appearence"))
	viewMenu.AddSeparator()

	return appMenu
}

func (a *App) LoadMenu() {
	a.updateMenu(a.settings.GetSettings().Shortcuts)
}

func (a *App) DisableMenu() {
	s := settings.SettingsValues_Shortcuts{}
	a.updateMenu(s)
}

func (a *App) updateMenu(s settings.SettingsValues_Shortcuts) {
	a.menu = a.newMenu(s)
	runtime.MenuSetApplicationMenu(a.ctx, a.menu)
	runtime.MenuUpdateApplicationMenu(a.ctx)
}
