package main

import (
	"context"
	"encoding/json"

	"github.com/d5/tengo/v2"
	"github.com/g2a-com/klio-lifecycle/internal/build"
)

func main() {
	build.Run(func(scriptText string, input map[string]interface{}) (output []string, err error) {
		// create a new Script instance
		script := tengo.NewScript([]byte(scriptText))

		// set values
		err = script.Add("input", fixTypes(input))
		if err != nil {
			return output, err
		}
		err = script.Add("output", []interface{}{})
		if err != nil {
			return output, err
		}

		// run the script
		compiled, err := script.RunContext(context.Background())
		if err != nil {
			return output, err
		}

		// retrieve values
		for _, o := range compiled.Get("output").Array() {
			output = append(output, o.(string))
		}

		return output, nil
	})
}

func fixTypes(in interface{}) (out interface{}) {
	txt, _ := json.Marshal(in)
	json.Unmarshal(txt, &out)
	return out
}
