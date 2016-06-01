// Special-Use mailboxes extension, as defined in RFC 6154.
package specialuse

import (
	"github.com/emersion/go-imap/client"
)

// The SPECIAL-USE capability.
const Capability = "SPECIAL-USE"

// Check if a client supports SPECIAL-USE.
func Supported(c *client.Client) bool {
	return c.Caps[Capability]
}
