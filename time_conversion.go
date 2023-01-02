package lorem_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// BenchmarkTimeConversion-10    	 3117019	       363.0 ns/op
// BenchmarkTimeConversion-10    	 3226982	       366.1 ns/op
func BenchmarkTimeConversion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		timeConversion("12:01:00PM")
		timeConversion("12:01:00AM")
	}
}

func timeConversion(s string) string {
	time := strings.Split(s, ":")
	h, m, s := time[0], time[1], time[2]
	day := strings.HasSuffix(s, "AM")
	night := strings.HasSuffix(s, "PM")

	fh, _ := strconv.Atoi(h)
	if night && fh < 12 {
		fh += 12
	}

	if day && fh == 12 {
		fh = 0
	}

	return fmt.Sprintf("%02d:%s:%s", fh, m, s[:2])
}
