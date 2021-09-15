package main

import (
	"github.com/dop251/goja"
	"github.com/g2a-com/klio-lifecycle/internal/build"
)

func main() {
	build.Run(func(script string, input map[string]interface{}) (output []string, err error) {
		vm := goja.New()

		vm.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))

		vm.Set("input", input)
		vm.Set("output", []interface{}{})

		_, err = vm.RunString(script)
		if err != nil {
			return output, err
		}

		err = vm.ExportTo(vm.Get("output"), &output)
		if err != nil {
			return output, err
		}

		return output, nil
	})
}
