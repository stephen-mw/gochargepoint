// Package v5 soap client was auto-generated using wsld2go. Small edits were made to get it working.
package v5

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"net"
	"net/http"
	"time"
)

// SOAPEncoder is a soap encoder
type SOAPEncoder interface {
	Encode(v interface{}) error
	Flush() error
}

// SOAPDecoder is a soap decoder
type SOAPDecoder interface {
	Decode(v interface{}) error
}

// SOAPHeader header
type SOAPHeader struct {
	XMLName xml.Name    `xml:"Header"`
	Content interface{} `xml:",omitempty"`
}

// SOAPEnvelope is a soap envelope
type SOAPEnvelope struct {
	XMLName xml.Name   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  SOAPHeader `xml:",omitempty"`
	Body    SOAPBody
}

// SOAPBody represents a soap body
type SOAPBody struct {
	XMLName   xml.Name    `xml:"Body"`
	XMLNSWsse string      `xml:"xmlns:wsse,attr"`
	Fault     *SOAPFault  `xml:",omitempty"`
	Content   interface{} `xml:",omitempty"`
}

// UnmarshalXML unmarshals SOAPBody xml
func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, _ xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

func (f *SOAPFault) Error() string {
	return f.String
}

// Predefined WSS namespaces to be used in
const (
	WssNsWSSE       string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
	WssNsWSU        string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"
	WssNsType       string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText"
	mtomContentType string = `multipart/related; start-info="application/soap+xml"; type="application/xop+xml"; boundary="%s"`
)

// WSSSecurityHeader is a security header
type WSSSecurityHeader struct {
	XMLName   xml.Name `xml:"wsse:Security"`
	XMLNSWsse string   `xml:"xmlns:wsse,attr"`

	MustUnderstand string `xml:"mustUnderstand,attr,omitempty"`

	Token *WSSUsernameToken `xml:",omitempty"`
}

// WSSUsernameToken is a wsse username token
type WSSUsernameToken struct {
	XMLName   xml.Name `xml:"wsse:UsernameToken"`
	XMLNSWsu  string   `xml:"xmlns:wsu,attr"`
	XMLNSWsse string   `xml:"xmlns:wsse,attr"`

	ID string `xml:"wsu:Id,attr,omitempty"`

	Username *WSSUsername `xml:",omitempty"`
	Password *WSSPassword `xml:",omitempty"`
}

// WSSUsername is your username
type WSSUsername struct {
	XMLName   xml.Name `xml:"wsse:Username"`
	XMLNSWsse string   `xml:"xmlns:wsse,attr"`

	Data string `xml:",chardata"`
}

// WSSPassword is your password
type WSSPassword struct {
	XMLName   xml.Name `xml:"wsse:Password"`
	XMLNSWsse string   `xml:"xmlns:wsse,attr"`
	XMLNSType string   `xml:"Type,attr"`

	Data string `xml:",chardata"`
}

func newWSSSecurityHeader(user, pass, tokenID, mustUnderstand string) *WSSSecurityHeader {
	hdr := &WSSSecurityHeader{XMLNSWsse: WssNsWSSE, MustUnderstand: mustUnderstand}
	hdr.Token = &WSSUsernameToken{XMLNSWsu: WssNsWSU, XMLNSWsse: WssNsWSSE, ID: tokenID}
	hdr.Token.Username = &WSSUsername{XMLNSWsse: WssNsWSSE, Data: user}
	hdr.Token.Password = &WSSPassword{XMLNSWsse: WssNsWSSE, XMLNSType: WssNsType, Data: pass}
	return hdr
}

type basicAuth struct {
	Login    string
	Password string
}

type options struct {
	tlsCfg           *tls.Config
	auth             *basicAuth
	timeout          time.Duration
	contimeout       time.Duration
	tlshshaketimeout time.Duration
	client           HTTPClient
	httpHeaders      map[string]string
	mtom             bool
}

var defaultOptions = options{
	timeout:          time.Duration(30 * time.Second),
	contimeout:       time.Duration(90 * time.Second),
	tlshshaketimeout: time.Duration(15 * time.Second),
}

// A Option sets options such as credentials, tls, etc.
type Option func(*options)

// WithHTTPClient is an Option to set the HTTP client to use
// This cannot be used with WithTLSHandshakeTimeout, WithTLS,
// WithTimeout options
func WithHTTPClient(c HTTPClient) Option {
	return func(o *options) {
		o.client = c
	}
}

// WithTLSHandshakeTimeout is an Option to set default tls handshake timeout
// This option cannot be used with WithHTTPClient
func WithTLSHandshakeTimeout(t time.Duration) Option {
	return func(o *options) {
		o.tlshshaketimeout = t
	}
}

