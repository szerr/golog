package golog

import (
	"./writer"
	"testing"
	"time"
)

func TestInfo(t *testing.T) {
	Level = WarnLevel
	SetWriter(writer.DateLog(""))
	Info(111111)
	Warn(111111)
	Err(3333)
	Fat(4444)
}
func TestLog(t *testing.T) {
	Level = WarnLevel
	SetWriter(writer.IntervalsLog("", time.Second))
	for i := 3; i > 0; i -= 1 {
		Info(111111)
		Warn(111111)
		Err(3333)
		Fat(4444)
		time.Sleep(time.Second)
	}
}
