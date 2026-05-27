package mail

import (
	"bytes"
	"crypto/tls"
	"net"
	"net/textproto"
	"sync"
	"time"

	"github.com/toorop/go-dkim"
)

// Email represents an email message.
type Email struct {
	from           string
	sender         string
	replyTo        string
	returnPath     string
	recipients     []string
	headers        textproto.MIMEHeader
	parts          []part
	attachments    []*File
	inlines        []*File
	Charset        string
	Encoding       encoding
	HeaderEncoding headerEncoding
	Error          error
	SMTPServer     *smtpClient
	DkimMsg        string

	// UseProvidedAddress if set to true will disable any parsing and
	// validation of addresses and uses the address provided by the user
	// without any modifications
	UseProvidedAddress bool

	// AllowEmptyAttachments if enabled, allows you you attach empty
	// items, a file without any associated data
	AllowEmptyAttachments     bool
	AllowDuplicateAddress     bool
	AddBccToHeader            bool
	preserveOriginalRecipient bool
	dsn                       []DSN
}

/*
SMTPServer represents a SMTP Server
If authentication is CRAM-MD5 then the Password is the Secret
*/
type SMTPServer struct {
	Authentication AuthType
	Encryption     Encryption
	Username       string
	Password       string
	Helo           string
	ConnectTimeout time.Duration
	SendTimeout    time.Duration
	Host           string
	Port           int
	KeepAlive      bool
	TLSConfig      *tls.Config

	// use custom dialer
	CustomConn net.Conn
}

// SMTPClient represents a SMTP Client for send email
type SMTPClient struct {
	mu                        sync.Mutex
	Client                    *smtpClient
	SendTimeout               time.Duration
	KeepAlive                 bool
	hasDSNExt                 bool
	preserveOriginalRecipient bool
	dsn                       []DSN
}

// part represents the different content parts of an email body.
type part struct {
	contentType string
	body        *bytes.Buffer
}

// Encryption type to enum encryption types (None, SSL/TLS, STARTTLS)
type Encryption int

// TODO: Remove EncryptionSSL and EncryptionTLS before launch v3

const (
	// EncryptionNone uses no encryption when sending email
	EncryptionNone Encryption = iota
	// EncryptionSSL: DEPRECATED. Use EncryptionSSLTLS. Sets encryption type to SSL/TLS when sending email
	EncryptionSSL
	// EncryptionTLS: DEPRECATED. Use EncryptionSTARTTLS. sets encryption type to STARTTLS when sending email
	EncryptionTLS
	// EncryptionSSLTLS sets encryption type to SSL/TLS when sending email
	EncryptionSSLTLS
	// EncryptionSTARTTLS sets encryption type to STARTTLS when sending email
	EncryptionSTARTTLS
)

// TODO: Remove last two indexes
var encryptionTypes = [...]string{"None", "SSL/TLS", "STARTTLS", "SSL/TLS", "STARTTLS"}

func (encryption Encryption) String() string { _ = "STUB: not implemented"; return "" }

type headerEncoding int

const (
	// HeaderEncodingNone turns off encoding on the message headers
	// https://www.rfc-editor.org/rfc/rfc6530#section-7.1
	HeaderEncodingNone headerEncoding = iota

	// TODO: Add Base64 encoding
	// HeaderEncodingBase64 sets the message header encoding to base64
	// https://www.rfc-editor.org/rfc/rfc2045#section-6.8
	// HeaderEncodingBase64

	// HeaderEncodingQ sets the message header encoding to Q encoding
	// https://www.rfc-editor.org/rfc/rfc2047#section-4.2
	HeaderEncodingQ
)

type encoding int

const (
	// EncodingNone turns off encoding on the message body
	EncodingNone encoding = iota
	// EncodingBase64 sets the message body encoding to base64
	EncodingBase64
	// EncodingQuotedPrintable sets the message body encoding to quoted-printable
	EncodingQuotedPrintable
)

