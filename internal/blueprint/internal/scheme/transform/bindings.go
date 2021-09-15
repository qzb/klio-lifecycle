//go:generate ../../../../../scripts/generate-transform-module.sh
package transform

import (
	"github.com/dop251/goja"
)

var ToInternal func(interface{}) interface{}

func init() {
	vm := goja.New()
	_, err := vm.RunString(SCRIPT)
	if err != nil {
		panic(err)
	}

	err = vm.ExportTo(vm.Get("toInternal"), &ToInternal)
	if err != nil {
		panic(err)
	}
}
