package protocol

const (
	COMPRESSION_GZIP = "gzip"
)

const PROTO_VERSION = "1.0"

const CLIENT_END_MSG = "SERVERGOODBYE"

type CommonProtoFields struct {
	ProtocolVersion string `json:"protocol_version"`
	StatusCode      int    `json:"status_code"`
	CompressionType string `json:"compression_type"`
}
