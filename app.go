package main

import (
	"TrojanUI/internal/files"
	"TrojanUI/internal/trojan"
	"context"
)

// App struct
type App struct {
	ctx            context.Context
	downloadStatus *files.DownloadStatus
	vpnActive      bool
	trojanInstance *trojan.TrojanController
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

// / Initializations
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

func (a *App) DownloadStatus() float64 {
	if a.downloadStatus == nil {
		return 0
	}
	return a.downloadStatus.Progress()
}

/// VPN control

func (a *App) StartVPN() bool {
	if a.vpnActive {
		return true
	}
	if a.trojanInstance == nil {
		a.trojanInstance = trojan.New()
	}
	if a.trojanInstance.Check() != nil {
		return false
	}
	err := a.trojanInstance.Start()
	return err == nil
}

func (a *App) StopVPN() bool {
	if !a.vpnActive {
		return true
	}
	err := a.trojanInstance.Stop()
	return err == nil
}
