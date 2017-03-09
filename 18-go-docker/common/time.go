package common

import "time"

func DelaySecond(n time.Duration) {
	time.Sleep(n * time.Second) // <------------ here
}
