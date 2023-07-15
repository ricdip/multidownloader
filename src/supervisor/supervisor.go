package supervisor

import (
	"fmt"
	"sync"

	bar "downloader/src/bar"
	download "downloader/src/download"
	"downloader/src/flags"
	printer "downloader/src/printer"
	utils "downloader/src/utils"
)

var waitGroup sync.WaitGroup

func Exec() {
	printer.PrintDownloadLinks(*flags.Links)

	utils.WaitUser()

	bar.CreateBars(len(*flags.Links))
	for i, link := range *flags.Links {
		waitGroup.Add(1)
		go download.Download(&waitGroup, link, i)
	}
	waitGroup.Wait()

	fmt.Println()
	fmt.Println("Exiting")
}