var encodingTypes = [...]string{"binary", "base64", "quoted-printable"}

func (encoding encoding) string() string { _ = "STUB: not implemented"; return "" }

type ContentType int

const (
	// TextPlain sets body type to text/plain in message body
	TextPlain ContentType = iota
	// TextHTML sets body type to text/html in message body
	TextHTML
	// TextCalendar sets body type to text/calendar in message body
	TextCalendar
	// TextAMP sets body type to text/x-amp-html in message body
	TextAMP
)

var contentTypes = [...]string{"text/plain", "text/html", "text/calendar", "text/x-amp-html"}

func (contentType ContentType) string() string { _ = "STUB: not implemented"; return "" }

type AuthType int

const (
	// AuthPlain implements the PLAIN authentication
	AuthPlain AuthType = iota
	// AuthLogin implements the LOGIN authentication
	AuthLogin
	// AuthCRAMMD5 implements the CRAM-MD5 authentication
	AuthCRAMMD5
	// AuthNone for SMTP servers without authentication
	AuthNone
	// AuthAuto (default) use the first AuthType of the list of returned types supported by SMTP
	AuthAuto
)

func (at AuthType) String() string { _ = "STUB: not implemented"; return "" }

/*
	DSN notifications

- 'NEVER' under no circumstances a DSN must be returned to the sender. If you use NEVER all other notifications will be ignored.

- 'SUCCESS' will notify you when your mail has arrived at its destination.

- 'FAILURE' will arrive if an error occurred during delivery.

- 'DELAY' will notify you if there is an unusual delay in delivery, but the actual delivery's outcome (success or failure) is not yet decided.

see https://tools.ietf.org/html/rfc3461 See section 4.1 for more information about NOTIFY
*/
type DSN int

const (
	NEVER DSN = iota
	FAILURE
	DELAY
	SUCCESS
)

var dsnTypes = [...]string{"NEVER", "FAILURE", "DELAY", "SUCCESS"}

func (dsn DSN) String() string { _ = "STUB: not implemented"; return "" }

// NewMSG creates a new email. It uses UTF-8 by default. All charsets: http://webcheatsheet.com/HTML/character_sets_list.php
func NewMSG() *Email { _ = "STUB: not implemented"; return nil }

// NewSMTPClient returns the client for send email
func NewSMTPClient() *SMTPServer { _ = "STUB: not implemented"; return nil }

// GetEncryptionType returns the encryption type used to connect to SMTP server
func (server *SMTPServer) GetEncryptionType() Encryption {
	_ = "STUB: not implemented"
	return *

	// GetError returns the first email error encountered
	new(Encryption)
}

func (email *Email) GetError() error {
	_ = "STUB: not implemented"

	// SetFrom sets the From address.
	return nil
}

func (email *Email) SetFrom(address string) *Email { _ = "STUB: not implemented"; return nil }

// SetSender sets the Sender address.
func (email *Email) SetSender(address string) *Email { _ = "STUB: not implemented"; return nil }

// SetReplyTo sets the Reply-To address.
func (email *Email) SetReplyTo(address string) *Email { _ = "STUB: not implemented"; return nil }

// SetReturnPath sets the Return-Path address. This is most often used
// to send bounced emails to a different email address.
func (email *Email) SetReturnPath(address string) *Email { _ = "STUB: not implemented"; return nil }

// AddTo adds a To address. You can provide multiple
// addresses at the same time.
func (email *Email) AddTo(addresses ...string) *Email { _ = "STUB: not implemented"; return nil }

// AddCc adds a Cc address. You can provide multiple
// addresses at the same time.
func (email *Email) AddCc(addresses ...string) *Email { _ = "STUB: not implemented"; return nil }

// AddBcc adds a Bcc address. You can provide multiple
// addresses at the same time.
func (email *Email) AddBcc(addresses ...string) *Email { _ = "STUB: not implemented"; return nil }

