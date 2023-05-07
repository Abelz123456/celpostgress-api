package utils

import (
	"fmt"
	"runtime"
)

func RuntimeMaxProcessors() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}
