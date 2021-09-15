package exec

import (
	"bufio"
	"bytes"

	log "github.com/g2a-com/klio-logger-go"
)

type writer struct {
	Logger *log.Logger
	Text   *string
}

func (w *writer) Write(data []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		text := scanner.Text()
		*w.Text += text
		w.Logger.Print(text)
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return len(data), nil
}
