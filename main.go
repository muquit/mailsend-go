/*
mailsend-go is a command line program to send mail via SMTP. It is the
golang implementation of the origical C mailsend. I wrote it because I was
getting tired of maintaining versions of Unix and Windows. In go, the binary for
all supported platforms can be cross compiled from one platform, in my case it
is MacOS.

License is MIT

Copyright © 2018 muquit@muquit.com

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the "Software"),
to deal in the Software without restriction, including without limitation
the rights to use, copy, modify, merge, publish, distribute, sublicense,
and/or sell copies of the Software, and to permit persons to whom the
Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included
in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package main

import (
	"bufio"
	"crypto/tls"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	gomail "gopkg.in/gomail.v2"
)

const (
	version = "1.0.10"
)

var (
	debug = false
)

// Attachment ...
type Attachment struct {
	FilePath       string
	AttachmentName string
	MimeType       string
	EncodingType   string
	Inline         bool
	Name           string
}

// Body ...
type Body struct {
	content      string
	mimeType     string
	EncodingType string
	Disposition  string
	CharacterSet string
}

// Auth ...
type Auth struct {
	Username string
	Password string
}

// Header ...
// muquit@muquit.com - - August-18-2018 19:53:31
// /////////////////////////////////////////////////////////////////////////////
type Header struct {
	name  string
	value string
}

// Options are the global flags for mailsend-go
// Use super simple validation:
// tag is "validate" followed by comma separated fields.
// for string type, fields are: string,required/optional,default,flag
// for numeric type, fields are: number,required/optional,default,min,max,flag
// for boolean type, fields are:  boolean,required/optional,default,flag
type Options struct {
	Copyright                bool
	Ipv4                     bool
	Ipv6                     bool
	Info                     bool
	SMTPServer               string `validate:"string,required,N/A,-smtp"`
	Port                     int    `validate:"number,optional,587,1,65535,-port"`
	Domain                   string `validate:"string,optional,localhost,-domain"`
	Subject                  string
	FromName                 string
	From                     string `validate:"string,required,N/A,-from"`
	ToName                   string
	To                       string `validate:"string,required,N/A,-to"`
	Cc                       string
	Bcc                      string
	MessageBody              string
	Name                     string
	ReplyToAddress           string
	RequestReadReciptAddress string
	ReturnPathAddress        string
	Ssl                      bool
	Quiet                    bool
	LogfilePath              string
	VerifyCert               bool
	PrintSMTPInfo            bool
	CharacterSet             string
}

// AddressList ...
type AddressList struct {
	name    string
	address string
}

// Mailsend is the main struct
type Mailsend struct {
	options     Options
	auth        Auth
	body        Body
	headers     []Header
	attachments []Attachment
	addressList []AddressList
}

var (
	mailsend = Mailsend{}
)

// NewAttachment ...
// muquit@muquit.com - - August-18-2018 19:33:38
// /////////////////////////////////////////////////////////////////////
func NewAttachment() *Attachment {
	attachment := new(Attachment)
	attachment.EncodingType = "base64"
	return attachment
}

// NewHeader ...
func NewHeader() *Header {
	header := new(Header)
	return header
}

// NewAuth ...
func NewAuth() *Auth {
	auth := new(Auth)
	return auth
}

// NewBody ...
func NewBody() *Body {
	body := new(Body)
	return body
}

// NewAddressList ...
func NewAddressList() *AddressList {
	addresslist := new(AddressList)
	return addresslist
}

// DefaultValidator ...
type DefaultValidator struct {
}

// StringValidator ...
type StringValidator struct {
	Required bool
	Default  string
	Flag     string
}

// NumberValidator performs numerical value validation.
type NumberValidator struct {
	Required bool
	Default  int
	Min      int
	Max      int
	Flag     string
}

// return true if file exists, false otherwise
func fileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

// Validator interface
type Validator interface {
	Validate(interface{}) (bool, error)
}

// Validate always return true and no error
func (v DefaultValidator) Validate(val interface{}) (bool, error) {
	return true, nil
}

func logDebug(format string, a ...interface{}) {
	if debug {
		log.Printf(format, a...)
	}
}

func logInfo(format string, a ...interface{}) {
	log.Printf(format, a...)
}

func logFile(format string, a ...interface{}) {
	if len(mailsend.options.LogfilePath) > 0 {
		log.Printf(format, a...)
	}
}

func fatalError(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Fprintf(os.Stderr, "ERROR: %s\n", msg)
	log.Printf("ERROR: %s\n", msg)
	os.Exit(1)
}

// Validate numeric value
func (v NumberValidator) Validate(val interface{}) (bool, error) {
	logDebug("in Number validator default: %d\n", v.Default)
	num := val.(int)
	logDebug("num: %d\n", num)
	if num == 0 { // not specified
		num = v.Default
	}

	if num < v.Min {
		return false, fmt.Errorf("%d must be greater than %v", num, v.Min)
	}

	if v.Max >= v.Min && num > v.Max {
		return false, fmt.Errorf("%d must be less than %v", num, v.Max)
	}

	return true, nil
}

// Validate string value
func (v StringValidator) Validate(val interface{}) (bool, error) {
	s := val.(string)
	l := len(s)
	if l == 0 {
		if v.Default != "N/A" {
			l = len(v.Default)
		}
	}

	if l == 0 {
		return false, fmt.Errorf("is required, specify with flag: %s", v.Flag)
	}

	return true, nil
}

// Returns validator struct corresponding to validation type
func getValidator(tag string) Validator {
	args := strings.Split(tag, ",")
	logDebug("args length: %d\n", len(args))
	switch args[0] {
	case "string":
		validator := StringValidator{}
		if args[1] == "required" {
			validator.Required = true
		}
		validator.Default = args[2]
		validator.Flag = args[3]
		return validator

	case "number":
		validator := NumberValidator{}
		if args[1] == "required" {
			validator.Required = true
		}
		logDebug("Flag: %s\n", args[5])
		fmt.Sscanf(strings.Join(args[2:], ","), "%d,%d,%d", &validator.Default, &validator.Min, &validator.Max)
		if args[5] == "-port" {
			if mailsend.options.Port == 0 {
				// Issue #33
				mailsend.options.Port = validator.Default
			}
		}
		return validator
	}

	return DefaultValidator{}
}

// Validate required members of structs using reflection
// muquit@muquit.com - - August-20-2018 20:38:12
// /////////////////////////////////////////////////////////////////////////////
func validateGlobalFlags() []error {
	errs := []error{}
	// validate global options
	options := mailsend.options
	v := reflect.ValueOf(options)
	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get("validate")
		if tag == "" {
			continue
		}
		logDebug("Tag: %s\n", tag)
		validator := getValidator(tag)
		// Perform validation
		valid, err := validator.Validate(v.Field(i).Interface())

		// Append error to results
		if !valid && err != nil {
			errs = append(errs, fmt.Errorf("%s %s", v.Type().Field(i).Name, err.Error()))
		}
	}
	return errs
}

func showExamplesAndExit() {
	/*
		exFile := "examples.txt"
			txt, err := box.FindString(exFile)
			if err != nil {
				fatalError("Could not open file %s\n", exFile)
			}
			fmt.Printf("%s\n", txt)
	*/
	PrintExamples()
	os.Exit(0)
}

