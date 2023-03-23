//go:build darwin

package system

import (
	"os/exec"
)

func ConfigureProxy(action string) error {
	if action == CMDStartProxy {
		// set setsocksfirewallproxy to localhost 1080
		err := exec.Command("networksetup", "-setsocksfirewallproxy", "Wi-Fi", "127.0.0.1", "1080").Run()
		if err != nil {
			return err
		}
		// enable setsocksfirewallproxystate
		err = exec.Command("networksetup", "-setsocksfirewallproxystate", "Wi-Fi", "on").Run()
		if err != nil {
			return err
		}
		return nil
	} else if action == CMDStopProxy {
		// disable setsocksfirewallproxystate
		err := exec.Command("networksetup", "-setsocksfirewallproxystate", "Wi-Fi", "off").Run()
		if err != nil {
			return err
		}
		return nil
	} else {
		return ErrInvalidAction
	}
}
