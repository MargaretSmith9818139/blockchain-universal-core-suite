package main

type LightClient struct {
	ClientID string
	Headers  []BlockHeader
	Latest   int
}

type BlockHeader struct {
	Height int
	Hash   string
	Prev   string
}

func NewLightClient(id string) *LightClient {
	return &LightClient{
		ClientID: id,
		Headers:  []BlockHeader{},
	}
}

func (l *LightClient) SyncHeader(header BlockHeader) {
	l.Headers = append(l.Headers, header)
	l.Latest = header.Height
}

func (l *LightClient) VerifyBlock(hash string) bool {
	for _, h := range l.Headers {
		if h.Hash == hash {
			return true
		}
	}
	return false
}
