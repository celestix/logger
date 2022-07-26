// Logger-Go (github.com/anonyindian/logger-go)
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

type Logger struct {
	Name        string
	Level       level
	NoColor     bool
	out         *os.File
	mutex       *sync.Mutex
	color       color
	minima      int
	projectName string
}

type LoggerOpts struct {
	ProjectName  string
	NoColor      bool
	MinimumLevel level
}

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

func (l *Logger) Print(v ...any) {
	if !l.shouldDo() {
		return
	}
	if l.NoColor {
		l.print(fmt.Sprint(v...))
	} else {
		l.colorPrint(levelToColor[l.Level], fmt.Sprint(v...))
	}
}

func (l *Logger) Printf(format string, a ...any) {
	if !l.shouldDo() {
		return
	}
	if l.NoColor {
		l.print(fmt.Sprintf(format, a...))
	} else {
		l.colorPrint(levelToColor[l.Level], fmt.Sprintf(format, a...))
	}
}

func (l *Logger) Println(v ...any) {
	if !l.shouldDo() {
		return
	}
	if l.NoColor {
		l.print(fmt.Sprintln(v...))
	} else {
		l.colorPrint(levelToColor[l.Level], fmt.Sprintln(v...))
	}
}

func (l *Logger) Printlnf(format string, a ...any) {
	if !l.shouldDo() {
		return
	}
	if l.NoColor {
		l.print(fmt.Sprintf(format+"\n", a...))
	} else {
		l.colorPrint(levelToColor[l.Level], fmt.Sprintf(format, a...))
	}
}

func (l *Logger) ChangeLevel(level level) *Logger {
	l.Level = level
	return l
}

func (l *Logger) ChangeMinima(minimumLevel level) *Logger {
	l.minima = LevelToNum[minimumLevel]
	return l
}

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
