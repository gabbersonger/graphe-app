package logger

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"sync"

	xerrors "github.com/mdobak/go-xerrors"
)

const (
	reset = "\033[0m"

	black        = 30
	red          = 31
	green        = 32
	yellow       = 33
	blue         = 34
	magenta      = 35
	cyan         = 36
	lightGray    = 37
	darkGray     = 90
	lightRed     = 91
	lightGreen   = 92
	lightYellow  = 93
	lightBlue    = 94
	lightMagenta = 95
	lightCyan    = 96
	white        = 97
)

func colorize(colorCode int, v string) string {
	return fmt.Sprintf("\033[%sm%s%s", strconv.Itoa(colorCode), v, reset)
}

type Handler struct {
	directory string
	filename  string
	filepath  string
	h         slog.Handler
	b         *bytes.Buffer
	m         *sync.Mutex
}

func (h *Handler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.h.Enabled(ctx, level)
}

func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &Handler{h: h.h.WithAttrs(attrs), b: h.b, m: h.m}
}

func (h *Handler) WithGroup(name string) slog.Handler {
	return &Handler{h: h.h.WithGroup(name), b: h.b, m: h.m}
}

func (h *Handler) StdError(message string) {
	fmt.Printf("%s | %s\n", colorize(red, "ERROR"), message)
}

func (h *Handler) Print(message string) {
	os.MkdirAll(h.directory, os.ModePerm)
	f, err := os.OpenFile(h.filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		h.StdError("Could not open log file")
	}
	_, err = f.WriteString(message)
	if err != nil {
		h.StdError("Could not write to log file")
	}
	f.Close()
}

func (h *Handler) Handle(ctx context.Context, r slog.Record) error {
	switch r.Level {
	case slog.LevelDebug:
		h.Print(fmt.Sprintf("%s | %s\n", colorize(darkGray, "DEBUG"), r.Message))
	case slog.LevelInfo:
		h.Print(fmt.Sprintf("%s | %s\n", colorize(cyan, "INFO"), r.Message))
	case slog.LevelWarn:
		h.Print(fmt.Sprintf("%s | %s\n", colorize(yellow, "WARN"), r.Message))
	case slog.LevelError:
		h.StdError(r.Message)
		fmt.Println(xerrors.StackTrace(xerrors.New("")))
		h.Print(fmt.Sprintf("%s | %s\n", colorize(red, "ERROR"), r.Message))
		os.Exit(1)
	default:
		panic(fmt.Sprintf("Unknown log level: %v\n", r.Level))
	}
	return nil
}

const LOG_FILE_NAME = "graphe.log"

func NewGrapheLogger() *slog.Logger {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("No home directory")
	}
	log_directory := filepath.Join(home_dir, "/Library/Logs/Graphe")
	err = os.MkdirAll(log_directory, os.ModePerm)

	b := &bytes.Buffer{}
	handler := &Handler{
		directory: log_directory,
		filename:  LOG_FILE_NAME,
		filepath:  path.Join(log_directory, LOG_FILE_NAME),

		b: b,
		h: slog.NewTextHandler(b, &slog.HandlerOptions{Level: slog.LevelDebug}),
		m: &sync.Mutex{},
	}
	return slog.New(handler)
}
