package protocol_commands

import (
	"encoding/json"
	"goqueue/pkg/protocol"
	"log"
)

type ConnectCommand struct {
}

func (cc *ConnectCommand) Execute(req []byte, resp *protocol.ProtoResponse) error {
	var r protocol.ProtoRequest
	err := json.Unmarshal(req, &r)
	log.Println(r)
	if err != nil {
		return err
	}
	//pp := protocol.NewPayloadParser(r.Payload)
	//err = pp.Parse()
	//if err != nil {
	//	return err
	//}
	cf := protocol.CommonProtoFields{ProtocolVersion: protocol.PROTO_VERSION, StatusCode: 200,
		CompressionType: protocol.COMPRESSION_GZIP}
	*resp = protocol.ProtoResponse{CommonProtoFields: cf}
	return nil
}
