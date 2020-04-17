```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	var t *time.Timer
	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		t.Reset(randomDuration())
	})
	time.Sleep(5 * time.Second)
}

func randomDuration() time.Duration {
    return time.Duration(rand.Int63n(1e9))
```
Above snippet when ran with go run -race filename produces

```
alitopaloglu~/go/src/github.com/AliTopal89 (master)$ go run concurrency/race_detector/race_detector.go
953.017725ms
1.038013796s
1.706020482s
1.943919132s
2.234044685s
2.788070264s
3.421402404s
3.758101296s
3.946055609s
4.428080639s
alitopaloglu~/go/src/github.com/AliTopal89 (master)$ go test concurrency/race_detector/race_detector.go -race?   	command-line-arguments	[no test files]
alitopaloglu~/go/src/github.com/AliTopal89 (master)$ go test --race concurrency/race_detector/race_detector.go
?   	command-line-arguments	[no test files]
alitopaloglu~/go/src/github.com/AliTopal89 (master)$ go run --race concurrency/race_detector/race_detector.go
953.2977ms
==================
WARNING: DATA RACE
Read at 0x00c000010018 by goroutine 7:
  main.main.func1()
      /Users/alitopaloglu/go/src/github.com/AliTopal89/concurrency/race_detector/race_detector.go:14 +0x123

Previous write at 0x00c000010018 by main goroutine:
  main.main()
      /Users/alitopaloglu/go/src/github.com/AliTopal89/concurrency/race_detector/race_detector.go:12 +0x18e

Goroutine 7 (running) created at:
  time.goFunc()
      /usr/local/Cellar/go/1.12.4/libexec/src/time/sleep.go:169 +0x51
==================

```

The race detector shows the problem: an unsynchronized read and write of the variable t from different goroutines. If the initial timer duration is very small, the timer function may fire before the main goroutine has assigned a value to t and so the call to t.Reset is made with a nil t.

```go
for time.Since(start) < 5*time.Second {
    <-reset
    t.Reset(randomDuration())
}
```

Here the main goroutine is wholly responsible for setting and resetting the Timer t and a new reset channel communicates the need to reset the timer in a thread-safe way.
