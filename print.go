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

func (l *Logger) print(b string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.output(l.rawNewline())
	l.output(b)
}

func (l *Logger) output(b string) {
	l.out.Write([]byte(b))
}
