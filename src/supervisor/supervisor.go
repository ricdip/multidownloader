package supervisor

import (
	"fmt"
	"sync"

	"multidownloader/src/bar"
	"multidownloader/src/download"
	"multidownloader/src/flags"
	"multidownloader/src/log"

	zlog "github.com/rs/zerolog/log"
)

var waitGroup sync.WaitGroup

func Exec() {
	log.Init()

	zlog.Trace().Caller().Str("function", "supervisor.Exec").Msg("start")

	for i, v := range *flags.Links {
		zlog.Info().Str("link", v).Int("index", i).Msg("detected")
	}

	if !(flags.AutoYes || flags.QuietLog) {
		zlog.Info().Msg("Press [ENTER] to continue...")
		fmt.Scanln()
	}

	if !flags.QuietLog {
		bar.CreateBars(len(*flags.Links))
	}

	for i, link := range *flags.Links {
		waitGroup.Add(1)
		go download.Download(&waitGroup, link, i)
	}

	waitGroup.Wait()

	zlog.Trace().Caller().Str("function", "supervisor.Exec").Msg("exit")
}
