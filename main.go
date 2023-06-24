package main

import (
	"fmt"
	"sync"

	bar "downloader/mod/bar"
	download "downloader/mod/download"
	printer "downloader/mod/printer"
	utils "downloader/mod/utils"
)

var waitGroup sync.WaitGroup

func main() {
	utils.CheckArgs()

	links := utils.GetDownloadLinks()
	if len(links) == 0 {
		printer.FatalMsg("Error: no links found in args\n")
	}

	printer.PrintDownloadLinks(links)

	utils.WaitUser()

	bar.CreateBars(len(links))
	for i, link := range links {
		waitGroup.Add(1)
		go download.Download(&waitGroup, link, i)
	}
	waitGroup.Wait()

	fmt.Println()
	fmt.Println("Exiting")
}
