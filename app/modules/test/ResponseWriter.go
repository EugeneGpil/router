package test

import "net/http"

type ResponseWriter struct {
	messages [][]byte
}

func (writer *ResponseWriter) Header() http.Header {
	return http.Header{}
}

func (writer *ResponseWriter) Write(message []byte) (int, error) {
	writer.messages = append(writer.messages, message)

	return 0, nil
}

func (writer *ResponseWriter) WriteHeader(statusCode int) {}

func (writer *ResponseWriter) GetMessages() [][]byte {
	return writer.messages
}