func foundAnotherCommand(arg string) bool {
	if arg == "attach" || arg == "oneline" || arg == "body" || arg == "auth" || arg == "header" { // another command found
		return true
	}
	return false
}

func parseHeaderCommandParams(args []string, command string) int {
	argc := len(args)
	h := NewHeader()
	j := 1
	max := 4 // update if new options are added
	for i := 1; i < argc; i++ {
		arg := args[i]
		showHelp(arg)
		if foundAnotherCommand(arg) {
			break
		}
		if i > max {
			break
		}
		if arg == "-name" || arg == "--name" {
			i++
			if i == argc {
				fatalError("Missing value with %s for command %s\n", arg, command)
			}
			h.name = args[i]
			j = i
		}
		if arg == "-value" || arg == "--value" {
			i++
			if i == argc {
				fatalError("Missing value with %s for command %s\n", arg, command)
			}
			h.value = args[i]
			j = i
		}
	}
	if h.name == "" {
		fatalError("Missing header name for command %s\n", command)
	}
	if h.value == "" {
		fatalError("Missing header value for command %s\n", command)
	}
	mailsend.headers = append(mailsend.headers, *h)
	if j > max {
		j = max
	}
	return j
}

// Parse all the valid flags of attach command
// muquit@muquit.com - - August-18-2018 19:36:47
// /////////////////////////////////////////////////////////////////////////////
func parseAttachCommandParams(args []string, command string) int {
	argc := len(args)
	a := NewAttachment()
	j := 1
	max := 6 // update if new options are added
	for i := 1; i < argc; i++ {
		arg := args[i]
		showHelp(arg)
		if foundAnotherCommand(arg) {
			break
		}
		if i > max {
			break
		}
		if arg == "-file" || arg == "--file" {
			i++
			if i == argc {
				fatalError("Missing value with %s for command %s\n", arg, command)
			}
			if !fileExists(args[i]) {
				fatalError("Attachment file %s does not exist\n", args[i])
			}
			a.FilePath = args[i]
			j = i
		}
		if arg == "-mime-type" || arg == "--mime-type" {
			i++
			if i == argc {
				fatalError("Missing value with %s for command %s\n", arg, command)
			}
			a.MimeType = args[i]
			j = i
		}
		if arg == "-inline" || arg == "--inline" {
			a.Inline = true
			j = i
		}
		if arg == "-name" || arg == "--name" {
			i++
			if i == argc {
				fatalError("Missing value with %s for command %s\n", arg, command)
			}
			a.Name = args[i]
			logDebug("Name: " + a.Name)
			j = i
		}
	}
	if len(a.FilePath) == 0 {
		fatalError("No file specified with -file for for command %s\n", command)
	}
	mailsend.attachments = append(mailsend.attachments, *a)
	if j > max {
		j = max
	}
	logDebug("> Encoding Base64 %T\n", gomail.Base64)
	return j
}

