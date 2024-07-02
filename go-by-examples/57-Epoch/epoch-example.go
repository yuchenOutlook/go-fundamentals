package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	p := fmt.Println
	p(now)

	p(now.Unix())
	p(now.UnixMilli())
	p(now.UnixMicro())
	p(time.Unix(now.Unix(), 0))
	p(time.Unix(0, now.UnixNano()))
}