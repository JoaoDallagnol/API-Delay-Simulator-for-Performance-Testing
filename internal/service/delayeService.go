package service

import (
	"errors"
	"math/rand"
	"time"
)

func FixedDelay(duration time.Duration) {
	time.Sleep(duration)
}

func RandomFailure(failureRate float64) bool {
	return rand.Float64() < failureRate
}

func CustomDelay(duration int) error {
	if duration < 0 {
		return ErrInvalidDelay
	}
	time.Sleep(time.Duration(duration) * time.Second)
	return nil
}

var ErrInvalidDelay = errors.New("invalid delay parameter")
