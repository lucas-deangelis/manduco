//go:build trace

package main

import (
	"fmt"
	"time"
)

var now = time.Now()

func since(msg string) {
	newNow := time.Now()
	fmt.Printf("%s: %s\n", msg, newNow.Sub(now))
	now = newNow
}
