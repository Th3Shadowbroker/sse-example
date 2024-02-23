package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var server *http.Server

func Start(port uint) {
	var mux = http.NewServeMux()
	mux.Handle("/sse", new(Handler))

	server = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}

func Stop() {
	if server != nil {
		if err := server.Shutdown(context.TODO()); err != nil {
			panic(err)
		}
	}
}

func AwaitReadiness(url string, retries uint, intervalSec uint) {
	for i := 0; i < int(retries); i++ {
		response, _ := http.Head(url)
		if response != nil && response.StatusCode == http.StatusOK {
			return
		}

		time.Sleep(time.Duration(intervalSec) * time.Second)
	}
	panic(errors.New("server didn't reach ready state in time"))
}
