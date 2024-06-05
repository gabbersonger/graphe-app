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

func (a *App) Menu() *menu.Menu {
	appMenu := menu.NewMenu()

	grapheMenu := appMenu.AddSubmenu("Graphe")
	grapheMenu.AddText("About Graphe", nil, menuCallbackEmit(a, "global:about"))
	grapheMenu.AddSeparator()
	grapheMenu.AddText("Settings", keys.CmdOrCtrl(","), menuCallbackEmit(a, "global:settings"))
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
	fileMenu.AddText("Passage Mode", keys.CmdOrCtrl("P"), menuCallbackEmit(a, "app:mode", "passage"))
	fileMenu.AddText("Search Mode", keys.CmdOrCtrl("F"), menuCallbackEmit(a, "app:mode", "search"))
	fileMenu.AddSeparator()
	fileMenu.AddText("Choose Version...", keys.CmdOrCtrl("D"), menuCallbackEmit(a, "ui:modal", "version"))
	fileMenu.AddText("Choose Text...", keys.CmdOrCtrl("T"), menuCallbackEmit(a, "ui:modal", "text"))
	fileMenu.AddSeparator()
	fileMenu.AddText("Functions", keys.CmdOrCtrl("E"), menuCallbackEmit(a, "ui:modal", "functions"))
	fileMenu.AddText("Appearence", keys.CmdOrCtrl("R"), menuCallbackEmit(a, "ui:modal", "appearence"))
	fileMenu.AddSeparator()
	fileMenu.AddText("Sidebar", keys.CmdOrCtrl("\\"), menuCallbackEmit(a, "ui:sidebar:toggle"))
	fileMenu.AddSeparator()
	fileMenu.AddText("Theme toggle", keys.CmdOrCtrl("L"), menuCallbackEmit(a, "ui:theme:toggle"))

	appMenu.Append(menu.EditMenu())

	return appMenu
}
