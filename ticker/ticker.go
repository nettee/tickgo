package ticker

import (
	"time"
)

func Wait(duration time.Duration) {
	timer := time.NewTimer(duration)
	<- timer.C
}


