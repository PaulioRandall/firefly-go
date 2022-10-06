package dataproc

import (
	"errors"
	"fmt"

	"github.com/PaulioRandall/firefly-go/workflow/inout"
)

// Process is a function designed for processing a value from one stream before
// it goes into another. It accepts the previous, current, and next item in
// stream. The first return value is to be written out if the second boolean
// value is true, else it is removed from processsing.
//
// Note that a values marked for removal are scrubbed from history, that is,
// it will not become the previous value in the subsequent Process call, it
// will be the one before that.
type Process[T any] func(prev, curr, next T) (T, bool, error)

func Stream[T any](r inout.Reader[T], w inout.Writer[T], p Process[T]) error {
	var prev, curr, next T
	var e error

	for r.More() {
		prev = curr
		curr = next
		next, e = readNext(r)

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

func readNext[T any](r inout.Reader[T]) (T, error) {
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
