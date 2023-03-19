package internal

import (
	"encoding/json"
	"time"
)

const (
	COMPRESSION_GZIP = "gzip"
)

const (
	StatusOK                    = 200
	StatusCommandNotFound       = 400
	StatusBrokerInShutdownState = 500
)

const (
	ClientProducer = iota
	ClientConsumer
)

type ClientRequest struct {
	ClientType    int               `json:"client_type"`
	Command       string            `json:"command"`
	CommandParams map[string]string `json:"command_params"`
}

func (c ClientRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c)

}

type CommonResponse struct {
	StatusCode  uint16 `json:"status_code"`
	MessageBody string `json:"message_body"`
}

type CommandResponse struct {
	CommonResponse
}

type ProtocolResponse struct {
	CommonResponse
	TimeStamp       time.Time         `json:"time_stamp"`
	Headers         map[string]string `json:"headers"`
	CompressionType string            `json:"compression_type"`
}

func (p ProtocolResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(&p)
}