func showHelp(arg string) {
	if arg == "-h" || arg == "-help" || arg == "--h" || arg == "--help" {
		showUsageAndExit()
	}
}

func parseAuthCommandParams(args []string, command string) int {
	argc := len(args)
	j := 1
	max := 4
	auth := NewAuth()
	for i := 1; i < argc; i++ {
		arg := args[i]
		if foundAnotherCommand(arg) {
			break
		}
		showHelp(arg)
		if i > max {
			break
		}

		if arg == "-user" {
			i++
			if i == argc {
				fatalError("Missing value with %s for command %s\n", arg, command)
			}
			auth.Username = args[i]
			j = i
		}
		if arg == "-pass" {
			i++
			if i == argc {
				fatalError("Missing value with %s for command %s\n", arg, command)
			}
			auth.Password = args[i]
			j = i
		}
	}
	if len(auth.Username) == 0 {
		fatalError("No auth username specified with -user with command %s\n", command)
	}
	if len(auth.Password) == 0 {
		evar := "SMTP_USER_PASS"
		pass, ok := os.LookupEnv(evar)
		if !ok {
			fatalError("No auth password specified with -pass or env variable %s for command %s\n", evar, command)
		}
		auth.Password = pass
	}
	mailsend.auth = *auth
	if j > max {
		j = max
	}
	return j
}

