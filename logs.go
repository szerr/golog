package golog

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"sync"
)

var (
	Level       = 0
	logger      *log.Logger
	RunTimeLock = new(sync.Mutex)
	Wg          = new(sync.WaitGroup)
)

const (
	InfoLevel = iota
	WarnLevel
	ErrorLevel
	FatalLevel

	CallSkip = 2
)

//等待所有log输出
func Wait() {
	Wg.Wait()
}

func init() {
	SetWriter(os.Stdout)
}

func SetWriter(out io.Writer) {
	logger = log.New(out, "", log.Ldate|log.Ltime)
}

//公用的log输出函数
func Println(flag string, s ...interface{}) {
	RunTimeLock.Lock()
	_, file, line, ok := runtime.Caller(CallSkip)
	RunTimeLock.Unlock()
	Wg.Add(1)
	go func() {
		defer Wg.Done()
		if !ok {
			file = "???"
			line = 0
		}
		logger.Print(flag, " ", file+":", line, " ", fmt.Sprintln(s...))
	}()
}

func Info(s ...interface{}) {
	if Level <= InfoLevel {
		Println("Info", s...)
	}
}

func Warn(s ...interface{}) {
	if Level <= WarnLevel {
		Println("Warn", s...)
	}
}

func Err(s ...interface{}) {
	if Level <= ErrorLevel {
		Println("Error", s...)
	}
}

func Fat(s ...interface{}) {
	if Level <= FatalLevel {
		Println("Fatal", s...)
	}
}
