package files

import (
	"io"
	"net/http"
	"os"
	"sync/atomic"
)

type DownloadStatus struct {
	TotalSize  int64
	Downloaded int64
	Done       int32
	Failed     bool
}

func (ds *DownloadStatus) IsDone() bool {
	return atomic.LoadInt32(&ds.Done) == 1
}

func DownloadFile(url, dest string, ds *DownloadStatus) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	ds.TotalSize = resp.ContentLength

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	ds.Done = 1

	return nil
}
