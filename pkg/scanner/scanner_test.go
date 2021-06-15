package scanner

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

type mockScrollReader struct {
	idx    int
	scroll []rune
}

func (sr mockScrollReader) More() bool {
	return len(sr.scroll) > sr.idx
}

func (sr mockScrollReader) Peek() (rune, error) {
	if !sr.More() {
		return rune(0), errors.New("EOF")
	}
	return sr.scroll[sr.idx], nil
}

func (sr mockScrollReader) Read() (rune, error) {
	ru, e := sr.Peek()
	if e != nil {
		return rune(0), e
	}
	sr.idx++
	return ru, nil
}

func (sr mockScrollReader) PutBack(ru rune) error {
	last := sr.scroll[sr.idx:]
	first := append(sr.scroll[:sr.idx], ru)
	sr.scroll = append(first, last...)
	return nil
}

func TestScanAll(t *testing.T) {

	sr := mockScrollReader{
		scroll: []rune("1 + 2"),
	}

	exp := []string{"1", "+", "2"}
	act, e := ScanAll(sr)

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)

}
