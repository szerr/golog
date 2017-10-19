package writer

import (
	"io"
	"os"
	"time"
)

//可以自定义时间间隔的writher
func IntervalsLog(path string, intervals time.Duration) io.Writer {
	var writer io.Writer = &intervalsLog{Path: path, intervals: intervals}
	return writer
}

type intervalsLog struct {
	intervals time.Duration
	time      time.Time
	fd        *os.File
	Path      string
}

func (d *intervalsLog) Write(p []byte) (n int, err error) {
	nowTime := time.Now()
	if d.time.Before(nowTime) {
		d.time = nowTime.Add(d.intervals)
		d.fd.Close()
		d.fd, err = os.OpenFile(d.Path+nowTime.Format("2006-01-02 15:04:05")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return 0, err
		}
	}
	return d.fd.Write(p)
}
