package mail

// TODO: Remove this file before launch v3

// AddAttachment. DEPRECATED. Use Attach method. Allows you to add an attachment to the email message.
// You can optionally provide a different name for the file.
func (email *Email) AddAttachment(file string, name ...string) *Email {
	_ = "STUB: not implemented"
	return nil
}

// AddAttachmentData. DEPRECATED. Use Attach method. Allows you to add an in-memory attachment to the email message.
func (email *Email) AddAttachmentData(data []byte, filename, mimeType string) *Email {
	_ = "STUB: not implemented"
	return nil
}

// AddAttachmentBase64. DEPRECATED. Use Attach method. Allows you to add an attachment in base64 to the email message.
// You need provide a name for the file.
func (email *Email) AddAttachmentBase64(b64File, name string) *Email {
	_ = "STUB: not implemented"
	return nil
}

// AddInline. DEPRECATED. Use Attach method. Allows you to add an inline attachment to the email message.
// You can optionally provide a different name for the file.
func (email *Email) AddInline(file string, name ...string) *Email {
	_ = "STUB: not implemented"
	return nil
}

// AddInlineData. DEPRECATED. Use Attach method. Allows you to add an inline in-memory attachment to the email message.
func (email *Email) AddInlineData(data []byte, filename, mimeType string) *Email {
	_ = "STUB: not implemented"
	return nil
}

// AddInlineBase64. DEPRECATED. Use Attach method. Allows you to add an inline in-memory base64 encoded attachment to the email message.
// You need provide a name for the file. If mimeType is an empty string, attachment mime type will be deduced
// from the file name extension and defaults to application/octet-stream.
func (email *Email) AddInlineBase64(b64File, name, mimeType string) *Email {
	_ = "STUB: not implemented"
	return nil
}
