package sync

import "sync"

type Counter struct {
	// mutex contains filtered or unexported fields
	sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

func main() {

}
