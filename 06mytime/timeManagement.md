# Time Management in Go

In Golang, memory management happens automatically in Go.
The allocation and de-allocation happens automatically.

The 2 main commands/keywords are:
    1. new() -> memory allocated but not init
    2. make() -> allocated and init

> zeroed storage is storage where you can not place data.

The Garbage Collection takes place automatically using the `GC` or `Garabage Collector`.
Anything that becomes useless/out of scope or Nil. It is taken care of using GC.

The information about all this can be found in the `runtime` package.
This package contains all the information about the system settings, envs,
and hardware.

The `GOGC` variable in `go env` sets the garbage collection percentage. By default it is 100, but
can be changed using runtime functions.

Another important aspect in Go, for handling memory is Pointers.