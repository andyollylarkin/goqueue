package protocol

type PayloadParser struct {
}

func NewPayloadParser(payload []byte) *PayloadParser {
	return &PayloadParser{}
}

func (pp *PayloadParser) Parse() error { return nil }
