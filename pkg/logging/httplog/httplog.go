package httplog

import (
	"binarytree/pkg/logging"
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func New(logger LoggerWithFields) *Logger {
	return &Logger{
		logger: logger,
	}
}

type Logger struct {
	logger LoggerWithFields
}

type LoggerWithFields interface {
	ErrorWithFields(logging.Fields)
	InfoWithFields(logging.Fields)
}

func (l *Logger) LogHTTP(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		logWriter := httpLogWriter{w, 200, nil}

		reqBody := getRequestBody(r)
		next(&logWriter, r, ps)

		logWriteFunc := l.logger.InfoWithFields

		if logWriter.statusCode > 299 {
			logWriteFunc = l.logger.ErrorWithFields
		}

		logWriteFunc(logging.Fields{
			"path":          r.URL.Path,
			"headers":       r.Header,
			"request_body":  string(reqBody),
			"response_body": string(logWriter.getBytes()),
			"status_code":   logWriter.statusCode,
			"query_args":    r.URL.Query(),
		})

	}
}

func getRequestBody(r *http.Request) []byte {

	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return bodyBytes
}

type httpLogWriter struct {
	http.ResponseWriter
	statusCode int
	bytes      []byte
}

func (lw *httpLogWriter) WriteHeader(status int) {
	lw.statusCode = status
	lw.ResponseWriter.WriteHeader(status)
}

func (lw *httpLogWriter) Write(p []byte) (int, error) {
	lw.bytes = make([]byte, len(p))
	for i := range p {
		lw.bytes[i] = p[i]
	}
	return lw.ResponseWriter.Write(p)
}

func (lw *httpLogWriter) getBytes() []byte {
	return lw.bytes
}
