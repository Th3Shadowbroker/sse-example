package server

import (
	"fmt"
	"net/http"
)

func Start(port uint16) {
	http.Handle("/sse", new(Handler))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
