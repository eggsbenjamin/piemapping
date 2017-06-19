package commons

import (
	"fmt"
	"log"
	"os"
)

const (
	Ldate = 1 << iota
	Ltime
	Lmicroseconds
	Llongfile
)

type LevelledLogWriter interface {
	Info(...interface{})
	Infof(string, ...interface{})
	Error(...interface{})
	Errorf(string, ...interface{})
}

//	wrapper for logging functionality to enable uniform, prefixed levelled logs throughout the service
type Logger struct {
	info  *log.Logger
	error *log.Logger
}

//	constructor
func NewLogger(pref string, dt int) *Logger {
	iPref := fmt.Sprintf("[%s][Info] ", pref)
	ePref := fmt.Sprintf("[%s][Error] ", pref)
	return &Logger{
		info:  log.New(os.Stdout, iPref, dt),
		error: log.New(os.Stderr, ePref, dt),
	}
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Print(v...)
}

func (l *Logger) Infof(f string, v ...interface{}) {
	l.info.Printf(f, v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.error.Print(v...)
}

func (l *Logger) Errorf(f string, v ...interface{}) {
	l.error.Printf(f, v...)
}

//	noop logger
type NoopLogger struct{}

func (l *NoopLogger) Info(v ...interface{}) {}

func (l *NoopLogger) Infof(f string, v ...interface{}) {}

func (l *NoopLogger) Error(v ...interface{}) {}

func (l *NoopLogger) Errorf(f string, v ...interface{}) {}

//	gets the id associated with the different log date/time formats (https://golang.org/src/log/log.go?s=2383:2675#L27)
func GetLogLevelId(f string) int {
	lvls := map[string]int{
		"Ldate":         1,
		"Ltime":         2,
		"Lmicroseconds": 3,
		"Llongfile":     4,
		"Lshortfile":    5,
		"LUTC":          6,
		"LstdFlags":     7,
	}
	return lvls[f]
}
