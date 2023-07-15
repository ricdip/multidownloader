package utils

import (
	"fmt"
	"net/url"

	printer "downloader/src/printer"
)

func CheckDownloadLinks(links []string) error {
	if len(links) == 0 {
		return fmt.Errorf("no links found")
	}

	for _, arg := range links {
		_, err := url.ParseRequestURI(arg)

		if err != nil {
			return err
		}
	}

	return nil
}

func WaitUser() {
	printer.WaitUserMsg()
	fmt.Scanln()
}
