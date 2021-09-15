package main

import (
	"github.com/g2a-com/klio-lifecycle/internal/build"
)

func main() {
	build.Run(func(script string, input map[string]interface{}) (output []string, err error) {
		return output, nil
	})
}
