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
	"os"
)

var def = New(os.Stderr, &LoggerOpts{
	NoColor: true,
})

// Print function is used to print the provided arguments with default settings.
//
// Use New() function to create your own logger.
func Print(v ...any) {
	def.Print(v...)
}

// Println function is used to print the provided arguments in a newline with default settings.
//
// Use New() function to create your own logger.
func Println(v ...any) {
	def.Println(v...)
}

// Printf function is used to format the provided string with the arguments and print the result with default settings.
//
// Use New() function to create your own logger.
func Printf(format string, a ...any) {
	def.Printf(format, a...)
}

// Printlnf function is used to format the provided string with the arguments and print the result in a newline with default settings.
//
// Use New() function to create your own logger.
func Printlnf(format string, a ...any) {
	def.Printlnf(format, a...)
}
