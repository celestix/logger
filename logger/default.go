package logger

import (
	"fmt"
	"os"
)

var def = New(os.Stderr, "")

func Print(v ...any) {
	def.Print(v...)
}

func Println(v ...any) {
	def.Println(v...)
}

func Printf(format string, a ...any) {
	def.Printf(format, a...)
}

func (l *Logger) TestColor(c color, s string) {
	l.colorPrint(c, s)
}

func TestC(a ...any) {
	// red        color = "\033[31m"
	// green      color = "\033[32m"
	// yellow     color = "\033[33m"
	// blue       color = "\033[34m"
	// purple     color = "\033[35m"
	// cyan       color = "\033[36m"
	// white      color = "\033[37m"
	def.TestColor(red, fmt.Sprintln(a...))
	def.TestColor(green, fmt.Sprintln(a...))
	def.TestColor(yellow, fmt.Sprintln(a...))
	def.TestColor(blue, fmt.Sprintln(a...))
	def.TestColor(purple, fmt.Sprintln(a...))
	def.TestColor(cyan, fmt.Sprintln(a...))
	def.TestColor(white, fmt.Sprintln(a...))
	def.TestColor(colorReset, fmt.Sprintln(a...))
}
