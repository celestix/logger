package logger

func (l *Logger) print(b string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.output(b)
}

func (l *Logger) output(b string) {
	l.out.Write([]byte(b))
}
