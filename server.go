package specialuse

import (
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/server"
)

type extension struct {}

func (ext *extension) Capabilities(state imap.ConnState) []string {
	var caps []string

	if state & imap.AuthenticatedState != 0 {
		caps = append(caps, Capability)
	}

	return caps
}

func (ext *extension) Command(name string) server.HandlerFactory {
	return nil
}

func NewExtension() server.Extension {
	return &extension{}
}