// AddAddresses allows you to add addresses to the specified address header.
func (email *Email) AddAddresses(header string, addresses ...string) *Email {
	_ = "STUB: not implemented"
	return nil
}

// check for a valid address header

// check to see if the addresses are valid

// ignore empty addresses

// check for more than one address

// other address types can have more than one address

// save the address

// delete the current "From" to set the new
// when "From" need to be changed in the message

// check that the address was added to the recipients list

// make sure the from and sender addresses are different

// add Bcc only if AddBccToHeader is true

// add all addresses to the headers except for Bcc and Return-Path

// add the address to the headers

// addAddress adds an address to the address list if it hasn't already been added
func addAddress(addressList []string, address string, allowDuplicateAddress bool) ([]string, error) {
	_ = "STUB: not implemented"
	return nil,

		// loop through the address list to check for dups
		nil
}

type Priority int

const (
	// PriorityLow sets the email Priority to Low
	PriorityLow Priority = iota
	// PriorityHigh sets the email Priority to High
	PriorityHigh
)

// SetPriority sets the email message Priority. Use with
// either "High" or "Low".
func (email *Email) SetPriority(priority Priority) *Email { _ = "STUB: not implemented"; return nil }

// SetDate sets the date header to the provided date/time.
// The format of the string should be YYYY-MM-DD HH:MM:SS Time Zone.
//
// Example: SetDate("2015-04-28 10:32:00 CDT")
func (email *Email) SetDate(dateTime string) *Email { _ = "STUB: not implemented"; return nil }

// Try to parse the provided date/time

// SetSubject sets the subject of the email message.
func (email *Email) SetSubject(subject string) *Email { _ = "STUB: not implemented"; return nil }

// SetListUnsubscribe sets the Unsubscribe address.
func (email *Email) SetListUnsubscribe(address string) *Email {
	_ = "STUB: not implemented"
	return nil
}

// SetDkim adds DomainKey signature to the email message (header+body)
func (email *Email) SetDkim(options dkim.SigOptions) *Email { _ = "STUB: not implemented"; return nil }

// SetBody sets the body of the email message.
func (email *Email) SetBody(contentType ContentType, body string) *Email {
	_ = "STUB: not implemented"
	return nil
}

// SetBodyData sets the body of the email message from []byte
func (email *Email) SetBodyData(contentType ContentType, body []byte) *Email {
	_ = "STUB: not implemented"
	return nil
}

// AddHeader adds the given "header" with the passed "value".
func (email *Email) AddHeader(header string, values ...string) *Email {
	_ = "STUB: not implemented"
	return nil
}

// check that there is actually a value

// Set header to correct canonical Mime

// AddHeaders is used to add multiple headers at once
func (email *Email) AddHeaders(headers textproto.MIMEHeader) *Email {
	_ = "STUB: not implemented"
	return nil
}

// AddAlternative allows you to add alternative parts to the body
// of the email message. This is most commonly used to add an
// html version in addition to a plain text version that was
// already added with SetBody.
func (email *Email) AddAlternative(contentType ContentType, body string) *Email {
	_ = "STUB: not implemented"
	return nil
}

// AddAlternativeData allows you to add alternative parts to the body
// of the email message. This is most commonly used to add an
// html version in addition to a plain text version that was
// already added with SetBody.
func (email *Email) AddAlternativeData(contentType ContentType, body []byte) *Email {
	_ = "STUB: not implemented"
	return nil
}

// SetDSN sets the delivery status notification list, only is set when SMTP server supports DSN extension
//
// To preserve the original recipient of an email message, for example, if it is forwarded to another address, set preserveOriginalRecipient to true
func (email *Email) SetDSN(dsn []DSN, preserveOriginalRecipient bool) *Email {
	_ = "STUB: not implemented"
	return nil
}

// GetFrom returns the sender of the email, if any
func (email *Email) GetFrom() string { _ = "STUB: not implemented"; return "" }

