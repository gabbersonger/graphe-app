package app

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	data "graphe/internal/data"
	settings "graphe/internal/settings"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	env EnvironmentInfo
	ctx context.Context

	data     *data.DataDB
	settings *settings.SettingsDB
	menu     *menu.Menu
}

type EnvironmentInfo struct {
	Arch      string `json:"arch"`
	BuildType string `json:"buildType"`
	Platform  string `json:"platform"`
	Version   string `json:"version"`

	HomeDirectory string `json:"homeDirectory"`
	DataDirectory string `json:"dataDirectory"`
	LogDirectory  string `json:"logDirectory"`
}

func (a *App) assert(cond bool, msg string) {
	if !cond {
		if a.ctx != nil {
			runtime.LogFatal(a.ctx, msg)
		} else {
			panic(msg)
		}
	}
}

func NewApp(version string) *App {
	a := &App{}

	a.env.Version = strings.TrimSpace(version)
	a.assert(len(a.env.Version) > 0, "Invalid app version")

	homeDir, err := os.UserHomeDir()
	a.assert(err == nil, "No home directory")
	a.assert(len(homeDir) > 0, "Home directory had 0 length")
	a.env.HomeDirectory = homeDir

	a.env.DataDirectory = filepath.Join(a.env.HomeDirectory, "/Library/Application Support/Graphe")
	err = os.MkdirAll(a.env.DataDirectory, os.ModePerm)
	a.assert(err == nil, "Could not create/access folder for data directory")
	// FIXME: above only works for mac

	a.env.LogDirectory = filepath.Join(a.env.HomeDirectory, "/Library/Logs/Graphe")
	err = os.MkdirAll(a.env.LogDirectory, os.ModePerm)
	a.assert(err == nil, "Could not create/access folder for log directory")

	return a
}

func (a *App) Startup(ctx context.Context) {
	runtime.LogInfo(ctx, "Runcycle: Startup")

	a.ctx = ctx
	a.assert(ctx != nil, "Invalid context")

	wailsEnv := runtime.Environment(ctx)

	a.env.Arch = wailsEnv.Arch
	a.assert(len(a.env.Arch) > 0, "Invalid Arch")

	a.env.BuildType = wailsEnv.BuildType
	a.assert(len(a.env.BuildType) > 0, "Invalid BuildType")

	a.env.Platform = wailsEnv.Platform
	a.assert(len(a.env.Platform) > 0, "Invalid Platform")

	dbFile := a.env.DataDirectory + "/graphe.db"
	a.data = data.CreateDB(a.ctx, dbFile)
	a.assert(a.data != nil, "Invalid DataDB created")
	a.assert(a.data.Ping(), "Error connecting to DataDB")

	settingsFile := a.env.DataDirectory + "/settings.db"
	a.settings = settings.CreateDB(a.ctx, settingsFile)
	a.assert(a.settings != nil, "Invalid SettingsDB created")
	a.assert(a.settings.Ping(), "Error connecting to SettingsDB")

	a.LoadMenu()
	a.assert(a.menu != nil, "Error creating menu")
}

func (a *App) Shutdown(ctx context.Context) {
	runtime.LogInfo(ctx, "Runcycle: Shutdown")
	a.data.Shutdown()
	a.settings.Shutdown()
}

func (a *App) GetEnvironmentInfo() EnvironmentInfo {
	return a.env
}
