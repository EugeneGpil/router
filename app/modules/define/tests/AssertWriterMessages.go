package tests

import (
	"net/http"
	"net/url"

	"github.com/EugeneGpil/tester"

	httpTester "github.com/EugeneGpil/httpTester"
)

func AssertWriterMessages(messages [][]byte) {
	urlObj, err := url.Parse("/")
	tester.AssertNil(err)

	request := http.Request{
		Method: http.MethodGet,
		URL:    urlObj,
	}

	handler, _ := Mux.Handler(&request)

	writer := httpTester.GetTestResponseWriter()

	handler.ServeHTTP(writer, &request)

	writerMessages := writer.GetMessages()

	tester.AssertLen(messages, len(messages))

	for index, message := range writerMessages {
		tester.AssertSame(message, messages[index])
	}
}