// WithRequestTimeout is an Option to set default end-end connection timeout
// This option cannot be used with WithHTTPClient
func WithRequestTimeout(t time.Duration) Option {
	return func(o *options) {
		o.contimeout = t
	}
}

// WithBasicAuth is an Option to set BasicAuth
func WithBasicAuth(login, password string) Option {
	return func(o *options) {
		o.auth = &basicAuth{Login: login, Password: password}
	}
}

// WithTLS is an Option to set tls config
// This option cannot be used with WithHTTPClient
func WithTLS(tls *tls.Config) Option {
	return func(o *options) {
		o.tlsCfg = tls
	}
}

// WithTimeout is an Option to set default HTTP dial timeout
func WithTimeout(t time.Duration) Option {
	return func(o *options) {
		o.timeout = t
	}
}

// WithHTTPHeaders is an Option to set global HTTP headers for all requests
func WithHTTPHeaders(headers map[string]string) Option {
	return func(o *options) {
		o.httpHeaders = headers
	}
}

// WithMTOM is an Option to set Message Transmission Optimization Mechanism
// MTOM encodes fields of type Binary using XOP.
func WithMTOM() Option {
	return func(o *options) {
		o.mtom = true
	}
}

type wssSecurity struct {
	XMLName        xml.Name `xml:"Security"`
	MustUnderstand string   `xml:"mustUnderstand,attr"`
	UsernameToken  struct {
		Text     string `xml:",chardata"`
		Username string `xml:"Username"`
		Password struct {
			Text string `xml:",chardata"`
			Type string `xml:"Type,attr"`
		} `xml:"Password"`
	} `xml:"UsernameToken"`
}

// NewWssSecurityHeader returns the security header used in chargepoint requests
func newWssSecurityHeader(u string, p string) SOAPHeader {
	sec := new(wssSecurity)
	sec.MustUnderstand = "1"
	sec.UsernameToken.Username = u
	sec.UsernameToken.Password.Text = p

	return SOAPHeader{
		Content: *sec,
	}
}

// Client is soap client
type Client struct {
	url    string
	opts   *options
	header SOAPHeader
}

// HTTPClient is a client which can make HTTP requests
// An example implementation is net/http.Client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// NewSoapClient creates new SOAP client instance
func NewSoapClient(user string, password string, uri string) *Client {
	opts := defaultOptions
	return &Client{
		url:    uri,
		opts:   &opts,
		header: newWssSecurityHeader(user, password),
	}
}

// CallContext performs HTTP POST request with a context
func (s *Client) CallContext(ctx context.Context, soapAction string, request, response interface{}) error {
	return s.call(ctx, soapAction, request, response)
}

// Call performs HTTP POST request
func (s *Client) Call(soapAction string, request, response interface{}) error {
	return s.call(context.Background(), soapAction, request, response)
}

func (s *Client) call(ctx context.Context, soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{}
	envelope.Header = s.header
	envelope.Body.XMLNSWsse = "ns1"

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)
	var encoder SOAPEncoder
	if s.opts.mtom {
		encoder = newMtomEncoder(buffer)
	} else {
		encoder = xml.NewEncoder(buffer)
	}

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.opts.auth != nil {
		req.SetBasicAuth(s.opts.auth.Login, s.opts.auth.Password)
	}

	req.WithContext(ctx)

	if s.opts.mtom {
		req.Header.Add("Content-Type", fmt.Sprintf(mtomContentType, encoder.(*mtomEncoder).Boundary()))
	} else {
		req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	}
	req.Header.Add("SOAPAction", soapAction)
	req.Header.Set("User-Agent", "gowsdl/0.1")
	if s.opts.httpHeaders != nil {
		for k, v := range s.opts.httpHeaders {
			req.Header.Set(k, v)
		}
	}
	req.Close = true

	client := s.opts.client
	if client == nil {
		tr := &http.Transport{
			TLSClientConfig: s.opts.tlsCfg,
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				d := net.Dialer{Timeout: s.opts.timeout}
				return d.DialContext(ctx, network, addr)
			},
			TLSHandshakeTimeout: s.opts.tlshshaketimeout,
		}
		client = &http.Client{Timeout: s.opts.contimeout, Transport: tr}
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}

	mtomBoundary, err := getMtomHeader(res.Header.Get("Content-Type"))
	if err != nil {
		return err
	}

	var dec SOAPDecoder
	if mtomBoundary != "" {
		dec = newMtomDecoder(res.Body, mtomBoundary)
	} else {
		dec = xml.NewDecoder(res.Body)
	}

	if err := dec.Decode(respEnvelope); err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}
