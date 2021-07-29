package main

import (
	"os"
	"time"

	"github.com/AliTopal89/clockface" // REPLACE THIS!
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
