package client

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/th3shadowbroker/sse-example/messaging"
)

type SSEClient struct {
	Messages chan messaging.Message
}

func NewSSEClient() *SSEClient {
	return &SSEClient{
		Messages: make(chan messaging.Message),
	}
}

func (c *SSEClient) Connect(url string) {
	res, req, err := c.getResponse(url)
	if err != nil {
		panic(err)
	}
	//defer req.Body.Close()

	reader := bufio.NewReader(res.Body)

	for {
		select {

		case <-req.Context().Done():
			return

		default:
			line, err := reader.ReadBytes('\n')
			if err != nil {
				if os.IsTimeout(err) {
					c.Connect(url)
					return
				}
				panic(err)
			}

			var message messaging.Message
			if err := json.Unmarshal(line, &message); err != nil {
				panic(err)
			}

			c.Messages <- message
		}
	}
}

func (c *SSEClient) getResponse(url string) (*http.Response, *http.Request, error) {
	var client = http.Client{
		Timeout: 10 * time.Second,
	}

	var request, _ = http.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("Accept", "application/stream+json")

	response, err := client.Do(request)
	return response, request, err
}
