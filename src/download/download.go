package download

import (
	"io"
	"net/http"
	"os"
	"path"
	"sync"

	"multidownloader/src/bar"

	zlog "github.com/rs/zerolog/log"
)

func downloadAux(link string, index int) (int64, error) {
	resp, err := http.Get(link)
	size := resp.ContentLength

	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	file, err := os.Create(path.Base(resp.Request.URL.String()))
	if err != nil {
		return -1, err
	}
	defer file.Close()

	ch := make(chan struct{})
	go bar.ShowBar(ch, size, file, index)

	read, err := io.Copy(file, resp.Body)

	<-ch
	return read, err
}

func Download(waitGroup *sync.WaitGroup, link string, index int) {
	zlog.Trace().Caller().Str("function", "download.Download").Msg("start")
	defer waitGroup.Done()

	downloadAux(link, index)

	zlog.Trace().Caller().Str("function", "download.Download").Msg("exit")
}
