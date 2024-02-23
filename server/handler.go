package server

import (
	"log"
	"net/http"
	"time"

	"github.com/th3shadowbroker/sse-example/messaging"
)

type Handler struct{}

func (h *Handler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {

	case http.MethodHead:
		h.writeStatus(writer, http.StatusOK)

	case http.MethodGet:
		writer.Header().Add("Content-Type", "application/stream+json")
		for {
			select {

			case <-req.Context().Done():
				return

			default:
				var message = messaging.NewMessage()
				h.writeBody(writer, message.Bytes())
				log.Printf("Sent message %s\n", message.Id)
				time.Sleep(5 * time.Second)

			}
		}

	default:
		h.writeStatus(writer, http.StatusNotFound)

	}
}

func (*Handler) writeBody(writer http.ResponseWriter, data []byte) {
	writer.Write(data)
	writer.(http.Flusher).Flush()
}

func (*Handler) writeStatus(writer http.ResponseWriter, statusCode int) {
	writer.Write(nil)
	writer.(http.Flusher).Flush()
}