func parseBodyCommandParams(args []string, command string) int {
	argc := len(args)
	j := 1
	max := 6
	body := NewBody()
	for i := 1; i < argc; i++ {
		arg := args[i]
		showHelp(arg)
		if foundAnotherCommand(arg) {
			break
		}
		if i > max {
			break
		}
		if arg == "-file" {
			i++
			if i == argc {
				fatalError("Missing value with %s for command %s\n", arg, command)
			}
			if !fileExists(args[i]) {
				fatalError("File %s to add a mail body does not exist\n", args[i])
			}
			content := readFile(args[i])
			body.content = string(content)
			j = i
		} else if arg == "-m" || arg == "-msg" || arg == "-message" || arg == "--m" || arg == "--msg" || arg == "--message" {
			i++
			if i == argc {
				fatalError("Missing value with %s for command %s\n", arg, command)
			}
			body.content = args[i]
			j = i
		} else if arg == "-mime-type" || arg == "--mime-type" {
			i++
			if i == argc {
				fatalError("Missing value with %s for command %s\n", arg, command)
			}
			body.mimeType = args[i]
			j = i
		}
	}
	if len(body.content) == 0 {
		fatalError("Path of a text file or a message must be specified with -file or -msg for command %s\n", command)
	}
	if len(body.mimeType) == 0 {
		logDebug("Detecting MIME Type....\n")
		body.mimeType = contentType([]byte(body.content))
	}
	mailsend.body = *body
	if j > max {
		j = max
	}
	return j
}

func showUsageAndExit() {
	v := " Version: @($) mailsend-go v" + version
	usage := ` mailsend-go [options]
  Where the options are:
  -debug                 - Print debug messages
  -sub subject           - Subject
  -t to,to..*            - email address/es of the recipient/s. Required
  -list file             - file with list of email addresses. 
                           Syntax is: Name, email_address
  -fname name            - name of sender
  -f address*            - email address of the sender. Required
  -cc cc,cc..            - carbon copy addresses
  -bcc bcc,bcc..         - blind carbon copy addresses
  -rt rt                 - reply to address
  -smtp host/IP          - hostname/IP address of the SMTP server. Required
                           unless '-use' is set.
  -use mailprovider      - Arranges -smtp, -port and -ssl for you when using
                           a well known mailprovider. Allowed values:
                           gmail, yahoo, outlook, office365, gmx, zoho, aol
  -port port             - port of SMTP server. Default is 587
  -domain domain         - domain name for SMTP HELO. Default is localhost
  -info                  - Print info about SMTP server
  -ssl                   - SMTP over SSL. Default is StartTLS
  -verifyCert            - Verify Certificate in connection. Default is No
  -ex                    - show examples
  -help                  - show this help
  -q                     - quiet
  -log filePath          - write log messages to this file
  -cs charset            - Character set for text/HTML. Default is utf-8
  -V                     - show version and exit
  auth                   - Auth Command
   -user username*       - username for ESMTP authentication. Required
   -pass password*       - password for EMSPTP authentication. Required
  body                   - body command for attachment for mail body
   -msg msg              - message to show as body 
   -file path            - or path of a text/HTML file
   -mime-type type       - MIME type of the body content. Default is detected
  attach                 - attach command. Repeat for multiple attachments
   -file path*           - path of the attachment. Required
   -name name            - name of the attachment. Default is filename
   -mime-type type       - MIME-Type of the attachment. Default is detected
   -inline               - Set Content-Disposition to "inline". 
                           Default is "attachment"
  header                 - Header Command. Repeat for multiple headers
   -name header          - Header name
   -value value          - Header value

The options with * are required. 

Environment variables:
   SMTP_USER_PASS for auth password (-pass)
`

	usage = strings.Replace(usage, "\t", "    ", -1)
	fmt.Printf("%s\n\n%s\n", v, usage)
	os.Exit(0)
}

