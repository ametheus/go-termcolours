package main

import (
	"bufio"
	"flag"
	"github.com/thijzert/go-termcolours/timelog"
	"log"
	"os"
	"time"
)

var (
	referenceDuration = flag.Float64("duration", 1.0, "Reference duration for each line of output.")
)

func init() {
	flag.Parse()
	wr := timelog.NewWriter(os.Stdout)
	wr.ReferenceDuration = time.Duration(*referenceDuration * float64(time.Second))
	log.SetOutput(wr)
}

func main() {
	b := bufio.NewReader(os.Stdin)
	var l []byte
	var er error = nil
	for er == nil {
		l, er = b.ReadBytes('\n')
		if er == nil && len(l) > 0 {
			log.Printf("%s", l)
		}
	}
}
