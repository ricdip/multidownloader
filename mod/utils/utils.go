package utils

import (
    "os"
    "path"
    "fmt"

    printer "downloader/mod/printer"
)


func GetDownloadLinks() []string {
    return os.Args[1:]
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
