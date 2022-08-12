// Logger-Go (github.com/anonyindian/logger)
// Copyright (C) 2022
// Veer Pratap Singh <hello@veer.codes>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser Public License for more details.
//
// You should have received a copy of the GNU Lesser Public License
// along with this program. If not, see [http://www.gnu.org/licenses/].

package logger

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

// Logger is the main struct of the library which implements methods for logging.
//
// Name (type string): Name of the new module you want to inherit into.
// Level (type Level): Current running level of the logger.
// NoColor (type bool): Enable it if you want to disable colors.
type Logger struct {
	Name        string
	Level       Level
	NoColor     bool
	out         *os.File
	mutex       *sync.Mutex
	color       color
	minima      int
	projectName string
}

// LoggerOpts contain all the optional fields of New Logger.
type LoggerOpts struct {
	ProjectName  string
	NoColor      bool
	MinimumLevel Level
}

// New function creates a new logger.
func New(out *os.File, opts *LoggerOpts) *Logger {
	if opts == nil {
		opts = &LoggerOpts{}
	}
	if opts.ProjectName == "" {
		opts.ProjectName = "PROGRAM"
	}
	return &Logger{
		Level:       LevelMain,
		NoColor:     opts.NoColor,
		out:         out,
		mutex:       &sync.Mutex{},
		minima:      LevelToNum[opts.MinimumLevel],
		projectName: opts.ProjectName,
	}
}

// Print method is used to print the provided arguments with custom settings.
func (l *Logger) Print(v ...any) {
	if !l.shouldDo() {
		return
	}
	if l.NoColor {
		l.print(fmt.Sprint(v...))
	} else {
		l.colorPrint(LevelToColor[l.Level], fmt.Sprint(v...))
	}
}

// Printf method is used to format the provided string with the arguments and print the result with custom settings.
func (l *Logger) Printf(format string, a ...any) {
	if !l.shouldDo() {
		return
	}
	if l.NoColor {
		l.print(fmt.Sprintf(format, a...))
	} else {
		l.colorPrint(LevelToColor[l.Level], fmt.Sprintf(format, a...))
	}
}

// Println method is used to print the provided arguments in a newline with custom settings.
func (l *Logger) Println(v ...any) {
	if !l.shouldDo() {
		return
	}
	if l.NoColor {
		l.print(fmt.Sprintln(v...))
	} else {
		l.colorPrint(LevelToColor[l.Level], fmt.Sprintln(v...))
	}
}

// Printlnf method is used to format the provided string with the arguments and print the result in a newline with custom settings.
func (l *Logger) Printlnf(format string, a ...any) {
	if !l.shouldDo() {
		return
	}
	if l.NoColor {
		l.print(fmt.Sprintf(format+"\n", a...))
	} else {
		l.colorPrint(LevelToColor[l.Level], fmt.Sprintf(format+"\n", a...))
	}
}

// ChangeLevel method should be used to change the level of current logger.
func (l *Logger) ChangeLevel(level Level) *Logger {
	l.Level = level
	return l
}

// ChangeMinimumLevel method should be used to change the minimum printable level of current logger.
func (l *Logger) ChangeMinimumLevel(minimumLevel Level) *Logger {
	l.minima = LevelToNum[minimumLevel]
	return l
}

// Create method is used when you want to initalise logger for a new module, parent logger will be inherited with a few mutations.
func (l *Logger) Create(name string) *Logger {
	l1 := &(*l)
	l1.Name = fmt.Sprintf("%s[%s]", l.Name, strings.ToUpper(name))
	return l1
}

func (l *Logger) shouldDo() bool {
	return LevelToNum[l.Level] > l.minima
}

func (l *Logger) rawNewline() string {
	if l.Name != "" {
		return fmt.Sprintf("[%s][%s]%s ", l.projectName, l.Level, l.Name)
	}
	return fmt.Sprintf("[%s][%s] ", l.projectName, l.Level)
}
