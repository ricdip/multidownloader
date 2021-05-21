package printer

import (
    "fmt"
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
    fmt.Printf("Usage: %s link1 link2 ...\n", programName)
    fmt.Println("This program allows you to download the given links simultaneously.")
}

func ErrorMsg() {
    fmt.Println("Error: no links found in args")
}

func WaitUserMsg() {
    fmt.Println("Press [ENTER] to start downloading or CTRL+C to exit")
}
