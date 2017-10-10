package header

import (
	"io"
	"os"
	"time"
)

func DateLog(path string) io.Writer {
	var writer io.Writer = &dateLog{Path: path}
	return writer
}

type dateLog struct {
	day  int
	fd   *os.File
	Path string
}

func (d *dateLog) Write(p []byte) (n int, err error) {
	nowDay := time.Now().Day()
	if d.day != nowDay {
		d.day = nowDay
		d.fd.Close()
		d.fd, err = os.OpenFile(d.Path+time.Now().Format("2006-01-02")+".log", os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			return 0, err
		}
	}
	return d.fd.Write(p)
}
