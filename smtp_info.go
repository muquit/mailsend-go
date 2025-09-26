package main

import (
	"crypto/tls"
	"crypto/x509"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
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
func printSMTPInfo(server string, port int, domain string, ssl bool, verifyCert bool, printCerts bool) {
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
	if len(c.localName) == 0 {
		c.localName = "localhost"
	}
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
		//certChain := cs.PeerCertificates
		//cert := certChain[len(certChain)-1]
		//printCertInfo(cert, c.serverName)
		if printCerts {
			printTLSConnectionInfo(cs)
		}
	}
	c.Quit()
}

func printTLSConnectionInfo(state tls.ConnectionState) {
    fmt.Println("\n" + strings.Repeat("=", 60))
    fmt.Println("TLS CONNECTION INFORMATION")
    fmt.Println(strings.Repeat("=", 60))

    // TLS Version mapping
    versionMap := map[uint16]string{
        tls.VersionSSL30: "SSL 3.0",
        tls.VersionTLS10: "TLS 1.0",
        tls.VersionTLS11: "TLS 1.1",
        tls.VersionTLS12: "TLS 1.2",
        tls.VersionTLS13: "TLS 1.3",
    }

    version := fmt.Sprintf("0x%04x (Unknown)", state.Version)
    if v, ok := versionMap[state.Version]; ok {
        version = fmt.Sprintf("%s (0x%04x)", v, state.Version)
    }

    fmt.Printf("TLS Version:           %s\n", version)
    fmt.Printf("Cipher Suite:          0x%04x\n", state.CipherSuite)
    fmt.Printf("Server Certificates:   %d\n", len(state.PeerCertificates))
    fmt.Printf("Handshake Complete:    %v\n", state.HandshakeComplete)
    fmt.Printf("Did Resume:            %v\n", state.DidResume)
    fmt.Printf("Negotiated Protocol:   %s\n", state.NegotiatedProtocol)

    if len(state.PeerCertificates) == 0 {
        fmt.Println("\nNo certificates found")
        return
    }

    // Print certificate chain
    for i, cert := range state.PeerCertificates {
        fmt.Printf("\n%s\n", strings.Repeat("-", 60))
        if i == 0 {
            fmt.Printf("SERVER CERTIFICATE #%d (Leaf Certificate)\n", i+1)
        } else {
            fmt.Printf("INTERMEDIATE CERTIFICATE #%d\n", i+1)
        }
        fmt.Printf("%s\n", strings.Repeat("-", 60))

        printCertificateDetails(cert)
    }
}

