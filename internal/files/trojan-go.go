package files

import (
	"TrojanUI/internal/paths"
	"context"
	"io"
	"net/http"
	"os"
	"runtime"
	"time"
)

var goos string

func init() {
	// Get the current OS
	switch runtime.GOOS {
	case "darwin":
		goos = "macOS"
	case "linux":
		goos = "Linux"
	case "windows":
		goos = "Windows.exe"
	}
}

func RequireExecutableUpdate() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	// Fetch the latest hash from the server
	req, _ := http.NewRequestWithContext(ctx, "GET", "https://huggingface.co/acheong08/files/raw/main/trojan/trojan-"+goos+".md5", nil)
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return false
	}
	defer response.Body.Close()
	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return false
	}
	latest_hash := string(body)
	// Get the current hash
	current_hash, err := os.ReadFile(paths.HashPath)
	if err != nil {
		return true
	}
	// Compare the hashes
	return string(current_hash) != latest_hash
}

func DownloadExecutable() (*DownloadStatus, error) {
	status := &DownloadStatus{}
	// Download the executable
	go func() {
		err := DownloadFile("https://huggingface.co/acheong08/files/resolve/main/trojan/trojan-"+goos, paths.ExecPath, status)
		if err != nil {
			status.Failed = true
			return
		}
		// Change the permissions of the executable
		err = os.Chmod(paths.ExecPath, 0755)
		if err != nil {
			status.Failed = true
			return
		}
		// Get the latest hash
		response, err := http.Get("https://huggingface.co/acheong08/files/raw/main/trojan/trojan-" + goos + ".md5")
		if err != nil {
			status.Failed = true
			return
		}
		defer response.Body.Close()
		// Read the response body
		body, err := io.ReadAll(response.Body)
		if err != nil {
			status.Failed = true
			return
		}
		latest_hash := string(body)
		// Write the latest hash to the hash file
		err = os.WriteFile(paths.HashPath, []byte(latest_hash), 0644)
		if err != nil {
			status.Failed = true
			return
		}
	}()
	return status, nil
}
