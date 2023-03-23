//go:build linux

package system

import (
	"errors"
	"os"
	"os/exec"
)

var (
	ErrGnomeNotFound = errors.New("gnome not found")
	ErrInvalidAction = errors.New("invalid action")
	CMDStartProxy    = "start"
	CMDStopProxy     = "stop"
)

// ConfigureGnomeProxy configures the GNOME proxy settings
func ConfigureGnomeProxy(action string) error {
	if action == CMDStartProxy {
		// Detect if GNOME is running
		if !detectGnome() {
			return ErrGnomeNotFound
		}
		// Set the proxy settings
		err := exec.Command("gsettings", "set", "org.gnome.system.proxy", "mode", "manual").Run()
		if err != nil {
			return err
		}
		err = exec.Command("gsettings", "set", "org.gnome.system.proxy.http", "host", "127.0.0.1").Run()
		if err != nil {
			return err
		}
		err = exec.Command("gsettings", "set", "org.gnome.system.proxy.http", "port", "1080").Run()
		if err != nil {
			return err
		}
		return nil
	} else if action == CMDStopProxy {
		// Detect if GNOME is running
		if !detectGnome() {
			return ErrGnomeNotFound
		}
		// Set the proxy settings
		err := exec.Command("gsettings", "set", "org.gnome.system.proxy", "mode", "none").Run()
		if err != nil {
			return err
		}
		return nil
	} else {
		return ErrInvalidAction
	}
}

func detectGnome() bool {
	_, err := os.Stat("/usr/bin/gsettings")
	return err == nil
}
