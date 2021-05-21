package bar

import (
    "time"
    "os"
    "fmt"

    "github.com/cheggaaa/pb"
)


var Bars []*pb.ProgressBar
var BarPool = &pb.Pool{}

func ShowBar(ch chan bool, size int64, file *os.File, index int) {
    Bars[index].SetTotal64(size)
    Bars[index].Set64(0)
    Bars[index].Prefix(fmt.Sprintf("link %d: ", index))
    Bars[index].SetUnits(pb.U_BYTES)

    var fileStat os.FileInfo
    for (fileStat == nil || fileStat.Size() < size) {
        fileStat, _ := file.Stat()
        Bars[index].Set64(fileStat.Size())
        time.Sleep(time.Millisecond)
        if fileStat.Size() == size {
            break;
        }
    }

    time.Sleep(time.Second)

    Bars[index].Finish()
    ch <- true
}

func CreateBars(n_bars int) {
    for i := 0; i < n_bars; i++ {
        bar := pb.StartNew(0)
        Bars = append(Bars, bar)
    }
    BarPool.Add(Bars...)
    BarPool.Start()
}
