// Package logging contains primitive json logger
package logging

import (
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("must return not nil logger", func(t *testing.T) {
		tl := New(nil)

		if tl == nil {
			t.Error("expected not nil, got  nil")
		}
	})

	t.Run("must return not nil logger with non nil inner logger", func(t *testing.T) {
		tl := New(nil)

		if tl == nil {
			t.Error("expected not nil, got  nil")
		}

		if tl.logger == nil {
			t.Error("expected not nil, got nil")
		}
	})
}

func TestLogger_InfoWithFields(t *testing.T) {
	t.Run("must write data into inner writer", func(t *testing.T) {
		tWriter := &testWriter{}
		tl := New(tWriter)

		tl.InfoWithFields(Fields{
			"var": "test",
		})

		if len(tWriter.bb) == 0 {
			t.Error("expected some payload in logger, got len == 0")
		}
	})
}

type testWriter struct {
	bb []byte
}

func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.bb = p
	log.Println(string(tw.bb))
	return len(p), nil
}

func TestLogger_ErrorWithFields(t *testing.T) {
	t.Run("must write data into inner writer", func(t *testing.T) {
		tWriter := &testWriter{}
		tl := New(tWriter)

		tl.ErrorWithFields(Fields{
			"var": "test",
		})

		if len(tWriter.bb) == 0 {
			t.Error("expected some payload in logger, got len == 0")
		}
	})
}
