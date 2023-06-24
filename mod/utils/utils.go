package utils

import (
	"fmt"
	"net/url"
	"os"
	"path"

	printer "downloader/mod/printer"
)

func GetDownloadLinks() []string {
	args := os.Args[1:]

	for _, arg := range args {
		_, err := url.ParseRequestURI(arg)

		if err != nil {
			printer.FatalMsg("Error: illegal URL `%s`\n", arg)
		}
	}

	return args
}

func CheckArgs() {
	if len(os.Args) > 1 && (os.Args[1] == "help" || os.Args[1] == "--help" || os.Args[1] == "-h") {
		printer.HelpMsg(path.Base(os.Args[0]))
		os.Exit(0)
	}
}

func WaitUser() {
	printer.WaitUserMsg()
	fmt.Scanln()
}
