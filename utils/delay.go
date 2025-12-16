package utils

import (
	"math/rand"
	"time"
)

// HumanDelay waiting time
func HumanDelay(min, max int) {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(max-min+1)+min) * time.Second)
}
