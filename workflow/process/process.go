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
type ProcessItem[T any] func(prev, curr, next T) (T, bool, error)

func Process[T any](
	r inout.Reader[T],
	w inout.Writer[T],
	p ProcessItem[T],
) error {

	var prev, curr, next T
	var more bool
	var e error

	if !r.More() {
		return nil
	}

	if next, more, e = readNext(r); e != nil {
		return fmt.Errorf("[dataproc.Stream] Failed to read next value: %w", e)
	}

	for more {

		prev = curr
		curr = next
		next, more, e = readNext(r)

		if e != nil {
			return fmt.Errorf("[dataproc.Stream] Failed to read next value: %w", e)
		}

		v, keep, e := p(prev, curr, next)
		if e != nil {
			return fmt.Errorf("[dataproc.Stream] Failed to process value: %w", e)
		}

		if !keep {
			curr = prev
			continue
		}

		if e = w.Write(v); e != nil {
			return fmt.Errorf("[dataproc.Stream] Failed to write value: %w", e)
		}
	}

	return nil
}

func readNext[T any](r inout.Reader[T]) (T, bool, error) {
	var zero T

	if !r.More() {
		return zero, false, nil
	}

	v, e := r.Read()
	if errors.Is(e, inout.EOF) {
		return zero, false, nil
	}

	if e != nil {
		return zero, false, e // TODO: wrap error
	}

	return v, true, nil
}
