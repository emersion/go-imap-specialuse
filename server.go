package specialuse

import (
	"github.com/emersion/go-imap/common"
	"github.com/emersion/go-imap/server"
)

// Enable this extension on a server. The backend must support this extension
// too.
func NewServer(s *server.Server) {
	s.RegisterCapability(Capability, common.AuthenticatedState)
}
