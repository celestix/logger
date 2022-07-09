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
	l.output(s)
	l.output(string(colorReset))
}
