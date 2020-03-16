package logging

import (
	"io"
	"fmt"
	"log"
    "runtime"
)

type LoggingLevel int

const (
	INFO LoggingLevel = iota + 1
	DEBUG
	WARNING
	FATAL
)


var (
    BasicLogLevel      =  DEBUG
)

type Logging struct {
	lg *log.Logger
	level LoggingLevel
}

func New(writer io.Writer, format string, flags int) Logging{
	log := Logging{
		lg: log.New(writer, format, flags),
		level: WARNING,
	}
	return log
}

func (l *Logging) SetLevel(lv LoggingLevel) {
	l.level = lv
}
func (l *Logging) SetBasicLevel(lv LoggingLevel) {
    BasicLogLevel = lv
}

func (l *Logging) OutputLog(lv LoggingLevel, calldepth int, s string) error {
    if lv >= l.level || lv >= BasicLogLevel {
    	/*
        if (l.flag & Llevel) != 0 {
            return l.Output(calldepth, fmt.Sprintf("[%v] %s", lv, s))
        }
        */
        return l.lg.Output(calldepth, s)
    }
    return nil
}

func color_str(s string, color int, ending string) string{
    if runtime.GOOS == "windows" {
        return s+ending
    }
    return fmt.Sprintf("\x1b[%dm%s\x1b[0m%s", color, s, ending)
}

func (l *Logging) Info(format interface{}, v ...interface{}) {
    l.OutputLog(
    	INFO, 3,
    	fmt.Sprintf(color_str("[INFO]", 34, "  ")+fmt.Sprint(format), v...),
    )
}

func (l *Logging) Debug(format interface{}, v ...interface{}) {
    l.OutputLog(
    	DEBUG, 3,
    	fmt.Sprintf(color_str("[DEBUG]", 32, " ")+fmt.Sprint(format), v...),
    )
}

func (l *Logging) Warning(format interface{}, v ...interface{}) {
    l.OutputLog(
    	WARNING, 3,
    	fmt.Sprintf(color_str("[WARN]", 33, "  ")+fmt.Sprint(format), v...),
    )
}

func (l *Logging) Fatal(format interface{}, v ...interface{}) {
    l.OutputLog(
    	FATAL, 3,
    	fmt.Sprintf(color_str("[FATAL]", 31, " ")+fmt.Sprint(format), v...),
    )
}