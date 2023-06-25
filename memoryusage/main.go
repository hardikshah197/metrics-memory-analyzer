package memoryprinter

import (
    "runtime"
    "fmt"
    "time"
)

func PrintSampleMemoryAllocation() {
    // Print our starting memory usage (should be around 0mb)
    PrintMemUsage()

    var overall [][]int
    for i := 0; i<4; i++ {

        // Allocate memory using make() and append to overall (so it doesn't get 
        // garbage collected). This is to create an ever increasing memory usage 
        // which we can track. We're just using []int as an example.
        a := make([]int, 0, 999999)
        overall = append(overall, a)

        // Print our memory usage at each interval
        PrintMemUsage()
        time.Sleep(time.Second)
    }

    // Clear our memory and print usage, unless the GC has run 'Alloc' will remain the same
    overall = nil
    PrintMemUsage()

    // Force GC to clear up, should see a memory drop
    runtime.GC()
    PrintMemUsage()
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number 
// of garage collection cycles completed.
func PrintMemUsage() runtime.MemStats {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    // For info on each, see: https://golang.org/pkg/runtime/#MemStats
    fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
    fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
    fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
    fmt.Printf("\tNumGC = %v\n", m.NumGC)

    return m
}

func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}


// Alloc uint64
// Alloc is bytes of allocated heap objects.
// "Allocated" heap objects include all reachable objects, as well as unreachable objects that the garbage collector has not yet freed.
// Specifically, Alloc increases as heap objects are allocated and decreases as the heap is swept and unreachable objects are freed.
// Sweeping occurs incrementally between GC cycles, so these two processes occur simultaneously, and as a result Alloc tends to change smoothly (in contrast with the sawtooth that is typical of stop-the-world garbage collectors).
// TotalAlloc uint64
// TotalAlloc is cumulative bytes allocated for heap objects.
// TotalAlloc increases as heap objects are allocated, but unlike Alloc and HeapAlloc, it does not decrease when objects are freed.
// Sys uint64
// Sys is the total bytes of memory obtained from the OS.
// Sys is the sum of the XSys fields below. Sys measures the virtual address space reserved by the Go runtime for the heap, stacks, and other internal data structures. It's likely that not all of the virtual address space is backed by physical memory at any given moment, though in general it all was at some point.
// NumGC uint32
// NumGC is the number of completed GC cycles.