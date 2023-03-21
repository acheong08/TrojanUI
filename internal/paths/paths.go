package paths

import (
	"os"
	"os/user"
	"path/filepath"
)

var BaseDir string
var ConfigPath string
var ExecPath string
var HashPath string

func init() {
	user, _ := user.Current()
	BaseDir = filepath.Join(user.HomeDir, ".config/trojan")
	// Make sure the base directory exists
	os.MkdirAll(BaseDir, 0755)
	ConfigPath = filepath.Join(BaseDir, "config.json")
	ExecPath = filepath.Join(BaseDir, "trojan-go")
	HashPath = filepath.Join(BaseDir, "trojan-go.md5")
}
