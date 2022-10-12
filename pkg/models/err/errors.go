package err

import (
	"errors"
)

func Is(a, b error) bool {
	return errors.Is(a, b)
}

func Unwrap(e error) error {
	return errors.Unwrap(e)
}
