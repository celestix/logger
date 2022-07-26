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

// Level type is to define the logging levels, which would be used for saturating logs.
//
// Note: In case you create new levels, they must be added to the LevelToNum and LevelToColor maps.
type Level string

var (
	// LevelDebug is the DEBUG level which should be used for debugging.
	LevelDebug Level = "DEBUG"
	// LevelMain is the MAIN level which should be used for logging main arguments.
	LevelMain Level = "MAIN"
	// LevelInfo is the INFO level which should be used for logging all basic arguments.
	LevelInfo Level = "INFO"
	// LevelError is the ERROR level which should be used for logging all errors.
	LevelError Level = "ERROR"
	// LevelCritical is the CRITICAL level which should be used for logging all critical things.
	LevelCritical Level = "CRITICAL"

	// LevelToNum map converts the provided level to its integral value.
	//
	// One should add custom created levels to this map for proper functioning.
	LevelToNum = map[Level]int{
		LevelDebug:    -1,
		LevelInfo:     0,
		LevelMain:     1,
		LevelError:    2,
		LevelCritical: 3,
	}

	// LevelToColor map converts the provided level to its defined color.
	//
	// One should add custom created levels to this map for proper functioning.
	LevelToColor = map[Level]color{
		LevelDebug:    white,
		LevelInfo:     blue,
		LevelMain:     green,
		LevelError:    purple,
		LevelCritical: red,
	}
)
