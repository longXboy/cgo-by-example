package main

import (
	"fmt"
	"runtime"
)

/*
#include "lib/lib.h"
#include <stdlib.h>       // for free()
#cgo LDFLAGS: -L. -llib
*/
import "C"

func main() {
	runtime.GOMAXPROCS(4)
	for i := 0; i <= 20; i++ {
		go func() {
			fmt.Println("sleep now")
			C.sleepwhile()
			fmt.Println("sleep after")
		}()
	}
	select {}
}
