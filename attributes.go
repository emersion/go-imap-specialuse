package specialuse

// Mailbox attributes defined in RFC 6154 section 2.
const (
	// This mailbox presents all messages in the user's message store.
	All = "\\All"
	// This mailbox is used to archive messages.
	Archive = "\\Archive"
	// This mailbox is used to hold draft messages -- typically, messages that are
	// being composed but have not yet been sent.
	Drafts = "\\Drafts"
	// This mailbox presents all messages marked in some way as "important".
	Flagged = "\\Flagged"
	// This mailbox is where messages deemed to be junk mail are held.
	Junk = "\\Junk"
	// This mailbox is used to hold copies of messages that have been sent.
	Sent = "\\Sent"
	// This mailbox is used to hold messages that have been deleted or marked for
	// deletion.
	Trash = "\\Trash"
	// This mailbox contains messages that are likely important to the user.
	// (defined in RFC 8457)
	Important = "\\Important"
)
