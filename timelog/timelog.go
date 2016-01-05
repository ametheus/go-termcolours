package timelog

import (
	"fmt"
	tc "github.com/ametheus/go-termcolours"
	"github.com/ametheus/go-termcolours/heatmap"
	"io"
	"time"
)

type Writer struct {
	ReferenceDuration time.Duration
	base              io.Writer
	now               time.Time
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{ReferenceDuration: 1 * time.Minute,
		base: w,
		now:  time.Now()}
}

func (w *Writer) Write(in []byte) (n int, err error) {
	t1 := time.Now()
	d := t1.Sub(w.now)
	w.now = t1

	if len(in) > 0 && in[len(in)-1] == '\n' {
		in = in[0 : len(in)-1]
	}

	perc := heatmap.Percentage(100.0 * float64(d) / float64(w.ReferenceDuration))
	dd := []byte(tc.Foreground24(perc, fmtDur(d)))

	line := make([]byte, len(in)+len(dd)+2)
	copy(line, dd)
	line[len(dd)] = '\t'
	copy(line[len(dd)+1:], in)
	line[len(in)+len(dd)+1] = '\n'

	return w.base.Write(line)
}

func fmtDur(d time.Duration) string {
	if d < time.Millisecond {
		return fmt.Sprintf("%.4g\u00b5s", float64(d)/float64(time.Microsecond))
	} else if d < time.Second {
		return fmt.Sprintf("%.4gms", float64(d)/float64(time.Millisecond))
	} else if d < (5 * time.Minute) {
		return fmt.Sprintf("%.4gs", float64(d)/float64(time.Second))
	} else if d < 2*time.Hour {
		return fmt.Sprintf("%.4gm", float64(d)/float64(time.Minute))
	} else if d < 72*time.Hour {
		return fmt.Sprintf("%.4gh", float64(d)/float64(time.Hour))
	} else {
		return fmt.Sprintf("%.4gd", float64(d)/(24.0*float64(time.Hour)))
	}
}
