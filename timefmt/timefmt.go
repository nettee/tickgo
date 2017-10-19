package timefmt

import (
	"time"
	"github.com/cactus/gostrftime"
	//"math"
	//"fmt"
)

const (
	fmtStrNano = "%Y-%m-%d %H:%M:%S (+.%N)"
	fmtStr = "%Y-%m-%d %H:%M:%S"
)

func Fmt(t time.Time) string {
	return gostrftime.Format(fmtStr, t)
}

func FmtNano(t time.Time) string {
	return gostrftime.Format(fmtStrNano, t)
}