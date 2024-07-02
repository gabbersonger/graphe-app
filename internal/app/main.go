package app

import (
	"context"
	"graphe/internal/database"
	"graphe/internal/settings"
	"os"
	"path/filepath"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type EnvironmentInfo struct {
	Arch      string `json:"arch"`
	BuildType string `json:"buildType"`
	Platform  string `json:"platform"`
	Version   string `json:"version"`

	HomeDirectory string `json:"homeDirectory"`
	DataDirectory string `json:"dataDirectory"`
	LogDirectory  string `json:"logDirectory"`
}

type App struct {
	env      EnvironmentInfo
	ctx      context.Context
	db       *database.GrapheDB
	settings *settings.Settings
	menu     *menu.Menu
}

func (a *App) check(e error) {
	if e != nil {
		runtime.LogFatal(a.ctx, e.Error())
	}
}

func (a *App) GetEnvironmentInfo() EnvironmentInfo {
	return a.env
}

func (a *App) Throw(s string) {
	runtime.LogFatal(a.ctx, s)
}

func NewApp(version string) *App {
	a := &App{}
	a.env.Version = strings.TrimSpace(version)
	homeDir, err := os.UserHomeDir()
	a.check(err)
	a.env.HomeDirectory = homeDir
	a.env.DataDirectory = filepath.Join(a.env.HomeDirectory, "/Library/Application Support/Graphe")
	a.env.LogDirectory = filepath.Join(a.env.HomeDirectory, "/Library/Logs/Graphe")
	os.MkdirAll(a.env.DataDirectory, os.ModePerm)
	os.MkdirAll(a.env.LogDirectory, os.ModePerm)
	return a
}

func (a *App) Startup(ctx context.Context) {
	runtime.LogInfo(ctx, "Runcycle: Startup")

	a.ctx = ctx

	wailsEnv := runtime.Environment(ctx)
	a.env.Arch = wailsEnv.Arch
	a.env.BuildType = wailsEnv.BuildType
	a.env.Platform = wailsEnv.Platform

	dbFile := a.env.DataDirectory + "/graphe.db"
	a.db = database.Startup(a.ctx, dbFile)

	settingsFile := a.env.DataDirectory + "/settings.db"
	a.settings = settings.Startup(a.ctx, settingsFile)
	a.LoadMenu()
}

func (a *App) Shutdown(ctx context.Context) {
	runtime.LogInfo(ctx, "Runcycle: Shutdown")

	a.db.Shutdown()
	a.settings.Shutdown()
}

func (a *App) GetMenu() *menu.Menu {
	return a.menu
}
