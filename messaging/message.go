package messaging

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Message struct {
	Id string `json:"id"`
}

func NewMessage() Message {
	return Message{
		Id: uuid.NewString(),
	}
}

func (m *Message) Bytes() []byte {
	messageBytes, _ := json.Marshal(m)
	return append(messageBytes, byte('\n'))
}
