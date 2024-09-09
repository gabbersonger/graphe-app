package main

import (
	"embed"
	"graphe/internal/data"
	"graphe/internal/logger"
	"graphe/internal/scripture"
	"graphe/internal/settings"
	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	logger := logger.NewGrapheLogger()
	scripture_service := scripture.NewScriptureService(logger)
	settings_db := settings.NewSettingsDB(logger)
	data_db := data.NewDataDB(logger, scripture_service)

	app := application.New(application.Options{
		Name:        "Graphe",
		Description: "Original language Bible study",
		Services: []application.Service{
			application.NewService(scripture_service),
			application.NewService(settings_db),
			application.NewService(data_db),
		},
		Logger: logger,
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title: "Graphe",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(27, 38, 54),
		URL:              "/",
	})

	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
