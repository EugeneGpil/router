package DefineRoutes

import (
	"net/http"
	"net/url"

	"github.com/EugeneGpil/tester"
)

func AssertWriterMessages(messages [][]byte) {
	urlObj, err := url.Parse("/")
	tester.AssertNil(err)

	request := http.Request{
		Method: http.MethodGet,
		URL:    urlObj,
	}

	handler, _ := Mux.Handler(&request)

	writer := tester.GetTestResponseWriter()

	handler.ServeHTTP(writer, &request)

	writerMessages := writer.GetMessages()

	tester.AssertLen(messages, len(messages))

	for index, message := range writerMessages {
		tester.AssertSame(message, messages[index])
	}
}
