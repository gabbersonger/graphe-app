package internal

import (
	"log"
	"os"
	"path"
	"strings"
)

var showError = true

type AppLogger struct {
	directory string
	filename  string
	filepath  string
}

func NewAppLogger(directory, filename string) *AppLogger {
	return &AppLogger{
		directory: directory,
		filename:  filename,
		filepath:  path.Join(directory, filename),
	}
}

func (l *AppLogger) Print(message string) {
	os.MkdirAll(l.directory, os.ModePerm)
	f, err := os.OpenFile(l.filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil && showError {
		log.Fatal("Could not open logfile")
	}

	if _, err = f.WriteString(message); err != nil {
		if showError {
			log.Fatal("Could not write to logfile")
			showError = false
		} else {
			showError = true
		}
	} else {
		showError = true
	}

	f.Close()
}

func (l *AppLogger) Println(message string) {
	l.Print(message + "\n")
}

func (l *AppLogger) Trace(message string) {
	l.Println("TRACE | " + message)
}

func (l *AppLogger) Debug(message string) {
	if strings.HasPrefix(message, "[ExternalAssetHandler]") || strings.HasPrefix(message, "[AssetHandler]") {
		return
	}
	l.Println("DEBUG | " + message)
}

func (l *AppLogger) Info(message string) {
	l.Println("INFO  | " + message)
}

func (l *AppLogger) Warning(message string) {
	l.Println("WARN  | " + message)
}

func (l *AppLogger) Error(message string) {
	l.Println("ERROR | " + message)
}

func (l *AppLogger) Fatal(message string) {
	l.Println("FATAL | " + message)
	os.Exit(1)
}
