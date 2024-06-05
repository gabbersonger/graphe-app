package main

import (
	"embed"
	"graphe/internal"
	"graphe/internal/app"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

var (
	//go:embed all:frontend/dist
	assets embed.FS

	//go:embed build/appicon.png
	appIcon []byte

	//go:embed build/info/title.txt
	title string

	//go:embed build/info/version.txt
	version string

	//go:embed build/info/comment.txt
	comment string

	//go:embed build/info/copyright.txt
	copyright string
)

func main() {
	// Create an instance of the app structure
	app := app.NewApp(version)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  title,
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		Menu:             app.Menu(),
		OnStartup:        app.Startup,
		OnShutdown:       app.Shutdown,
		Bind: []interface{}{
			app,
		},

		Logger:             internal.NewAppLogger(app.GetEnvironmentInfo().LogDirectory, title+".log"),
		LogLevel:           logger.TRACE,
		LogLevelProduction: logger.INFO,

		Mac: &mac.Options{
			TitleBar: mac.TitleBarHiddenInset(),
			About: &mac.AboutInfo{
				Title:   title + "\n" + comment,
				Message: copyright,
				Icon:    appIcon,
			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
