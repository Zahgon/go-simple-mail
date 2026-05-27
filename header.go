// headers.go implements "Q" encoding as specified by RFC 2047.
//Modified from https://github.com/joegrasse/mime to use with Go Simple Mail

package mail

import (
	"bufio"
	"io"
)

type encoder struct {
	w         *bufio.Writer
	charset   string
	encoding  headerEncoding
	usedChars int
}

// newEncoder returns a new mime header encoder that writes to w. The c
// parameter specifies the name of the character set of the text that will be
// encoded. The u parameter indicates how many characters have been used
// already.
func newEncoder(w io.Writer, c string, encoding headerEncoding, u int) *encoder {
	_ = "STUB: not implemented"
	return nil
}

// encode encodes p using the encoding scheme specified in e
// If all chars are printable ascii chars, no encoding is performed.
// Limits line length to 75 characters and folds lines as necessary.
func (e *encoder) encode(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// some lines we encode end in "
//maxLineLength := 75 - 1

// prevent header injection

// check to see if we have all printable characters

// all characters are printable. just do line folding

// split the line where necessary

// check line length

// first word on newline needs a space in front

// reset since not on the first line anymore

/*word*/

// else block can only be HeaderEncodingQ as of now

// A single encoded word can not be longer than 75 characters

// encode the character

// Check line length

// reset since not on the first line anymore

// encode takes a string and position in that string and encodes one utf-8
// character. It then returns the encoded string and number of runes in the
// character.
func encode(text []byte, i int) (encodedString string, runeLength int) {
	_ = "STUB: not implemented"
	return "", 0
}

// secureHeader removes all unnecessary values to prevent
// header injection
func secureHeader(text []byte) []byte { _ = "STUB: not implemented"; return nil }

// isVchar returns true if c is an RFC 5322 VCHAR character.
func isVchar(c byte) bool {
	_ = "STUB: not implemented"
	// Visible (printing) characters.
	return false
}

// isWSP returns true if c is a WSP (white space).
// WSP is a space or horizontal tab (RFC5234 Appendix B).
func isWSP(c byte) bool { _ = "STUB: not implemented"; return false }
