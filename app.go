package main

import (
	"TrojanUI/internal/files"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx            context.Context
	downloadStatus *files.DownloadStatus
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) RequiresUpdate() bool {
	update_required := files.RequireExecutableUpdate()
	return update_required
}

func (a *App) DownloadVPN() string {
	var err error
	a.downloadStatus, err = files.DownloadExecutable()
	if err != nil {
		return "failed"
	}
	return "loading"
}

func (a *App) DownloadProgress() float64 {
	return a.downloadStatus.Progress()
}

func (a *App) CheckConfig() bool {
	config, err := files.ReadConfig()
	if err != nil {
		return false
	}
	if config == "" {
		return false
	}
	return true
}
