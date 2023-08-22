package types

import "net/http"

type ResponseWriter struct{}

var messages [][]byte

// Implementation of net/http.responseWriter interface
func (writer ResponseWriter) Header() http.Header {
	return http.Header{}
}

// Implementation of net/http.responseWriter interface
func (writer ResponseWriter) Write(message []byte) (int, error) {
	messages = append(messages, message)

	return 0, nil
}

// Implementation of net/http.responseWriter interface
func (writer ResponseWriter) WriteHeader(statusCode int) {}

func (writer *ResponseWriter) GetMessages() [][]byte {
	return messages
}
