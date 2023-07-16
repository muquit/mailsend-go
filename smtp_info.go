package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"net/textproto"
	"regexp"
	"strings"
	"time"
)

// smtpClient ...
// taken from golang's smtp.go. smtp package eats up code and message, I need to print smtp info,
// hence the upgly duplicating
type smtpClient struct {
	// Text is the textproto.Conn used by the smtpClient. It is exported to allow for
	// clients to add extensions.
	Text *textproto.Conn
	// keep a reference to the connection so it can be used to create a TLS
	// connection later
	conn net.Conn
	// whether the smtpClient is using TLS
	tls        bool
	serverName string
	// map of supported extensions
	ext map[string]string
	// supported auth mechanisms
	auth       []string
	localName  string // the name to use in HELO/EHLO
	didHello   bool   // whether we've said HELO/EHLO
	helloError error  // the error from the hello
}

func smtpAddr(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

// muquit@muquit.com - 2018-10-26 12:20:59
func printSMTPInfo(server string, port int, domain string, ssl bool, verifyCert bool) {
	addr := smtpAddr(server, port)
	var (
		conn       net.Conn
		err        error
		doStartTLS bool
	)
	if !ssl {
		doStartTLS = true
	} else {
		doStartTLS = false
	}
	host, _, _ := net.SplitHostPort(addr)

	if !ssl {
		conn, err = net.DialTimeout("tcp", addr, 5*time.Second)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true, // always skip
			ServerName:         host,
		}
		conn, err = tls.Dial("tcp", addr, tlsConfig)
		if err != nil {
			log.Fatal(err)
		}
	}

	text := textproto.NewConn(conn)

	code, msg, err := text.ReadResponse(220)
	if err != nil {
		text.Close()
		log.Fatal(err)
	}
	printMsg("[S] %d %s\n", code, msg)

	c := &smtpClient{
		Text:       text,
		conn:       conn,
		serverName: host,
		localName:  domain}
	_, c.tls = conn.(*tls.Conn)

	// HELO
	printMsg("[C] HELO %s\n", c.localName)
	_, _, err = c.cmd(250, "HELO %s", c.localName)
	if err != nil {
		text.Close()
		log.Fatal(err)
	}
	var startTLS bool
	if doStartTLS {
		startTLS, _ = c.Extension("STARTTLS")
		if startTLS {
			printMsg("[C] STARTTLS\n")
			config := &tls.Config{
				InsecureSkipVerify: !verifyCert,
				ServerName:         c.serverName,
			}
			err = c.StartTLS(config)
			if err != nil {
				text.Close()
				log.Fatal(err)
			}
		}
	}

	if ssl || startTLS {
		cs, _ := c.TLSConnectionState()
		certChain := cs.PeerCertificates
		cert := certChain[len(certChain)-1]
		printCertInfo(cert, c.serverName)
	}
	c.Quit()
}

func printCertInfo(cert *x509.Certificate, serverName string) {
	printMsg("Certificate of %s:\n", serverName)
	printMsg(" Version: %d (%#x)\n", cert.Version, cert.Version)
	printMsg(" Serial Number: %d (%#x)\n", cert.SerialNumber, cert.SerialNumber)
	printMsg(" Signature Algorithm: %s\n", cert.SignatureAlgorithm)
	printMsg(" Subject: %s\n", cert.Subject)
	printMsg(" Issuer: %s\n", cert.Issuer.CommonName)
	printMsg(" Not before: %s\n", cert.NotBefore.String())
	printMsg(" Not after: %s\n", cert.NotAfter.String())

}

// Close ...
func (c *smtpClient) Close() error {
	return c.Text.Close()
}

func (c *smtpClient) cmd(expectCode int, format string, args ...interface{}) (int, string, error) {
	id, err := c.Text.Cmd(format, args...)
	if err != nil {
		return 0, "", err
	}

	c.Text.StartResponse(id)
	defer c.Text.EndResponse(id)
	code, msg, err := c.Text.ReadResponse(expectCode)
	return code, msg, err

}

// StartTLS sends the STARTTLS command and encrypts all further communication.
// Only servers that advertise the STARTTLS extension support this function.
func (c *smtpClient) StartTLS(config *tls.Config) error {
	if err := c.hello(); err != nil {
		return err
	}
	code, msg, err := c.cmd(220, "STARTTLS")
	if err != nil {
		return err
	}
	printMsg("[S] %d-%s\n", code, msg)
	//	fmt.Printf("code %d, str: %s\n", code, str)
	c.conn = tls.Client(c.conn, config)
	c.Text = textproto.NewConn(c.conn)
	c.tls = true
	return c.ehlo()

}

// hello runs a hello exchange if needed.

func (c *smtpClient) hello() error {
	if !c.didHello {
		c.didHello = true
		err := c.ehlo()
		if err != nil {
			c.helloError = c.helo()
		}
	}
	return c.helloError
}
func printMsg(format string, a ...interface{}) {
	if !mailsend.options.Quiet {
		fmt.Printf(format, a...)
	}
}

// ehlo sends the EHLO (extended hello) greeting to the server. It
// should be the preferred greeting for servers that support it.
func (c *smtpClient) ehlo() error {
	printMsg("[C] EHLO %s\n", c.localName)
	code, msg, err := c.cmd(250, "EHLO %s", c.localName)
	if err != nil {
		return err
	}

	// msg contains response with new lines. So add the code at the
	// beginning of each line to make the output looks like mailsend
	re := regexp.MustCompile(`r?\n`)
	fmsg := re.ReplaceAllString(msg, "\n[S] 250-")

	printMsg("[S] %d-%s\n", code, fmsg)
	ext := make(map[string]string)
	extList := strings.Split(msg, "\n")
	if len(extList) > 1 {
		extList = extList[1:]
		for _, line := range extList {
			args := strings.SplitN(line, " ", 2)
			if len(args) > 1 {
				ext[args[0]] = args[1]
			} else {
				ext[args[0]] = ""
			}
		}
	}

	if mechs, ok := ext["AUTH"]; ok {
		c.auth = strings.Split(mechs, " ")
	}
	c.ext = ext
	return err
}

// helo sends the HELO greeting to the server. It should be used only when the
// server does not support ehlo.
func (c *smtpClient) helo() error {
	c.ext = nil
	_, _, err := c.cmd(250, "HELO %s", c.localName)
	return err
}

// Extension reports whether an extension is support by the server.
// The extension name is case-insensitive. If the extension is supported,
// Extension also returns a string that contains any parameters the
// server specifies for the extension.
func (c *smtpClient) Extension(ext string) (bool, string) {
	if err := c.hello(); err != nil {
		return false, ""
	}
	if c.ext == nil {
		return false, ""
	}
	ext = strings.ToUpper(ext)
	param, ok := c.ext[ext]
	return ok, param
}

// TLSConnectionState returns the client's TLS connection state.
// The return values are their zero values if StartTLS did
// not succeed.
func (c *smtpClient) TLSConnectionState() (state tls.ConnectionState, ok bool) {
	tc, ok := c.conn.(*tls.Conn)
	if !ok {
		return
	}
	return tc.ConnectionState(), true
}

// Quit sends the QUIT command and closes the connection to the server.
func (c *smtpClient) Quit() error {
	if err := c.hello(); err != nil {
		return err
	}
	printMsg("[C] QUIT\n")
	code, msg, err := c.cmd(221, "QUIT")
	if err != nil {
		return err
	}
	printMsg("[S] %d-%s\n", code, msg)
	return c.Text.Close()
}
