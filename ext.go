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
)

func (l *Logger) Mainln(v ...any) {
	if LevelToNum[LevelMain] < l.minima {
		return
	}
	if l.NoColor {
		l.print(fmt.Sprintln(v...))
	} else {
		l.colorPrint(LevelToColor[LevelMain], fmt.Sprintln(v...))
	}
}

func (l *Logger) Mainlnf(format string, a ...any) {
	if LevelToNum[LevelMain] < l.minima {
		return
	}
	if l.NoColor {
		l.print(fmt.Sprintf(format+"\n", a...))
	} else {
		l.colorPrint(LevelToColor[LevelMain], fmt.Sprintf(format+"\n", a...))
	}
}

func (l *Logger) Errorln(v ...any) {
	if LevelToNum[LevelError] < l.minima {
		return
	}
	if l.NoColor {
		l.print(fmt.Sprintln(v...))
	} else {
		l.colorPrint(LevelToColor[LevelError], fmt.Sprintln(v...))
	}
}

func (l *Logger) Errorlnf(format string, a ...any) {
	if LevelToNum[LevelError] < l.minima {
		return
	}
	if l.NoColor {
		l.print(fmt.Sprintf(format+"\n", a...))
	} else {
		l.colorPrint(LevelToColor[LevelError], fmt.Sprintf(format+"\n", a...))
	}
}

func (l *Logger) Debugln(v ...any) {
	if LevelToNum[LevelDebug] < l.minima {
		return
	}
	if l.NoColor {
		l.print(fmt.Sprintln(v...))
	} else {
		l.colorPrint(LevelToColor[LevelDebug], fmt.Sprintln(v...))
	}
}

func (l *Logger) Debuglnf(format string, a ...any) {
	if LevelToNum[LevelDebug] < l.minima {
		return
	}
	if l.NoColor {
		l.print(fmt.Sprintf(format+"\n", a...))
	} else {
		l.colorPrint(LevelToColor[LevelDebug], fmt.Sprintf(format+"\n", a...))
	}
}

func (l *Logger) Infoln(v ...any) {
	if LevelToNum[LevelInfo] < l.minima {
		return
	}
	if l.NoColor {
		l.print(fmt.Sprintln(v...))
	} else {
		l.colorPrint(LevelToColor[LevelInfo], fmt.Sprintln(v...))
	}
}

func (l *Logger) Infolnf(format string, a ...any) {
	if LevelToNum[LevelInfo] < l.minima {
		return
	}
	if l.NoColor {
		l.print(fmt.Sprintf(format+"\n", a...))
	} else {
		l.colorPrint(LevelToColor[LevelInfo], fmt.Sprintf(format+"\n", a...))
	}
}

func (l *Logger) Criticalln(v ...any) {
	if LevelToNum[LevelCritical] < l.minima {
		return
	}
	if l.NoColor {
		l.print(fmt.Sprintln(v...))
	} else {
		l.colorPrint(LevelToColor[LevelCritical], fmt.Sprintln(v...))
	}
}

func (l *Logger) Criticallnf(format string, a ...any) {
	if LevelToNum[LevelCritical] < l.minima {
		return
	}
	if l.NoColor {
		l.print(fmt.Sprintf(format+"\n", a...))
	} else {
		l.colorPrint(LevelToColor[LevelCritical], fmt.Sprintf(format+"\n", a...))
	}
}

func (l *Logger) Fatalln(v ...any) {
	if LevelToNum[LevelDebug] < l.minima {
		os.Exit(1)
		return
	}
	if l.NoColor {
		l.print(fmt.Sprintln(v...))
	} else {
		l.colorPrint(LevelToColor[LevelDebug], fmt.Sprintln(v...))
	}
	os.Exit(1)
}

func (l *Logger) Fatallnf(format string, a ...any) {
	if LevelToNum[LevelCritical] < l.minima {
		os.Exit(1)
		return
	}
	if l.NoColor {
		l.print(fmt.Sprintf(format+"\n", a...))
	} else {
		l.colorPrint(LevelToColor[LevelCritical], fmt.Sprintf(format+"\n", a...))
	}
	os.Exit(1)
}
