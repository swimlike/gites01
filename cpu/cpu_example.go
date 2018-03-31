// Copyright (c) 2010 Joseph D Poirier
// Distributable under the terms of The New BSD License
// that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"
	"runtime"
	"github.com/jpoirier/cpu"
)

func main() {
	fmt.Println("\nAn OS views physical cores and hyper-threads as logical processors \nin a multi-package, multi-core, multi-threading environment.")
	fmt.Println("Note that the term package refers to a physical processor\nand system refers to multiple packages.\n")

	fmt.Printf("              ----------- %s -----------\n", cpu.PackageVersion)

	fmt.Println("physical processors (aka packages)                        : ", cpu.Processors)
	fmt.Println("on line logical processors in the system                  : ", cpu.OnlnProcs) // also via cpu.OnlineProcs()
	fmt.Println("maximum logical processors in the system                  : ", cpu.MaxProcs) // also via cpu.ConfProcs()

	fmt.Println("cpuid present                                             : ", cpu.CpuidPresent)
//	fmt.Println("cpuid restricted                                          : ", cpu.CpuidRestricted)
	fmt.Println("hardware-threading supported                              : ", cpu.HardwareThreading)
	fmt.Println("processor L2 cache line size (bytes)                      : ", cpu.ProcessorL2CacheLine)
	fmt.Println("processor L2 cache size (bytes)                           : ", cpu.ProcessorL2Cache)
	fmt.Println("processor family                                          : ", cpu.ProcessorFamily)
	fmt.Println("vendor name                                               : ", cpu.Vendor)

	fmt.Println("\n    --- processor hardware capability  ---")
	fmt.Println("logical processors per physical processor                 : ", cpu.LogicalProcsPkg)
	fmt.Println("physical cores per physical processor                     : ", cpu.PhysicalCoresPkg)
	fmt.Println("hyper-threading enabled                                   : ", cpu.HyperThreadingEnabled)
	fmt.Println("hyper-threading logical processors per physical processor : ", cpu.HyperThreadingProcsPkg)

	// show the exported functions for completeness
	fmt.Println("")
	fmt.Println("on line logical processors in the system                  : ", cpu.OnlineProcs())
	fmt.Println("maximum logical processors in the system                  : ", cpu.ConfProcs())

	// set Go's runtime processor count
	runtime.GOMAXPROCS(int(cpu.MaxProcs))
}
