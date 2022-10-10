// Package processor eases the processing of a stream of items
package inout

import (
	"errors"
	"fmt"
)

// ProcessItem is a function designed for processing a value from one stream
// before it goes into another. It accepts the previous, current, and next item
// in stream. The first return value is to be written out if the second boolean
// value is true, else it is removed from processsing.
//
// Note that a values marked for removal are scrubbed from history, that is,
// it will not become the previous value in the subsequent ProcessItem call, it
// will be the one before that.
type ProcessItem[In, Out comparable] func(prev, curr, next In) (Out, error)

func Process[In, Out comparable](
	r Reader[In],
	w Writer[Out],
	p ProcessItem[In, Out],
) error {

	var (
		zeroIn, prev, curr, next In
		zeroOut                  Out
		e                        error
	)

	if next, e = readNext(r); e != nil {
		return fmt.Errorf("[process.Process] Failed to read next value: %w", e)
	}

	for next != zeroIn {

		prev = curr
		curr = next
		next, e = readNext(r)

		if e != nil {
			return fmt.Errorf("[process.Process] Failed to read next value: %w", e)
		}

		out, e := p(prev, curr, next)
		if e != nil {
			return fmt.Errorf("[process.Process] Failed to process value: %w", e)
		}

		if out == zeroOut {
			curr = prev
			continue
		}

		if e = w.Write(out); e != nil {
			return fmt.Errorf("[process.Process] Failed to write value: %w", e)
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
	if errors.Is(e, EOF) {
		return zeroIn, nil
	}

	if e != nil {
		return zeroIn, e // TODO: wrap error
	}

	return in, nil
}
