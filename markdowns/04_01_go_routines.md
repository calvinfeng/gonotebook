# Go-routines
## Basics
Before diving into goroutines, readers should have completed the concurrency section of Go Tour and
have a high level understanding of channels. Channels are the means for goroutines to communicate 
with each other. The famous Golang mantra is the following.

> Don't communicate by sharing memory; share memory by communicating

So What is a goroutine? The common definition is *a lightweight thread of execution*. We can create 
a goroutine easily using the `go` keyword.
```golang
func main() {
    ch := make(chan []byte)
    go func(recv chan int) {
        resp, _ := http.Get("https://example.com")
        defer resp.Body.Close()
        jsonBytes, _ := ioutil.ReadAll(res.Body)
        recv <- jsonBytes
    }(ch)

    fmt.Println("Program continues...")
    // Do other things

    // Then wait for HTTP response to come back
    result := <-ch
    fmt.Println("Done, result is", string(result))
}
```

With goroutine, we can make asynchronous network IO and continue the program without the need to 
wait for the response to come back. Obviously, network IO is just one of the many examples. We can
do parallelized processing like a mini map reduce and fully utilize each CPU core we have on our 
modern days machine. A typical Intel processsor has 4 cores at least. The high end models even have
8 or 16 cores.

## Details 
If we read carefully, we should know that goroutine is not an OS thread. It is an abstraction of a 
thread! An actual OS thread is typically assigned with multiple goroutines. It is Go runtime's job
to manage how many OS threads to create and how goroutines are mapped to operating system threads.
So now let's deep dive into how does Go runtime scheduler works.

### G Struct
Each goroutine is described by a struct named `G` which can be found on [here][1]. The struct keeps 
track of various runtime information, like stack and status. The status indicates whether the 
goroutine is blocked, runnable, or running. 

### M Struct
TBW

### P Struct
TBW

### Context Switch
The cost of context switching is around 50~100 nanoseconds, quoting from some internet sources.

[1]: https://github.com/golang/go/blob/master/src/runtime/runtime2.go)