// create slice of addresses from comma separated address
func makeRecipientAddresses(to string) []string {
	var addresses []string
	addrs := strings.Split(to, ",")
	for i := range addrs {
		addr := addrs[i]
		if len(addr) > 0 {
			// remove leading and trailing spaces
			addresses = append(addresses, strings.TrimSpace(addr))
		}
	}
	return addresses
}

func constructMail(fromName string, fromAddress string, toName string, toAddress string) *gomail.Message {
	o := mailsend.options
	m := gomail.NewMessage()

	if len(o.CharacterSet) > 0 {
		m = gomail.NewMessage(gomail.SetCharset(o.CharacterSet))
	}

	if len(fromName) > 0 {
		m.SetAddressHeader("From", fromAddress, fromName)
	} else {
		m.SetHeader("From", fromAddress)
	}

	if len(toName) > 0 {
		m.SetAddressHeader("To", toAddress, toName) // Off for now
	} else {
		recipients := makeRecipientAddresses(toAddress)
		addresses := make([]string, len(recipients))
		for i, to := range recipients {
			logDebug(" To: %s\n", to)
			addresses[i] = m.FormatAddress(to, "")
		}
		// Issue #2
		m.SetHeader("To", addresses...)
	}

	if len(o.Cc) > 0 {
		logDebug("Setting Carbon Copy: %s\n", o.Cc)
		recipients := makeRecipientAddresses(o.Cc)
		addresses := make([]string, len(recipients))
		for i, cc := range recipients {
			logDebug(" Cc: %s\n", cc)
			addresses[i] = m.FormatAddress(cc, "")
		}
		// Issue #2
		m.SetHeader("Cc", addresses...)
	}

	if len(o.Bcc) > 0 {
		logDebug("Setting Bind Carbon Copy: %s\n", o.Bcc)
		recipients := makeRecipientAddresses(o.Bcc)
		addresses := make([]string, len(recipients))
		for i, bcc := range recipients {
			logDebug(" Bcc: %s\n", bcc)
			addresses[i] = m.FormatAddress(bcc, "")
		}
		// Issue #2
		m.SetHeader("Bcc", addresses...)
	}

	if len(o.ReplyToAddress) > 0 {
		logDebug("Setting ReplyTo: %s\n", o.ReplyToAddress)
		m.SetHeader("reply-to", m.FormatAddress(o.ReplyToAddress, ""))
	}

	m.SetHeader("Subject", o.Subject)
	xmailer := fmt.Sprintf(" @(#) mailsend-go v%s, %s", version, runtime.GOOS)
	m.SetHeader("X-Mailer", xmailer)
	m.SetHeader("X-Copyright", "MIT. It is illegal to use this software for Spamming")

	// set custom headers if specified
	for _, h := range mailsend.headers {
		m.SetHeader(h.name, h.value)
	}

	if len(mailsend.body.content) > 0 {
		logDebug("Attach body\n")
		// Replace \n with real new line. Issue #22
		msg := strings.Replace(mailsend.body.content, `\n`, "\n", -1)
		m.SetBody(mailsend.body.mimeType, msg)
	}

	for _, a := range mailsend.attachments {
		if len(a.MimeType) > 0 {
			logDebug("Setting MIME-TYPE of the message to: %s\n", a.MimeType)
			mtype := map[string][]string{"Content-Type": {a.MimeType}}
			if !a.Inline {
				logDebug("Disposition is attach\n")
				if len(a.Name) > 0 {
					m.Attach(a.FilePath, gomail.SetHeader(mtype), gomail.Rename(a.Name))
				} else {
					m.Attach(a.FilePath, gomail.SetHeader(mtype))
				}
			} else {
				logDebug("Disposition is inline\n")
				m.Embed(a.FilePath, gomail.SetHeader(mtype))
			}
		} else {
			if !a.Inline {
				logDebug("Attach: %s\n", a.FilePath)
				if len(a.Name) > 0 {
					logDebug("Name: %s\n", a.Name)
					m.Attach(a.FilePath, gomail.Rename(a.Name))
				} else {
					m.Attach(a.FilePath)
				}
			} else {
				logDebug("Inline: %s\n", a.FilePath)
				m.Embed(a.FilePath)
			}
		}
	}

	return m
}

