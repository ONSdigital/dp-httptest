package httpmocks

import (
	"io"
	"net/http"
)

// ReadCloserMock is a mock implementation of io.ReadCloser
type ReadCloserMock struct {
	GetEntityFunc func() ([]byte, error)
	done          bool
	IsClosed      bool
}

// Fulfill the io.Reader interface.
func (mock *ReadCloserMock) Read(p []byte) (n int, err error) {
	if mock.done {
		return 0, io.EOF
	}

	b, err := mock.GetEntityFunc()
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

// NewReadCloserMock construct a new ReadCloserMock.
func NewReadCloserMock(getEntityFunc func() ([]byte, error)) *ReadCloserMock {
	return &ReadCloserMock{
		GetEntityFunc: getEntityFunc,
		done:          false,
		IsClosed:      false,
	}
}

// NewResponseMock construct a new *http.Response with the provided body and status code.
func NewResponseMock(body *ReadCloserMock, status int) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       body,
	}
}
