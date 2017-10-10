package golog

import (
	"./header"
	"testing"
)

func TestInfo(t *testing.T) {
	Level = WarnLevel
	SetWriter(header.DateLog(""))
	Info(111111)
	Warn(111111)
	Err(3333)
	Fat(4444)
}
