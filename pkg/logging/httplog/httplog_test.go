package httplog

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_httpLogWriter_getBytes(t *testing.T) {
	t.Run("inner bytes must be equal", func(t *testing.T) {
		lw := &httpLogWriter{}
		lw.bytes = []byte{1, 2, 3, 4, 5}
		bb := lw.getBytes()

		if len(bb) != len(lw.bytes) {
			t.Errorf("inner bytes len must be equal to result len, got %d, expected %d", len(bb), len(lw.bytes))
		}
	})
}

func Test_httpLogWriter_Write(t *testing.T) {
	t.Run("must write bytes into inner writer", func(t *testing.T) {
		recorder := &httptest.ResponseRecorder{Body: bytes.NewBuffer([]byte{})}
		lw := &httpLogWriter{recorder, 200, nil}

		lw.Write([]byte{1, 2, 3})

		if len(lw.bytes) == 0 {
			t.Errorf("len must be not 0, got %d", len(lw.bytes))
		}

		if len(recorder.Body.Bytes()) == 0 {
			t.Errorf("recorder len must be not 0, got %d", len(recorder.Body.Bytes()))
		}
	})
}

func Test_httpLogWriter_WriteHeader(t *testing.T) {
	t.Run("must write status code", func(t *testing.T) {
		recorder := &httptest.ResponseRecorder{}

		lw := httpLogWriter{recorder, 200, nil}

		lw.WriteHeader(http.StatusAccepted)

		if lw.statusCode != http.StatusAccepted {
			t.Errorf("expected status code %d, got %d", http.StatusAccepted, lw.statusCode)
		}

		if recorder.Code != http.StatusAccepted {
			t.Errorf("expected recorders.status code %d , got %d", http.StatusAccepted, recorder.Code)
		}
	})
}

func Test_getRequestBody(t *testing.T) {
	t.Run("both request body and result must be not empty", func(t *testing.T) {

		buf := bytes.NewBuffer([]byte{1, 2, 3, 4, 5})

		req, _ := http.NewRequest(http.MethodPost, "localhost", buf)

		bb := getRequestBody(req)

		if len(bb) == 0 {
			t.Error("expected result bytes len != 0")
		}

		bodyBytes, err := ioutil.ReadAll(req.Body)
		if err != nil {
			t.Errorf("expected ReadAll(req.Body) err == nil , got %d", err)
		}

		if len(bodyBytes) == 0 {
			t.Errorf("expected len(request.Body.Bytes) != 0")
		}

		if len(bodyBytes) != len(bb) {
			t.Errorf("body and result bytes is not equal, must be equal, got %d, %d", len(bodyBytes), len(bb))
		}
	})
}
