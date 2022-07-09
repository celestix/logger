package logger

import (
	"fmt"
	"os"
	"sync"
)

type Logger struct {
	out   *os.File
	mutex *sync.Mutex
	sep   string
}

func (l *Logger) Print(v ...any) {
	l.print(fmt.Sprint(v...))
}

func (l *Logger) Printf(format string, a ...any) {
	l.print(fmt.Sprintf(format, a...))
}

func (l *Logger) Println(v ...any) {
	l.print(fmt.Sprintln(v...))
}

func (l *Logger) Printlnf(format string, a ...any) {
	l.print(fmt.Sprintf(format+"\n", a...))
}

func New(out *os.File, separator string) *Logger {
	if separator == "" {
		separator = " "
	}
	return &Logger{out: out, sep: separator, mutex: &sync.Mutex{}}
}
