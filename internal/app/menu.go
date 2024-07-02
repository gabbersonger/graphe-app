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
	grapheMenu.AddSeparator()
	grapheMenu.AddText("Settings", keys.CmdOrCtrl(","), menuCallbackEmit(a, "graphe:mode", "settings"))
	grapheMenu.AddText("App", keys.CmdOrCtrl("<"), menuCallbackEmit(a, "graphe:mode", "app"))
	grapheMenu.AddSeparator()
	grapheMenu.AddText("Check for Updates", nil, menuCallbackEmit(a, "global:update_check"))
	grapheMenu.AddSeparator()
	grapheMenu.AddText("Open Data Directory", nil, menuCallbackOpenFolder(a, a.env.DataDirectory))
	grapheMenu.AddText("Open Log Directory", nil, menuCallbackOpenFolder(a, a.env.LogDirectory))
	grapheMenu.AddText("Purge logs", nil, menuCallbackEmit(a, "global:purge_logs"))
	grapheMenu.AddSeparator()
	grapheMenu.AddText("Hide Graphe", keys.CmdOrCtrl("H"), func(cd *menu.CallbackData) { runtime.Hide(a.ctx) })
	grapheMenu.AddText("Quit Graphe", keys.CmdOrCtrl("Q"), func(cd *menu.CallbackData) { runtime.Quit(a.ctx) })

	fileMenu := appMenu.AddSubmenu("File")
	fileMenu.AddText("Passage Mode", keys.CmdOrCtrl("P"), menuCallbackEmit(a, "window:workspace:mode", "passage"))
	fileMenu.AddText("Search Mode", keys.CmdOrCtrl("F"), menuCallbackEmit(a, "window:workspace:mode", "search"))
	fileMenu.AddSeparator()
	fileMenu.AddText("Choose Version...", keys.CmdOrCtrl("D"), menuCallbackEmit(a, "window:workspace:modal", "version"))
	fileMenu.AddText("Choose Text...", keys.CmdOrCtrl("T"), menuCallbackEmit(a, "window:workspace:modal", "text"))
	fileMenu.AddSeparator()
	fileMenu.AddText("Functions", keys.CmdOrCtrl("E"), menuCallbackEmit(a, "window:workspace:sidebar", "functions"))
	fileMenu.AddText("Appearence", keys.CmdOrCtrl("R"), menuCallbackEmit(a, "window:workspace:sidebar", "appearence"))
	fileMenu.AddSeparator()
	fileMenu.AddText("Sidebar", keys.CmdOrCtrl("\\"), menuCallbackEmit(a, "window:workspace:sidebar:toggle"))
	fileMenu.AddSeparator()
	fileMenu.AddText("Reset Zoom", keys.CmdOrCtrl("0"), menuCallbackEmit(a, "graphe:setting", []string{"appearence", "zoom"}, "reset"))
	fileMenu.AddText("Zoom In", keys.CmdOrCtrl("+"), menuCallbackEmit(a, "graphe:setting", []string{"appearence", "zoom"}, "in"))
	fileMenu.AddText("Zoom Out", keys.CmdOrCtrl("-"), menuCallbackEmit(a, "graphe:setting", []string{"appearence", "zoom"}, "out"))

	appMenu.Append(menu.EditMenu())

	return appMenu
}

func (a *App) ChangeMenu() {
	a.menu = a.newMenu()
	runtime.MenuSetApplicationMenu(a.ctx, a.menu)
	runtime.MenuUpdateApplicationMenu(a.ctx)
}
