package httpmocks

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/ONSdigital/log.go/log"
)

// ReadCloserMock is a mock implementation of io.ReadCloser
type ReadCloserMock struct {
	DoReadFunc func() ([]byte, error)
	done       bool
	IsClosed   bool
}

// Fulfill the io.Reader interface.
func (mock *ReadCloserMock) Read(p []byte) (n int, err error) {
	if mock.done {
		return 0, io.EOF
	}

	b, err := mock.DoReadFunc()
	if err != nil {
		return 0, err
	}

	for i, b := range b {
		p[i] = b
	}
	mock.done = true
	return len(b), nil
}

// Close fulfill the io.Closer interface
func (mock *ReadCloserMock) Close() error {
	mock.IsClosed = true
	return nil
}

// NewReadCloserMock construct a new ReadCloserMock b & err are the values to return when body.Read is called.
func NewReadCloserMock(b []byte, err error) *ReadCloserMock {
	return &ReadCloserMock{
		done:     false,
		IsClosed: false,
		DoReadFunc: func() ([]byte, error) {
			return b, err
		},
	}
}

// NewResponseMock construct a new *http.Response with the provided body and status code.
func NewResponseMock(body *ReadCloserMock, status int) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       body,
	}
}

// GetEntityBytes json marshall the value and return the produced []byte. Invokes t.FailNow() if there is an error while marshalling the value.
func GetEntityBytes(t *testing.T, i interface{}) []byte {
	body, err := json.Marshal(i)
	if err != nil {
		log.Event(nil, "failed to json marshal value", log.Error(err))
		t.FailNow()
	}
	return body
}
