package handlers

import "net/http"

func ExampleHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello!"))
}
