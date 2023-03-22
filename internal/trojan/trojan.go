package trojan

import (
	"TrojanUI/internal/paths"
	"os"
	"os/exec"
	"sync"
)

type TrojanController struct {
	cmd  *exec.Cmd
	lock sync.Mutex
}

func New() *TrojanController {
	return &TrojanController{}
}

func (tc *TrojanController) Start() error {
	tc.lock.Lock()
	defer tc.lock.Unlock()

	if tc.cmd != nil {
		return ErrAlreadyRunning
	}

	tc.cmd = exec.Command(paths.ExecPath, "-config", paths.ConfigPath)
	err := tc.cmd.Start()
	if err != nil {
		tc.cmd = nil
		return err
	}

	return nil
}

func (tc *TrojanController) Stop() error {
	tc.lock.Lock()
	defer tc.lock.Unlock()

	if tc.cmd == nil {
		return ErrNotRunning
	}

	err := tc.cmd.Process.Signal(os.Interrupt)
	if err != nil {
		return err
	}

	_, err = tc.cmd.Process.Wait()
	tc.cmd = nil

	return err
}

func (tc *TrojanController) Check() error {
	_, err := os.Stat(paths.ConfigPath)
	if os.IsNotExist(err) {
		return ErrMissingConfigFile
	} else if err != nil {
		return err
	}

	_, err = os.Stat(paths.ExecPath)
	if os.IsNotExist(err) {
		return ErrMissingExecutable
	} else if err != nil {
		return err
	}

	return nil
}
