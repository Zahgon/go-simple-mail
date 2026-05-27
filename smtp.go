// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in https://raw.githubusercontent.com/golang/go/master/LICENSE

// Package mail implements the Simple Mail Transfer Protocol as defined in RFC 5321.
// It also implements the following extensions:
//	8BITMIME  RFC 1652
//	SMTPUTF8  RFC 6531
//	AUTH      RFC 2554
//	STARTTLS  RFC 3207
//	SIZE      RFC 1870
// Additional extensions may be handled by clients using smtp.go in golang source code or pull request Go Simple Mail

// smtp.go file is a modification of smtp golang package what is frozen and is not accepting new features.

package mail

import (
	"crypto/tls"
	"io"
	"net"
	"net/textproto"
)

// A Client represents a client connection to an SMTP server.
type smtpClient struct {
	// Text is the textproto.Conn used by the Client.
	text *textproto.Conn
	// keep a reference to the connection so it can be used to create a TLS
	// connection later
	conn net.Conn
	// whether the Client is using TLS
	tls        bool
	serverName string
	// map of supported extensions
	ext map[string]string
	// supported auth mechanisms
	a          []string
	localName  string // the name to use in HELO/EHLO
	didHello   bool   // whether we've said HELO/EHLO
	helloError error  // the error from the hello
}

// newClient returns a new smtpClient using an existing connection and host as a
// server name to be used when authenticating.
func newClient(conn net.Conn, host string) (*smtpClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close closes the connection.
func (c *smtpClient) close() error { _ = "STUB: not implemented"; return nil }

// hello runs a hello exchange if needed.
func (c *smtpClient) hello() error { _ = "STUB: not implemented"; return nil }

// hi sends a HELO or EHLO to the server as the given host name.
// Calling this method is only necessary if the client needs control
// over the host name used. The client will introduce itself as "localhost"
// automatically otherwise. If Hello is called, it must be called before
// any of the other methods.
func (c *smtpClient) hi(localName string) error { _ = "STUB: not implemented"; return nil }

// cmd is a convenience function that sends a command and returns the response
func (c *smtpClient) cmd(expectCode int, format string, args ...interface{}) (int, string, error) {
	_ = "STUB: not implemented"
	return 0, "", nil
}

// helo sends the HELO greeting to the server. It should be used only when the
// server does not support ehlo.
func (c *smtpClient) helo() error { _ = "STUB: not implemented"; return nil }

// ehlo sends the EHLO (extended hello) greeting to the server. It
// should be the preferred greeting for servers that support it.
func (c *smtpClient) ehlo() error { _ = "STUB: not implemented"; return nil }

// startTLS sends the STARTTLS command and encrypts all further communication.
// Only servers that advertise the STARTTLS extension support this function.
func (c *smtpClient) startTLS(config *tls.Config) error { _ = "STUB: not implemented"; return nil }

// authenticate authenticates a client using the provided authentication mechanism.
// A failed authentication closes the connection.
// Only servers that advertise the AUTH extension support this function.
func (c *smtpClient) authenticate(a auth) error { _ = "STUB: not implemented"; return nil }

// the last message isn't base64 because it isn't a challenge

// abort the AUTH

// mail issues a MAIL command to the server using the provided email address.
// If the server supports the 8BITMIME extension, Mail adds the BODY=8BITMIME
// parameter.
// If the server supports the SMTPUTF8 extension, Mail adds the
// SMTPUTF8 parameter.
// This initiates a mail transaction and is followed by one or more Rcpt calls.
func (c *smtpClient) mail(from string, extArgs ...map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// rcpt issues a RCPT command to the server using the provided email address.
// A call to Rcpt must be preceded by a call to Mail and may be followed by
// a Data call or another Rcpt call.
func (c *smtpClient) rcpt(to, dsn string) error { _ = "STUB: not implemented"; return nil }

type dataCloser struct {
	c *smtpClient
	io.WriteCloser
}

func (d *dataCloser) Close() error { _ = "STUB: not implemented"; return nil }

// data issues a DATA command to the server and returns a writer that
// can be used to write the mail headers and body. The caller should
// close the writer before calling any more methods on c. A call to
// Data must be preceded by one or more calls to Rcpt.
func (c *smtpClient) data() (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

// extension reports whether an extension is support by the server.
// The extension name is case-insensitive. If the extension is supported,
// extension also returns a string that contains any parameters the
// server specifies for the extension.
func (c *smtpClient) extension(ext string) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

// reset sends the RSET command to the server, aborting the current mail
// transaction.
func (c *smtpClient) reset() error { _ = "STUB: not implemented"; return nil }

// noop sends the NOOP command to the server. It does nothing but check
// that the connection to the server is okay.
func (c *smtpClient) noop() error { _ = "STUB: not implemented"; return nil }

// quit sends the QUIT command and closes the connection to the server.
func (c *smtpClient) quit() error { _ = "STUB: not implemented"; return nil }

// validateLine checks to see if a line has CR or LF as per RFC 5321
func validateLine(line string) error { _ = "STUB: not implemented"; return nil }
