package bar

import (
	"fmt"
	"os"
	"time"

	"github.com/cheggaaa/pb"

	zlog "github.com/rs/zerolog/log"
)

var bars []*pb.ProgressBar
var barPool = &pb.Pool{}

func ShowBar(ch chan struct{}, size int64, file *os.File, index int) {
	zlog.Trace().Caller().Str("function", "bar.ShowBar").Msg("start")

	bars[index].SetTotal64(size)
	bars[index].Set64(0)
	bars[index].Prefix(fmt.Sprintf("link %d: ", index))
	bars[index].SetUnits(pb.U_BYTES)

	var fileStat os.FileInfo
	for fileStat == nil || fileStat.Size() < size {
		fileStat, _ := file.Stat()
		bars[index].Set64(fileStat.Size())
		time.Sleep(time.Millisecond)
		if fileStat.Size() == size {
			break
		}
	}

	time.Sleep(time.Second)

	bars[index].Finish()
	ch <- struct{}{}

	zlog.Trace().Caller().Str("function", "bar.ShowBar").Msg("exit")
}

func CreateBars(n_bars int) {
	zlog.Trace().Caller().Str("function", "bar.CreateBars").Msg("start")

	for i := 0; i < n_bars; i++ {
		bar := pb.StartNew(0)
		bars = append(bars, bar)
	}
	barPool.Add(bars...)
	barPool.Start()

	zlog.Trace().Caller().Str("function", "bar.CreateBars").Msg("exit")
}
