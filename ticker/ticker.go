package ticker

import (
	"math"
	"time"
	"fmt"

	"github.com/nettee/tickgo/timefmt"
)

func Tick(nanosecond int) {
	toWait := int(math.Pow10(9)) - nanosecond
	timer := time.NewTimer(time.Nanosecond * time.Duration(toWait))
	<- timer.C

	ticker := time.NewTicker(time.Second)
	for t := range ticker.C {
		fmt.Println(timefmt.Fmt(t))
	}
}

