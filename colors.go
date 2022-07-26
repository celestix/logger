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

type color string

var (
	colorReset color = "\033[0m"
	red        color = "\033[31m"
	green      color = "\033[32m"
	yellow     color = "\033[33m"
	blue       color = "\033[34m"
	purple     color = "\033[35m"
	cyan       color = "\033[36m"
	white      color = "\033[37m"
)

func (l *Logger) colorPrint(c color, s string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.output(string(c))
	l.output(l.rawNewline())
	l.output(s)
	l.output(string(colorReset))
}
