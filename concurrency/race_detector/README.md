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
1.038707616s
1.709814068s
1.94922919s
2.23786492s
2.790617615s
3.423992935s
3.758566823s
3.947171137s
4.4330058s
5.196108986s
5.462972639s
Found 1 data race(s)
exit status 66

```

The race detector shows the problem: an unsynchronized read and write of the variable t from different goroutines. If the initial timer duration is very small, the timer function may fire before the main goroutine has assigned a value to t and so the call to t.Reset is made with a nil t.

```go
for time.Since(start) < 5*time.Second {
    <-reset
    t.Reset(randomDuration())
}
```

Here the main goroutine is wholly responsible for setting and resetting the Timer t and a new reset channel communicates the need to reset the timer in a thread-safe way.

and when you run it again 

```
alitopaloglu~/go/src/github.com/AliTopal89 (master)$ go run --race concurrency/race_detector/race_detector.go
953.328706ms
1.038765614s
1.710507119s
1.947087978s
2.234883857s
2.78771125s
3.424274345s
3.760746756s
3.945846468s
4.427299413s
5.193237478s
5.459074635s
```
