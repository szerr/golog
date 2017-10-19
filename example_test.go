package golog_test

import (
	log "github.com/szerr/golog"
	"github.com/szerr/golog/writer"
)

func Example_Test() {
	log.Level = log.WarnLevel         //设置log等级为warning
	log.SetWriter(writer.DateLog("")) //这样会输出到当前目录下按日期分割的文件
	log.Info(111111)                  //这一条不会输出
	log.Warn(111111)
	log.Err(3333)
	log.Fat(4444)
}
