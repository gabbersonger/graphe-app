package app

import (
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

func (a *App) newMenu() *menu.Menu {
	// TODO: fill this out and make shortcuts based on settings values

	appMenu := menu.NewMenu()

	grapheMenu := appMenu.AddSubmenu("Graphe")
	grapheMenu.AddText("About Graphe", nil, menuCallbackEmit(a, "global:about"))
	grapheMenu.AddText("Check for Updates", nil, menuCallbackEmit(a, "global:update_check"))
	grapheMenu.AddSeparator()
	grapheMenu.AddText("Settings", keys.CmdOrCtrl(","), menuCallbackEmit(a, "graphe:mode", "settings"))
	grapheMenu.AddText("Workspace", keys.CmdOrCtrl("<"), menuCallbackEmit(a, "graphe:mode", "workspace"))
	grapheMenu.AddSeparator()
	grapheMenu.AddText("Open Data Directory", nil, menuCallbackOpenFolder(a, a.env.DataDirectory))
	grapheMenu.AddText("Open Log Directory", nil, menuCallbackOpenFolder(a, a.env.LogDirectory))
	grapheMenu.AddText("Purge logs", nil, menuCallbackEmit(a, "global:purge_logs"))
	grapheMenu.AddSeparator()
	grapheMenu.AddText("Hide Graphe", keys.CmdOrCtrl("H"), func(cd *menu.CallbackData) { runtime.Hide(a.ctx) })
	grapheMenu.AddText("Quit Graphe", keys.CmdOrCtrl("Q"), func(cd *menu.CallbackData) { runtime.Quit(a.ctx) })

	appMenu.Append(menu.EditMenu())

	workspaceMenu := appMenu.AddSubmenu("Workspace")
	workspaceMenu.AddText("Passage Mode", keys.CmdOrCtrl("P"), menuCallbackEmit(a, "window:workspace:mode", "passage"))
	workspaceMenu.AddText("Search Mode", keys.CmdOrCtrl("F"), menuCallbackEmit(a, "window:workspace:mode", "search"))
	workspaceMenu.AddSeparator()
	workspaceMenu.AddText("Functions", keys.CmdOrCtrl("]"), menuCallbackEmit(a, "window:workspace:modal", "functions"))
	workspaceMenu.AddText("Analytics", keys.CmdOrCtrl("\\"), menuCallbackEmit(a, "window:workspace:sidebar", "toggle"))
	workspaceMenu.AddSeparator()
	workspaceMenu.AddText("Choose Version...", keys.CmdOrCtrl("D"), menuCallbackEmit(a, "window:workspace:modal", "version"))
	workspaceMenu.AddText("Choose Text...", keys.CmdOrCtrl("T"), menuCallbackEmit(a, "window:workspace:modal", "text"))

	_ = appMenu.AddSubmenu("Search")
	_ = appMenu.AddSubmenu("Functions")

	viewMenu := appMenu.AddSubmenu("View")
	viewMenu.AddText("Zoom In", keys.CmdOrCtrl("+"), menuCallbackEmit(a, "graphe:setting", []string{"appearence", "zoom"}, "in"))
	viewMenu.AddText("Zoom Out", keys.CmdOrCtrl("-"), menuCallbackEmit(a, "graphe:setting", []string{"appearence", "zoom"}, "out"))
	viewMenu.AddText("Reset Zoom", keys.CmdOrCtrl("0"), menuCallbackEmit(a, "graphe:setting", []string{"appearence", "zoom"}, "reset"))
	viewMenu.AddSeparator()
	viewMenu.AddText("Change Theme...", nil, menuCallbackEmit(a, "window:settings:section", "appearence"))
	viewMenu.AddSeparator()

	return appMenu
}

func (a *App) ChangeMenu() {
	a.menu = a.newMenu()
	runtime.MenuSetApplicationMenu(a.ctx, a.menu)
	runtime.MenuUpdateApplicationMenu(a.ctx)
}
