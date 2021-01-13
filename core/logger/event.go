package logger

// Subscribe on every log
func (t *Logger) Subscribe(handler func(text string, lvl Level)) {
	t.handlers = append(t.handlers, handler)
}

func (t *Logger) executeHandlers(text string, lvl Level) {
	for _, handler := range t.handlers {
		handler(text, lvl)
	}
}