func printCertificateDetails(cert *x509.Certificate) {
	fmt.Printf("Subject:               %s\n", cert.Subject.String())
	fmt.Printf("Issuer:                %s\n", cert.Issuer.String())
	fmt.Printf("Serial Number:         %s\n", cert.SerialNumber.String())
	fmt.Printf("Version:               %d\n", cert.Version)
	
	// Validity period with time until expiry
	fmt.Printf("Valid From:            %s\n", cert.NotBefore.Format("2006-01-02 15:04:05 UTC"))
	fmt.Printf("Valid Until:           %s\n", cert.NotAfter.Format("2006-01-02 15:04:05 UTC"))
	
	// Check if certificate is expired or will expire soon
	now := time.Now()
	if now.After(cert.NotAfter) {
		fmt.Printf("Status:                ⚠️  EXPIRED\n")
	} else if now.Add(30*24*time.Hour).After(cert.NotAfter) {
		daysLeft := int(cert.NotAfter.Sub(now).Hours() / 24)
		fmt.Printf("Status:                ⚠️  EXPIRES IN %d DAYS\n", daysLeft)
	} else {
		daysLeft := int(cert.NotAfter.Sub(now).Hours() / 24)
		fmt.Printf("Status:                ✅ Valid (%d days remaining)\n", daysLeft)
	}
	
	// Subject Alternative Names
	if len(cert.DNSNames) > 0 {
		fmt.Printf("DNS Names:             %s\n", strings.Join(cert.DNSNames, ", "))
	}
	if len(cert.IPAddresses) > 0 {
		ips := make([]string, len(cert.IPAddresses))
		for i, ip := range cert.IPAddresses {
			ips[i] = ip.String()
		}
		fmt.Printf("IP Addresses:          %s\n", strings.Join(ips, ", "))
	}
	if len(cert.EmailAddresses) > 0 {
		fmt.Printf("Email Addresses:       %s\n", strings.Join(cert.EmailAddresses, ", "))
	}
	
	// Key information
	fmt.Printf("Public Key Algorithm:  %s\n", cert.PublicKeyAlgorithm.String())
	fmt.Printf("Signature Algorithm:   %s\n", cert.SignatureAlgorithm.String())
	
	// Key usage
	if cert.KeyUsage != 0 {
		var usages []string
		if cert.KeyUsage&x509.KeyUsageDigitalSignature != 0 {
			usages = append(usages, "Digital Signature")
		}
		if cert.KeyUsage&x509.KeyUsageKeyEncipherment != 0 {
			usages = append(usages, "Key Encipherment")
		}
		if cert.KeyUsage&x509.KeyUsageKeyAgreement != 0 {
			usages = append(usages, "Key Agreement")
		}
		if cert.KeyUsage&x509.KeyUsageCertSign != 0 {
			usages = append(usages, "Certificate Signing")
		}
		if cert.KeyUsage&x509.KeyUsageCRLSign != 0 {
			usages = append(usages, "CRL Signing")
		}
		fmt.Printf("Key Usage:             %s\n", strings.Join(usages, ", "))
	}
	
	// Extended key usage
	if len(cert.ExtKeyUsage) > 0 {
		var extUsages []string
		for _, usage := range cert.ExtKeyUsage {
			switch usage {
			case x509.ExtKeyUsageServerAuth:
				extUsages = append(extUsages, "Server Authentication")
			case x509.ExtKeyUsageClientAuth:
				extUsages = append(extUsages, "Client Authentication")
			case x509.ExtKeyUsageEmailProtection:
				extUsages = append(extUsages, "Email Protection")
			case x509.ExtKeyUsageTimeStamping:
				extUsages = append(extUsages, "Time Stamping")
			case x509.ExtKeyUsageCodeSigning:
				extUsages = append(extUsages, "Code Signing")
			default:
				extUsages = append(extUsages, fmt.Sprintf("Unknown (%v)", usage))
			}
		}
		fmt.Printf("Extended Key Usage:    %s\n", strings.Join(extUsages, ", "))
	}
	
	// Fingerprints
	fmt.Printf("SHA-1 Fingerprint:     %s\n", formatFingerprint(cert.Raw, "sha1"))
	fmt.Printf("SHA-256 Fingerprint:   %s\n", formatFingerprint(cert.Raw, "sha256"))
	
	// Certificate Authority
	if cert.IsCA {
		fmt.Printf("Certificate Authority: Yes\n")
		if cert.MaxPathLen >= 0 {
			fmt.Printf("Max Path Length:       %d\n", cert.MaxPathLen)
		} else if cert.MaxPathLenZero {
			fmt.Printf("Max Path Length:       0\n")
		}
	} else {
		fmt.Printf("Certificate Authority: No\n")
	}
}

func formatFingerprint(certRaw []byte, algorithm string) string {
    var hash []byte
    switch algorithm {
    case "sha1":
        h := sha1.Sum(certRaw)
        hash = h[:]
    case "sha256":
        h := sha256.Sum256(certRaw)
        hash = h[:]
    default:
        return "Unknown algorithm"
    }

    hexStr := hex.EncodeToString(hash)
    // Format as XX:XX:XX...
    var formatted []string
    for i := 0; i < len(hexStr); i += 2 {
        formatted = append(formatted, strings.ToUpper(hexStr[i:i+2]))
    }
    return strings.Join(formatted, ":")
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