func sendMail() {
	logFile("- mailsend-go v%s starts -\n", version)
	o := mailsend.options
	logDebug("Subject: %s\n", o.Subject)
	logDebug("From: %s\n", o.From)
	logDebug("To: %s\n", o.To)

	logFile("Subject: %s\n", o.Subject)
	logFile("From: %s\n", o.From)
	logFile("To: %s\n", o.To)
	//	logDebug("To Name: %s\n", o.ToName)
	logDebug("SMTP server: %s\n", o.SMTPServer)
	logDebug("SMTP Port: %d\n", o.Port)
	logDebug("Setting From with name: %s,%s\n", o.From, o.FromName)

	logFile("SMTP server: %s\n", o.SMTPServer)
	logFile("SMTP Port: %d\n", o.Port)
	logFile("Setting From with name: %s,%s\n", o.From, o.FromName)

	var d *gomail.Dialer
	if mailsend.auth.Username != "" && mailsend.auth.Password != "" {
		logDebug("Using ESMTP Authentication")
		d = gomail.NewDialer(o.SMTPServer, o.Port, mailsend.auth.Username, mailsend.auth.Password)
	} else {
		logDebug("Not Using ESMTP Authentication")
		d = &gomail.Dialer{Host: o.SMTPServer, Port: o.Port}
	}

	// default is localhost
	d.LocalName = mailsend.options.Domain

	if mailsend.options.Ssl {
		d.SSL = true
	}
	logDebug("SSL? %t\n", d.SSL)
	if d.SSL {
		// always skip verification, it segfaults if the host is an IP address
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	} else {
		d.TLSConfig = &tls.Config{InsecureSkipVerify: !mailsend.options.VerifyCert}
	}
	s, err := d.Dial()
	if err != nil {
		fatalError("%s\n", err)
	}
	logDebug("Sending mail...")
	logFile("Sending mail...")

	// send mail to a list of users
	if len(mailsend.addressList) > 0 {
		// Issue #6
		// Send to To first
		m := constructMail(o.FromName, o.From, o.ToName, o.To)
		if err := gomail.Send(s, m); err != nil {
			log.Printf("ERROR: Could not send mail to %q: %v\n", o.To, err)
		}
		m.Reset()
		for _, r := range mailsend.addressList {
			m := constructMail(o.FromName, o.From, r.name, r.address)
			if err := gomail.Send(s, m); err != nil {
				log.Printf("ERROR: Could not send mail to %q: %v\n", r.address, err)
			}
			m.Reset()
		}
	} else {
		m := constructMail(o.FromName, o.From, o.ToName, o.To)
		if err := gomail.Send(s, m); err != nil {
			fatalError("%s\n", err)
		}

	}

	if !mailsend.options.Quiet {
		fmt.Printf("Mail Sent Successfully\n")
	}
	logFile("Mail Sent Successfully\n")

	logFile("- mailsend-go v%s ends -\n", version)
}

// return content of the file as string
// exit on error
func readFile(path string) []byte {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fatalError("Could not read file %s:%s\n", path, err)
	}
	return b
}

func contentType(content []byte) string {
	return http.DetectContentType(content)
}

func xprintSMTPInfo() {
	if mailsend.options.SMTPServer == "" {
		fatalError("Please specify SMTP server with flag -smtp or set it indirectly with -use")
	}
	if mailsend.options.Port == 0 {
		mailsend.options.Port = 587
	}
	logDebug("SMTP Server: %s:%d\n", mailsend.options.SMTPServer, mailsend.options.Port)
	logDebug("Domain: %s\n",mailsend.options.Domain)
	printSMTPInfo(mailsend.options.SMTPServer, mailsend.options.Port, mailsend.options.Domain, mailsend.options.Ssl, mailsend.options.VerifyCert)
}

