package mail

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/textproto"
	"strings"
)

type message struct {
	headers        textproto.MIMEHeader
	body           *bytes.Buffer
	writers        []*multipart.Writer
	parts          uint8
	cids           map[string]string
	charset        string
	encoding       encoding
	headerEncoding headerEncoding // Only None and Q are currently supported
}

func newMessage(email *Email) *message { _ = "STUB: not implemented"; return nil }

func encodeHeader(text string, charset string, encoding headerEncoding, usedChars int) string {
	_ = "STUB: not implemented"
	// create buffer
	return ""
}

// encode

// getHeaders returns the message headers
func (msg *message) getHeaders() (headers string) {
	_ = "STUB: not implemented"
	// if the date header isn't set, set it
	return ""
}

// encode and combine the headers

// getCID gets the generated CID for the provided text
func (msg *message) getCID(text string) (cid string) {
	_ = "STUB: not implemented"
	// set the date format to use
	return ""
}

// get the cid if we have one

// generate a new cid

// save it

// replaceCIDs replaces the CIDs found in a text string
// with generated ones
func (msg *message) replaceCIDs(input []byte) []byte { _ = "STUB: not implemented"; return nil }

// One process replaceCIDs if we have anything to replace

// regular expression to find cids

// replace all of the found cids with generated ones

// openMultipart creates a new part of a multipart message
func (msg *message) openMultipart(multipartType string) {
	_ = "STUB: not implemented"
	// create a new multipart writer
	return
}

// create the boundary

// if no existing parts, add header to main header group

// add header to multipart section

// closeMultipart closes a part of a multipart message
func (msg *message) closeMultipart() { _ = "STUB: not implemented"; return }

// base64Encode base64 encodes the provided text with line wrapping
func base64Encode(text []byte) []byte {
	_ = "STUB: not implemented"
	// create buffer
	return nil
}

// create base64 encoder that linewraps

// write the encoded text to buf

// qpEncode uses the quoted-printable encoding to encode the provided text
func qpEncode(text []byte) []byte {
	_ = "STUB: not implemented"
	// create buffer
	return nil
}

const maxLineChars = 76

type base64LineWrap struct {
	writer       io.Writer
	numLineChars int
}

func (e *base64LineWrap) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"

	// while we have more chars than are allowed
	return 0, nil
}

// write the chars we can

// write a line break

// reset the line count

// remove the chars that have been written

// set the num of chars written

// write what is left

func (msg *message) write(header textproto.MIMEHeader, body []byte, encoding encoding) {
	_ = "STUB: not implemented"
	return
}

func (msg *message) writeHeader(headers textproto.MIMEHeader) {
	_ = "STUB: not implemented"
	// if there are no parts add header to main headers
	return
}

// add header to multipart section

func (msg *message) writeBody(body []byte, encoding encoding) {
	_ = "STUB: not implemented"
	// encode and write the body
	return
}

func (msg *message) addBody(contentType string, body []byte) { _ = "STUB: not implemented"; return }

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string { _ = "STUB: not implemented"; return "" }

func (msg *message) addFiles(files []*File, inline bool) { _ = "STUB: not implemented"; return }
