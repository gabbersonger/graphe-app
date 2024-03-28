package app

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const max_db_conn = 5

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
	Env     EnvironmentInfo
	ctx     context.Context
	db_pool chan *GrapheDB
}

func (a *App) check(e error) {
	if e != nil {
		runtime.LogFatal(a.ctx, e.Error())
	}
}

func (a *App) Throw(s string) {
	runtime.LogFatal(a.ctx, s)
}

func NewApp(version string) *App {
	a := &App{}
	a.Env.Version = strings.TrimSpace(version)
	homeDir, err := os.UserHomeDir()
	a.check(err)
	a.Env.HomeDirectory = homeDir
	a.Env.DataDirectory = filepath.Join(a.Env.HomeDirectory, "/Library/Application Support/Graphe")
	a.Env.LogDirectory = filepath.Join(a.Env.HomeDirectory, "/Library/Logs/Graphe")
	os.MkdirAll(a.Env.DataDirectory, os.ModePerm)
	os.MkdirAll(a.Env.LogDirectory, os.ModePerm)
	return a
}

func (a *App) Startup(ctx context.Context) {
	runtime.LogInfo(ctx, "Runcycle: Startup")

	a.ctx = ctx

	wailsEnv := runtime.Environment(ctx)
	a.Env.Arch = wailsEnv.Arch
	a.Env.BuildType = wailsEnv.BuildType
	a.Env.Platform = wailsEnv.Platform

	a.setupDatabasePool()
}

func (a *App) Shutdown(ctx context.Context) {
	runtime.LogInfo(ctx, "Runcycle: Shutdown")

	a.closeDatabasePool()
}
