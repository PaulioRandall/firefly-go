// Package process eases the processing of a stream of items
package process

import (
	"errors"
	"fmt"

	"github.com/PaulioRandall/firefly-go/workflow/inout"
)

// ProcessItem is a function designed for processing a value from one stream
// before it goes into another. It accepts the previous, current, and next item
// in stream. The first return value is to be written out if the second boolean
// value is true, else it is removed from processsing.
//
// Note that a values marked for removal are scrubbed from history, that is,
// it will not become the previous value in the subsequent ProcessItem call, it
// will be the one before that.
type ProcessItem[T comparable] func(prev, curr, next T) (T, error)

func Process[T comparable](
	r inout.Reader[T],
	w inout.Writer[T],
	p ProcessItem[T],
) error {

	var zero, prev, curr, next T
	var e error

	if next, e = readNext(r); e != nil {
		return fmt.Errorf("[dataproc.Stream] Failed to read next value: %w", e)
	}

	for next != zero {

		prev = curr
		curr = next
		next, e = readNext(r)

		if e != nil {
			return fmt.Errorf("[dataproc.Stream] Failed to read next value: %w", e)
		}

		v, e := p(prev, curr, next)
		if e != nil {
			return fmt.Errorf("[dataproc.Stream] Failed to process value: %w", e)
		}

		if v == zero {
			curr = prev
			continue
		}

		if e = w.Write(v); e != nil {
			return fmt.Errorf("[dataproc.Stream] Failed to write value: %w", e)
		}
	}

	return nil
}

func readNext[T comparable](r inout.Reader[T]) (T, error) {
	var zero T

	if !r.More() {
		return zero, nil
	}

	v, e := r.Read()
	if errors.Is(e, inout.EOF) {
		return zero, nil
	}

	if e != nil {
		return zero, e // TODO: wrap error
	}

	return v, nil
}
