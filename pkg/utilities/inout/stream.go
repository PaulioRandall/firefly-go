package inout

import (
	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

// StreamItem is a function designed for processing a value from one stream
// before it goes into another. It accepts the previous, current, and next item
// in stream. The first return value is to be written out if the second boolean
// value is true, else it is removed from processsing.
//
// Note that a values marked for removal are scrubbed from history, that is,
// it will not become the previous value in the subsequent ProcessItem call, it
// will be the one before that.
type StreamItem[In, Out comparable] func(prev, curr, next In) (Out, error)

func Stream[In, Out comparable](
	r Reader[In],
	w Writer[Out],
	f StreamItem[In, Out],
) error {

	var (
		zeroIn, prev, curr, next In
		zeroOut                  Out
		e                        error
	)

	if next, e = readNext(r); e != nil {
		return ErrRead.Wrap(e, "Stream failed to read from reader")
	}

	for next != zeroIn {

		prev = curr
		curr = next
		next, e = readNext(r)

		if e != nil {
			return ErrRead.Wrap(e, "Stream failed to read from reader")
		}

		out, e := f(prev, curr, next)
		if e != nil {
			return ErrRead.Wrap(e, "Stream failed to stream item")
		}

		if out == zeroOut {
			curr = prev
			continue
		}

		if e = w.Write(out); e != nil {
			return ErrRead.Wrap(e, "Stream failed to write item to writer")
		}
	}

	return nil
}

func readNext[In comparable](r Reader[In]) (In, error) {
	var zeroIn In

	if !r.More() {
		return zeroIn, nil
	}

	in, e := r.Read()
	if err.Is(e, EOF) {
		return zeroIn, nil
	}

	if e != nil {
		return zeroIn, ErrRead.Wrap(e, "Failed to read from reader")
	}

	return in, nil
}
