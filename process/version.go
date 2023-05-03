package process

import (
	"fmt"
	"runtime"
)

// Version view version of mkvm
func Version() {
	fmt.Print("mkvm version: 1.0.0 alpha\n")
	fmt.Println("go version: ", runtime.Version())
	return
}
