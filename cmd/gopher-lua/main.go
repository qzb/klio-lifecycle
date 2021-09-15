package main

import (
	"fmt"

	"github.com/g2a-com/klio-lifecycle/internal/build"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func main() {
	build.Run(func(script string, input map[string]interface{}) (output []string, err error) {
		L := lua.NewState()

		L.SetGlobal("input", luar.New(L, input))
		L.SetGlobal("output", L.NewTable())

		defer L.Close()
		if err := L.DoString(script); err != nil {
			return output, err
		}

		table, ok := L.GetGlobal("output").(*lua.LTable)
		if !ok {
			return output, fmt.Errorf("output must be a table")
		}

		table.ForEach(func(_, val lua.LValue) {
			output = append(output, val.String())
		})

		return output, nil
	})
}
