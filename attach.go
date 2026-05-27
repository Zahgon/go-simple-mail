package mail

// File represents the file that can be added to the email message.
// You can add attachment from file in path, from base64 string or from []byte.
// You can define if attachment is inline or not.
// Only one, Data, B64Data or FilePath is supported. If multiple are set, then
// the first in that order is used.
type File struct {
	// FilePath is the path of the file to attach.
	FilePath string
	// ContentID is the contentID of the attachment. Optional. Used instead of Name to look up inline attachment in the body if provided.
	ContentID string
	// Name is the name of file in attachment. Required for Data and B64Data. Optional for FilePath.
	Name string
	// MimeType of attachment. If empty then is obtained from Name (if not empty) or FilePath. If cannot obtained, application/octet-stream is set.
	MimeType string
	// B64Data is the base64 string to attach.
	B64Data string
	// Data is the []byte of file to attach.
	Data []byte
	// Inline defines if attachment is inline or not.
	Inline bool
}

type attachType int

const (
	attachData attachType = iota
	attachB64
	attachFile
)

// Attach allows you to add an attachment to the email message.
// The attachment can be inlined
func (email *Email) Attach(file *File) *Email { _ = "STUB: not implemented"; return nil }

// if no alternative name was provided, get the filename

// get the mimetype

func getAttachmentType(file *File, allowEmptyAttachments bool) (attachType, error) {
	_ = "STUB: not implemented"
	// 1- data
	// 2- base64
	// 3- file
	return *new(attachType), nil
}

// first check if Data

// data requires a name

// check if base64

// B64Data requires a name

// check if file

// attachB64 does the low level attaching of the files but decoding base64
func (email *Email) attachB64(file *File) error {
	_ = "STUB: not implemented"

	// decode the string
	return nil
}

func (email *Email) attachFile(file *File) error { _ = "STUB: not implemented"; return nil }

// attachData does the low level attaching of the in-memory data
func (email *Email) attachData(file *File) {
	_ = "STUB: not implemented"
	// use inlines and attachments because is necessary to know if message has related parts and mixed parts
	return
}
