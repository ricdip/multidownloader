package printer

import (
	"fmt"
	"os"
)

func PrintDownloadLinks(links []string) {
	fmt.Println()
	fmt.Printf("%d links detected:\n", len(links))
	for i, v := range links {
		fmt.Printf("\t- %d: %s\n", i, v)
	}
	fmt.Println()
}

func HelpMsg(programName string) {
	fmt.Printf("This program allows you to download the given links simultaneously\n\n")

	fmt.Printf("Usage: %s link1 link2 ...\n", programName)
}

func FatalMsg(errStr string, args ...interface{}) {
	fmt.Printf(errStr, args...)
	os.Exit(1)
}

func WaitUserMsg() {
	fmt.Println("Press [ENTER] to start downloading or CTRL+C to exit")
}
