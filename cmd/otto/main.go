package main

import (
	"fmt"

	"github.com/g2a-com/klio-lifecycle/internal/build"
	"github.com/robertkrimen/otto"
)

func main() {
	build.Run(func(script string, input map[string]interface{}) (output []string, err error) {

		vm := otto.New()

		vm.Set("input", input)

		_, err = vm.Run("var output=[];\n" + script)
		if err != nil {
			return output, err
		}

		if outputVal, err := vm.Get("output"); err == nil {
			if lengthVal, err := outputVal.Object().Get("length"); err == nil {
				if length, err := lengthVal.ToInteger(); err == nil {
					for i := int64(0); i < length; i++ {
						if itemVal, err := outputVal.Object().Get(fmt.Sprint(i)); err == nil {
							if item, err := itemVal.ToString(); err == nil {
								output = append(output, item)
							}
						}
					}
				}
			}
		}

		return output, nil
	})
}