// GetRecipients returns a slice of recipients emails
func (email *Email) GetRecipients() []string { _ = "STUB: not implemented"; return nil }

func (email *Email) hasMixedPart() bool { _ = "STUB: not implemented"; return false }

func (email *Email) hasRelatedPart() bool { _ = "STUB: not implemented"; return false }

func (email *Email) hasAlternativePart() bool { _ = "STUB: not implemented"; return false }

// GetMessage builds and returns the email message (RFC822 formatted message)
func (email *Email) GetMessage() string { _ = "STUB: not implemented"; return "" }

// Send sends the composed email
func (email *Email) Send(client *SMTPClient) error { _ = "STUB: not implemented"; return nil }

// SendEnvelopeFrom sends the composed email with envelope
// sender. 'from' must be an email address.
func (email *Email) SendEnvelopeFrom(from string, client *SMTPClient) error {
	_ = "STUB: not implemented"
	return nil
}

// dial connects to the smtp server with the request encryption type
func dial(customConn net.Conn, host string, port string, encryption Encryption, config *tls.Config) (*smtpClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// do the actual dial

// TODO: Remove EncryptionSSL check before launch v3

// smtpConnect connects to the smtp server and starts TLS and passes auth
// if necessary
func smtpConnect(customConn net.Conn, host, port, helo string, encryption Encryption, config *tls.Config) (*smtpClient, error) {
	_ = "STUB: not implemented"
	// connect to the mail server
	return nil, nil
}

// send Helo

// STARTTLS if necessary
// TODO: Remove EncryptionTLS check before launch v3

func (server *SMTPServer) getAuth(a string) (auth, error) {
	_ = "STUB: not implemented"
	return *new(auth), nil
}

func (server *SMTPServer) validateAuth(c *smtpClient) error { _ = "STUB: not implemented"; return nil }

// Determine Auth type automatically from extension

// Connect returns the smtp client
func (server *SMTPServer) Connect() (*SMTPClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if there is a ConnectTimeout, setup the channel and do the connect under a goroutine

// send the result

// get the connect result or timeout result, which ever happens first

// no ConnectTimeout, just fire the connect

// Reset send RSET command to smtp client
func (smtpClient *SMTPClient) Reset() error { _ = "STUB: not implemented"; return nil }

// Noop send NOOP command to smtp client
func (smtpClient *SMTPClient) Noop() error { _ = "STUB: not implemented"; return nil }

// Quit send QUIT command to smtp client
func (smtpClient *SMTPClient) Quit() error { _ = "STUB: not implemented"; return nil }

// Close closes the connection
func (smtpClient *SMTPClient) Close() error { _ = "STUB: not implemented"; return nil }

// SendMessage sends a message (a RFC822 formatted message)
// 'from' must be an email address, recipients must be a slice of email address
func SendMessage(from string, recipients []string, msg string, client *SMTPClient) error {
	_ = "STUB: not implemented"
	return nil
}

// send does the low level sending of the email
func send(from string, to []string, msg string, client *SMTPClient) error {
	_ = "STUB: not implemented"
	//Check if client struct is not nil
	return nil
}

//Check if client is not nil

// if there is a SendTimeout, setup the channel and do the send under a goroutine

// no SendTimeout, just fire the sendMailProcess

// get the send result or timeout result, which ever happens first

func sendMailProcess(from string, to []string, msg string, c *SMTPClient) error {
	_ = "STUB: not implemented"
	return nil
}

// Set the sender

// Set the recipients

// Send the data command

// write the message

// check if keepAlive for close or reset
func checkKeepAlive(client *SMTPClient) { _ = "STUB: not implemented"; return }

func hasNeverDSN(dsnList []DSN) bool { _ = "STUB: not implemented"; return false }

func dsnToString(dsnList []DSN) []string { _ = "STUB: not implemented"; return nil }
