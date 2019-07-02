package main

import (
	"fmt"
	"github.com/antonmedv/expr"
)

func main() {
	env := map[string]interface{}{
		"foo": 1,
		"bar": struct{ Value int }{1},
	}

	out, _ := expr.Eval("foo + bar.Value", env)
	fmt.Println(out)
}
