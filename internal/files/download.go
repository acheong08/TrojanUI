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

func (ds *DownloadStatus) Write(p []byte) (int, error) {
	n := len(p)
	atomic.AddInt64(&ds.Downloaded, int64(n))
	return n, nil
}

func (ds *DownloadStatus) Progress() float64 {
	return float64(ds.Downloaded) / float64(ds.TotalSize) * 100
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

	mw := io.MultiWriter(out, ds)
	_, err = io.Copy(mw, resp.Body)
	atomic.StoreInt32(&ds.Done, 1)

	return err
}