// Address list file a comma separated Name, Address lines
func parseAddressListFile(listFile string) {
	csvFile, err := os.Open(listFile)
	if err != nil {
		fatalError("Could not open address list file %s", listFile)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			fatalError("Error parsing address list CSV file: %s", error)
		}
		// If line starts with # ignore. Issue #6
		comment := strings.HasPrefix(line[0], "#")
		if comment {
			continue
		}
		al := NewAddressList()
		al.name = line[0]
		al.address = strings.TrimSpace(line[1])
		mailsend.addressList = append(mailsend.addressList, *al)
	}
}

func openLogfile() {
	path := mailsend.options.LogfilePath
	if len(path) > 0 {
		f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			fatalError("Could not open log file %s for writing: %v", path, err)
		}
		log.SetOutput(f)
	}
}

func main() {
	args := os.Args
	if len(args) == 0 {
		showUsageAndExit()
	}
	args = args[1:]
	argc := len(args)
	for i := 0; i < argc; i++ {
		arg := args[i]
		if arg == "-debug" || arg == "--debug" {
			debug = true
		} else if arg == "-h" || arg == "-help" || arg == "--h" || arg == "--help" {
			showUsageAndExit()
		} else if arg == "-domain" || arg == "--domain" {
			i++
			if i == argc {
				fatalError("Missing value for %s\n", arg)
			}
			mailsend.options.Domain = args[i]
		} else if arg == "-ex" || arg == "--ex" || arg == "-example" || arg == "--example" {
			showExamplesAndExit()
		} else if arg == "-t" || arg == "-to" || arg == "--t" || arg == "--to" {
			i++
			if i == argc {
				fatalError("Missing value for %s\n", arg)
			}
			mailsend.options.To = args[i]
		} else if arg == "-cc" || arg == "--cc" {
			i++
			if i == argc {
				fatalError("Missing value for %s\n", arg)
			}
			mailsend.options.Cc = args[i]
		} else if arg == "-bcc" || arg == "--bcc" {
			i++
			if i == argc {
				fatalError("Missing value for %s\n", arg)
			}
			mailsend.options.Bcc = args[i]
		} else if arg == "-rt" || arg == "--rt" {
			i++
			if i == argc {
				fatalError("Missing value for %s\n", arg)
			}
			mailsend.options.ReplyToAddress = args[i]
		} else if arg == "-f" || arg == "-from" || arg == "--f" || arg == "--from" {
			i++
			if i == argc {
				fatalError("Missing value for %s\n", arg)
			}
			mailsend.options.From = args[i]
		} else if arg == "-fname" || arg == "--fname" {
			i++
			if i == argc {
				fatalError("Missing value for %s\n", arg)
			}
			mailsend.options.FromName = args[i]
		} else if arg == "-sub" || arg == "-subject" || arg == "--sub" || arg == "--subject" {
			i++
			if i == argc {
				fatalError("Missing value for %s\n", arg)
			}
			mailsend.options.Subject = args[i]
		} else if arg == "-use" || arg == "--use" {
			i++
			if i == argc {
				fatalError("Missing value for %s\n", arg)
			}
			if args[i] == "gmail" {
				mailsend.options.SMTPServer = "smtp.gmail.com"
				mailsend.options.Port = 587
			} else if args[i] == "yahoo" {
				mailsend.options.SMTPServer = "smtp.mail.yahoo.com"
				mailsend.options.Port = 465
				mailsend.options.Ssl = true
			} else if args[i] == "office365" {
				mailsend.options.SMTPServer = "smtp.office365.com"
				mailsend.options.Port = 587
			} else if args[i] == "outlook" {
				mailsend.options.SMTPServer = "smtp-mail.outlook.com"
				mailsend.options.Port = 587
			} else if args[i] == "gmx" {
				mailsend.options.SMTPServer = "smtp.gmx.com"
				mailsend.options.Port = 465
				mailsend.options.Ssl = true
			} else if args[i] == "zoho" {
				mailsend.options.SMTPServer = "smtp.zoho.com"
				mailsend.options.Port = 465
				mailsend.options.Ssl = true
			} else if args[i] == "aol" {
				mailsend.options.SMTPServer = "smtp.aol.com"
				mailsend.options.Port = 587
			} else {
				fatalError("Mail provider '%s' not known\n", args[i])
			}
		} else if arg == "-smtp" || arg == "--smtp" {
			i++
			if i == argc {
				fatalError("Missing value for %s\n", arg)
			}
			mailsend.options.SMTPServer = args[i]
		} else if arg == "-p" || arg == "-port" || arg == "--p" || arg == "--port" {
			i++
			if i == argc {
				fatalError("Missing value for %s\n", arg)
			}
			port, err := strconv.Atoi(args[i])
			mailsend.options.Port = port
			if err != nil {
				fatalError("Invalid Port %s specified with %s\n", args[i], arg)
			}
		} else if arg == "-list" || arg == "--list" {
			i++
			if i == argc {
				fatalError("Missing value for %s\n", arg)
			}
			if !fileExists(args[i]) {
				fatalError("List file %s does not exist\n", args[i])
			}
			parseAddressListFile(args[i])
			for _, al := range mailsend.addressList {
				fmt.Printf("Name: '%s', Email: '%s'\n", al.name, al.address)
			}
		} else if arg == "-log" || arg == "--log" {
			i++
			if i == argc {
				fatalError("Missing value for %s\n", arg)
			}
			mailsend.options.LogfilePath = args[i]
		} else if arg == "-cs" || arg == "--cs" {
			i++
			if i == argc {
				fatalError("Missing value for %s\n", arg)
			}
			mailsend.options.CharacterSet = args[i]
		} else if arg == "-V" || arg == "--V" {
			fmt.Printf("@(#) mailsend-go v%s\n", version)
			os.Exit(0)
		} else if arg == "-info" || arg == "--info" {
			mailsend.options.PrintSMTPInfo = true
		} else if arg == "-ssl" || arg == "--ssl" {
			mailsend.options.Ssl = true
		} else if arg == "-verifyCert" || arg == "--verifyCert" {
			mailsend.options.VerifyCert = true
		} else if arg == "-q" || arg == "-quiet" || arg == "--q" || arg == "--quiet" {
			mailsend.options.Quiet = true
		} else if arg == "body" {
			j := parseBodyCommandParams(args[i:], arg)
			i += j
		} else if arg == "attach" {
			j := parseAttachCommandParams(args[i:], arg)
			i += j
		} else if arg == "auth" {
			j := parseAuthCommandParams(args[i:], arg)
			i += j
		} else if arg == "header" {
			j := parseHeaderCommandParams(args[i:], arg)
			i += j
		} else {
			fatalError("Unknown option %s\n", arg)
		}
	}
	if mailsend.options.PrintSMTPInfo {
		xprintSMTPInfo()
		os.Exit(0)
	}

	logDebug("Number of attachments: %d\n", len(mailsend.attachments))
	for n, attachment := range mailsend.attachments {
		logDebug("%d, File: %s\n", n, attachment.FilePath)
		logDebug("%d, Encoding type: %s\n", n, attachment.EncodingType)
	}

	errors := validateGlobalFlags()
	if len(errors) > 0 {
		fmt.Printf("\nmailsend-go v%s\n\n", version)
		for _, err := range errors {
			fmt.Printf("ERROR: %s\n", err.Error())
		}
		fmt.Printf("\nRun with -h for help\n\n")
		os.Exit(1)
	}
	openLogfile()
	sendMail()
}
