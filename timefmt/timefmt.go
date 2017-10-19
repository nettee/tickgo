package timefmt

import (
	"time"
	"github.com/cactus/gostrftime"
)

const (
	fmtStr = "%Y-%m-%d %H:%M:%S (.%N)"
)

func Fmt(t time.Time) string {
	return gostrftime.Format(fmtStr, t)
}