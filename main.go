// promhisto is a small tool that help you tweak your prometheus.ExponentialBuckets
// parameters.
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	start  = flag.Float64("s", 0.0, "start")
	factor = flag.Float64("f", 0.0, "factor")
	count  = flag.Int("c", 0, "count")

	asDuration = flag.Bool("d", false, "render as duration")
)

func main() {
	flag.Parse()

	if *start == 0 || *factor == 0 || *count == 0 {
		flag.Usage()
		os.Exit(1)
	}

	b := prometheus.ExponentialBuckets(*start, *factor, *count)
	if *asDuration {
		var d []time.Duration
		for _, i := range b {
			d = append(d, time.Duration(i*1e9)*time.Nanosecond)
		}
		fmt.Println(d)
	} else {
		fmt.Println(b)
	}
}