package specialuse

import (
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/server"
)

type extension struct{}

func (ext *extension) Capabilities(c server.Conn) []string {
	if c.Context().State&imap.AuthenticatedState != 0 {
		return []string{Capability, ImportantCapability}
	}
	return nil
}

func (ext *extension) Command(name string) server.HandlerFactory {
	return nil
}

func NewExtension() server.Extension {
	return &extension{}
}
