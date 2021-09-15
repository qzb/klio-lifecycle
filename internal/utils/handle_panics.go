package utils

import (
	"fmt"
	"os"
	"strings"

	log "github.com/g2a-com/klio-logger-go"
)

func HandlePanics() {
	if err := recover(); err != nil {
		msg := fmt.Sprint(err)
		log.ErrorLogger().WithLevel(log.FatalLevel).Print(strings.ToUpper(string(msg[0])) + string(msg[1:]))
		os.Exit(1)
	}
}
