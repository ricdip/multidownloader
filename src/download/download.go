package download

import (
	"io"
	"net/http"
	"os"
	"path"
	"sync"

	"multidownloader/src/bar"
	"multidownloader/src/flags"

	zlog "github.com/rs/zerolog/log"
)

func downloadAux(link string, index int) (int64, error) {
	resp, err := http.Get(link)

	if err != nil {
		zlog.Error().Err(err).Msg("HTTP GET request")
		return -1, err
	}
	defer resp.Body.Close()

	file, err := os.Create(path.Base(resp.Request.URL.String()))
	if err != nil {
		zlog.Error().Err(err).Msg("file create")
		return -1, err
	}
	defer file.Close()

	var read int64
	if !flags.QuietLog {
		size := resp.ContentLength
		ch := make(chan struct{})

		go bar.ShowBar(ch, size, file, index)

		read, err = io.Copy(file, resp.Body)
		<-ch
	} else {
		read, err = io.Copy(file, resp.Body)
	}

	return read, err
}

func Download(waitGroup *sync.WaitGroup, link string, index int) {
	zlog.Trace().Caller().Str("function", "download.Download").Msg("start")
	defer waitGroup.Done()

	downloadAux(link, index)

	zlog.Trace().Caller().Str("function", "download.Download").Msg("exit")
}
