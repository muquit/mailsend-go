# Examples

Each example mailsend-go command is a single line. In Unix back slash \ 
can be used to continue in the next line. Also in Unix, use single quotes 
instead of double quotes, otherwise if input has any shell character like 
$ etc, it will get expanded by the shell.

## Show SMTP server information

### StartTLS will be used if server supports it

```
  mailsend-go -info -smtp smtp.gmail.com -port 587
```    

```
[S] 220 smtp.gmail.com ESMTP k185-v6sm17739711qkd.27 - gsmtp
[C] HELO localhost
[C] EHLO localhost
[S] 250-smtp.gmail.com at your service, [x.x.x.x]
[S] 250-SIZE 35882577
[S] 250-8BITMIME
[S] 250-STARTTLS
[S] 250-ENHANCEDSTATUSCODES
[S] 250-PIPELINING
[S] 250-CHUNKING
[S] 250-SMTPUTF8
[C] STARTTLS
[S] 220-2.0.0 Ready to start TLS
[C] EHLO localhost
[S] 250-smtp.gmail.com at your service, [x.x.x.x]
[S] 250-SIZE 35882577
[S] 250-8BITMIME
[S] 250-AUTH LOGIN PLAIN XOAUTH2 PLAIN-CLIENTTOKEN OAUTHBEARER XOAUTH
[S] 250-ENHANCEDSTATUSCODES
[S] 250-PIPELINING
[S] 250-CHUNKING
[S] 250-SMTPUTF8
Certificate of smtp.gmail.com:
 Version: 3 (0x3)
 Serial Number: 149685795415515161014990164765 (0x1e3a9301cfc7206383f9a531d)
 Signature Algorithm: SHA256-RSA
 Subject: CN=Google Internet Authority G3,O=Google Trust Services,C=US
 Issuer: GlobalSign
 Not before: 2017-06-15 00:00:42 +0000 UTC
 Not after: 2021-12-15 00:00:42 +0000 UTC
[C] QUIT
[S] 221-2.0.0 closing connection k185-v6sm17739711qkd.27 - gsmtp
```

### Use SSL. Note the port is different

```
  mailsend-go -info -smtp smtp.gmail.com -port 465 -ssl
```

### Print SMTP server certificate chain

```bash
  mailsend-go -info -smtp smtp.gmail.com -port 587 -printCerts
```

```bash
[S] 220 smtp.gmail.com ESMTP 6a1803df08f44-80166e2f32bsm34741116d6.41 - gsmtp
[C] HELO localhost
[C] EHLO localhost
[S] 250-smtp.gmail.com at your service, [x.x.x.x]
[S] 250-SIZE 35882577
[S] 250-8BITMIME
[S] 250-STARTTLS
[S] 250-ENHANCEDSTATUSCODES
[S] 250-PIPELINING
[S] 250-CHUNKING
[S] 250-SMTPUTF8
[C] STARTTLS
[S] 220-2.0.0 Ready to start TLS
[C] EHLO localhost
[S] 250-smtp.gmail.com at your service, [x.x.x.x]
[S] 250-SIZE 35882577
[S] 250-8BITMIME
[S] 250-AUTH LOGIN PLAIN XOAUTH2 PLAIN-CLIENTTOKEN OAUTHBEARER XOAUTH
[S] 250-ENHANCEDSTATUSCODES
[S] 250-PIPELINING
[S] 250-CHUNKING
[S] 250-SMTPUTF8
[C] QUIT
[S] 221-2.0.0 closing connection 6a1803df08f44-80166e2f32bsm34741116d6.41 - gsmtp

=== TLS Connection Information ===
TLS Version: TLS 1.3
Cipher Suite: 0x1301
Server Name: smtp.gmail.com
Negotiated Protocol:

Certificate Chain (3 certificates):

--- Certificate 1 ---
Subject: CN=smtp.gmail.com
Issuer: CN=WR2,O=Google Trust Services,C=US
Serial Number: 14461562026188826353951632455228095006
Not Before: 2025-09-08T08:36:45Z
Not After: 2025-12-01T08:36:44Z
Is CA: false
DNS Names: smtp.gmail.com
Key Usage: Digital Signature
Status: Valid

--- Certificate 2 ---
Subject: CN=WR2,O=Google Trust Services,C=US
Issuer: CN=GTS Root R1,O=Google Trust Services LLC,C=US
Serial Number: 170058220837755766831192027518741805976
Not Before: 2023-12-13T09:00:00Z
Not After: 2029-02-20T14:00:00Z
Is CA: true
Key Usage: Digital Signature, Certificate Sign, CRL Sign
Status: Valid

--- Certificate 3 ---
Subject: CN=GTS Root R1,O=Google Trust Services LLC,C=US
Issuer: CN=GlobalSign Root CA,OU=Root CA,O=GlobalSign nv-sa,C=BE
Serial Number: 159159747900478145820483398898491642637
Not Before: 2020-06-19T00:00:42Z
Not After: 2028-01-28T00:00:42Z
Is CA: true
Key Usage: Digital Signature, Certificate Sign, CRL Sign
Status: Valid

--- Certificate Fingerprints (Leaf) ---
SHA-1: 28:88:45:90:10:20:88:BA:87:2E:0E:7C:3A:12:D6:35:EC:26:AE:90
SHA-256: 6F:F8:E2:F5:D4:AE:5A:FF:92:4A:5F:AC:88:80:14:3A:30:33:7A:CF:EE:33:94:82:EF:2A:93:47:80:E4:18:EF
=====================================
```

```bash
   mailsend-go -info -smtp smtp.gmail.com -port 465 -ssl -printCerts
```

```bash
[S] 220 smtp.gmail.com ESMTP 6a1803df08f44-80166781d27sm35134546d6.45 - gsmtp
[C] HELO localhost
[C] EHLO localhost
[S] 250-smtp.gmail.com at your service, [x.x.x.x]
[S] 250-SIZE 35882577
[S] 250-8BITMIME
[S] 250-AUTH LOGIN PLAIN XOAUTH2 PLAIN-CLIENTTOKEN OAUTHBEARER XOAUTH
[S] 250-ENHANCEDSTATUSCODES
[S] 250-PIPELINING
[S] 250-CHUNKING
[S] 250-SMTPUTF8
[C] QUIT
[S] 221-2.0.0 closing connection 6a1803df08f44-80166781d27sm35134546d6.45 - gsmtp

=== TLS Connection Information ===
TLS Version: TLS 1.3
Cipher Suite: 0x1301
Server Name: smtp.gmail.com
Negotiated Protocol:

Certificate Chain (3 certificates):

--- Certificate 1 ---
Subject: CN=smtp.gmail.com
Issuer: CN=WR2,O=Google Trust Services,C=US
Serial Number: 14461562026188826353951632455228095006
Not Before: 2025-09-08T08:36:45Z
Not After: 2025-12-01T08:36:44Z
Is CA: false
DNS Names: smtp.gmail.com
Key Usage: Digital Signature
Status: Valid

--- Certificate 2 ---
Subject: CN=WR2,O=Google Trust Services,C=US
Issuer: CN=GTS Root R1,O=Google Trust Services LLC,C=US
Serial Number: 170058220837755766831192027518741805976
Not Before: 2023-12-13T09:00:00Z
Not After: 2029-02-20T14:00:00Z
Is CA: true
Key Usage: Digital Signature, Certificate Sign, CRL Sign
Status: Valid

--- Certificate 3 ---
Subject: CN=GTS Root R1,O=Google Trust Services LLC,C=US
Issuer: CN=GlobalSign Root CA,OU=Root CA,O=GlobalSign nv-sa,C=BE
Serial Number: 159159747900478145820483398898491642637
Not Before: 2020-06-19T00:00:42Z
Not After: 2028-01-28T00:00:42Z
Is CA: true
Key Usage: Digital Signature, Certificate Sign, CRL Sign
Status: Valid

--- Certificate Fingerprints (Leaf) ---
SHA-256: 6F:F8:E2:F5:D4:AE:5A:FF:92:4A:5F:AC:88:80:14:3A:30:33:7A:CF:EE:33:94:82:EF:2A:93:47:80:E4:18:EF
SHA-1: 28:88:45:90:10:20:88:BA:87:2E:0E:7C:3A:12:D6:35:EC:26:AE:90
=====================================
```

```bash
    mailsend-go -info -smtp smtp-mail.outlook.com -port 587 -printCerts
```

```bash
[S] 220 MN2PR01CA0065.outlook.office365.com Microsoft ESMTP MAIL Service ready at Sat, 27 Sep 2025 00:29:10 +0000 [08DDFAA3FED0C1E0]
[C] HELO localhost
[C] EHLO localhost
[S] 250-MN2PR01CA0065.outlook.office365.com Hello [x.x.x.x]
[S] 250-SIZE 157286400
[S] 250-PIPELINING
[S] 250-DSN
[S] 250-ENHANCEDSTATUSCODES
[S] 250-STARTTLS
[S] 250-8BITMIME
[S] 250-BINARYMIME
[S] 250-CHUNKING
[S] 250-SMTPUTF8
[C] STARTTLS
[S] 220-2.0.0 SMTP server ready
[C] EHLO localhost
[S] 250-MN2PR01CA0065.outlook.office365.com Hello [x.x.x.x]
[S] 250-SIZE 157286400
[S] 250-PIPELINING
[S] 250-DSN
[S] 250-ENHANCEDSTATUSCODES
[S] 250-AUTH LOGIN XOAUTH2
[S] 250-8BITMIME
[S] 250-BINARYMIME
[S] 250-CHUNKING
[S] 250-SMTPUTF8
[C] QUIT
[S] 221-2.0.0 Service closing transmission channel

=== TLS Connection Information ===
TLS Version: TLS 1.3
Cipher Suite: 0x1302
Server Name: smtp-mail.outlook.com
Negotiated Protocol:

Certificate Chain (2 certificates):

--- Certificate 1 ---
Subject: CN=outlook.com,O=Microsoft Corporation,L=Redmond,ST=Washington,C=US
Issuer: CN=DigiCert Cloud Services CA-1,O=DigiCert Inc,C=US
Serial Number: 10535063011692331098818316272276424549
Not Before: 2025-03-29T00:00:00Z
Not After: 2026-03-28T23:59:59Z
Is CA: false
DNS Names: *.clo.footprintdns.com, *.hotmail.com, *.internal.outlook.com, *.live.com, *.nrb.footprintdns.com, *.office.com, *.office365.com, *.outlook.com, *.outlook.office365.com, attachment.outlook.live.net, attachment.outlook.office.net, attachment.outlook.officeppe.net, attachments.office.net, attachments-sdf.office.net, ccs.login.microsoftonline.com, ccs-sdf.login.microsoftonline.com, hotmail.com, mail.services.live.com, office365.com, outlook.com, outlook.office.com, substrate.office.com, substrate-sdf.office.com
Key Usage: Digital Signature, Key Encipherment
Status: Valid

--- Certificate 2 ---
Subject: CN=DigiCert Cloud Services CA-1,O=DigiCert Inc,C=US
Issuer: CN=DigiCert Global Root CA,OU=www.digicert.com,O=DigiCert Inc,C=US
Serial Number: 20058375873168194746987232153701302504
Not Before: 2020-09-25T00:00:00Z
Not After: 2030-09-24T23:59:59Z
Is CA: true
Key Usage: Digital Signature, Certificate Sign, CRL Sign
Status: Valid

--- Certificate Fingerprints (Leaf) ---
SHA-1: A6:F7:EC:FB:2B:F6:31:B3:A8:4F:EB:B0:9F:FD:BB:4E:3B:0F:42:11
SHA-256: 4F:94:1A:8E:50:52:5E:09:24:4F:8F:FE:75:65:E1:6A:51:DD:10:47:04:74:94:6A:0F:BA:84:6A:86:E4:DE:8C
=====================================
```

### Use default settings for well known mail providers

Don't worry about the settings of -smtp, -port and -ssl for well known mail
providers. This works for gmail, yahoo, outlook, gmx, zoho and aol.

      mailsend-go -info -use gmail

## Send mail with a text message

Notice "auth" is a command and it takes -user and -pass arguments. "body" is
also a command and here it took -msg as an argument. The command "body" can
not repeat, if specified more than once, the last one will be used.

```
    mailsend-go -sub "Test"  -smtp smtp.gmail.com -port 587 \
     auth \
      -user jsnow@gmail.com -pass "secret" \
     -from "jsnow@gmail.com" -to  "mjane@example.com" \
     body \
       -msg "hello, world!\nThis is a message"
```                    
The embedded new line \\n will be converted to a real newline and the final
message will show up as two lines.

The environment variable "SMTP_USER_PASS" can be used instead of the flag
`-pass`.

## Send mail with a HTML message
```
    mailsend-go -sub "Test"  \
    -smtp smtp.gmail.com -port 587 \
    auth \
     -user jsnow@gmail.com -pass "secret" \
    -from "jsnow@gmail.com"  \
    -to  "mjane@example.com" -from "jsnow@gmail.com" \
    body \
     -msg "<b>hello, world!</b>"
```

## Attach a PDF file
MIME type will be detected. Content-Disposition will be set to "attachment",
Content-Transfer-Encoding will be "Base64". Notice, "attach" is a command it
took -file as an arg. The command "attach" can repeat.
```
    mailsend-go -sub "Test"  \
    -smtp smtp.gmail.com -port 587 \
    auth \
     -user jsnow@gmail.com -pass "secret" \
    -from "jsnow@gmail.com"  \
    -to  "mjane@example.com" -from "jsnow@gmail.com" \
    body \
     -msg "A PDF file is attached" \
    attach \
     -file "/path/file.pdf"
```
The name of the attachment will be file.pdf. To change the attachmetn name,
use the `-name` flag. e.g.

```
    attach -file "/path/file.pdf" -name "report.pdf"
```

## Attach a PDF file and an image
Notice, the "attach" command is repeated here.
```
    mailsend-go -sub "Test"  \
    -smtp smtp.gmail.com -port 587 \
    auth \
     -user jsnow@gmail.com -pass "secret" \
    -from "jsnow@gmail.com"  \
    -to  "mjane@example.com" -from "jsnow@gmail.com" \
    body \
     -msg "A PDF file and a PNG file is attached" \
    attach \
     -file "/path/file.pdf" \
    attach \
     -file "/path/file.png"
```
## Attach a PDF file and embed an image
Content-Disposition for the image will be set to "inline". It's an hint to the
mail reader to display the image on the page. Note: it is just a hint, it is
up to the mail reader to respect it or ignore it.
```
    mailsend-go -sub "Test"  \
    -smtp smtp.gmail.com -port 587 \
    auth \
     -user jsnow@gmail.com -pass "secret" \
    -from "jsnow@gmail.com"  \
    -to  "mjane@example.com" -from "jsnow@gmail.com" \
    body \
     -msg "A PDF file is attached, image should be displayed inline" \
    attach \
     -file "/path/file.pdf" \
    attach \
     -file "/path/file.png" \
     -inline
```
## Set Carbon Copy and Blind Carbon copy
```
    mailsend-go -sub "Testing -cc and -bcc" \
    -smtp smtp.gmail.com -port 587 \
    auth \
     -user example@gmail.com -pass "secret" \
     -to jsoe@example.com \
     -f "example@gmail.com" \
     -cc "user1@example.com,user2@example.com" \
     -bcc "foo@example.com" \
     body -msg "Testing Carbon Copy and Blind Carbon copy"
```
Cc addresses will be visible to the recipients but Bcc address will not be.

## Send mail to a list of users

Create a file with list of users. The syntax is ```Name,email_address``` in a line. Name can be empty but comma must be specified. Example of a list file:

```
    # This is a comment.
    # The syntax is Name,email address in a line. Name can be empty but comma 
    # must be specified
    John Snow,jsnow@example.com
    Mary Jane,mjane@example.com
    ,foobar@example.com
```

Specify the list file with ```-list``` flag. 

```
    mailsend-go -sub "Test sending mail to a list of users" \
    -smtp smtp.gmail.com -port 587 \
    auth \
     -user example@gmail.com -pass "secret" \
        -f "me@example.com" \
        -to "xyz@example.com" \
        body \
        -msg "This is a test of sendmail mail to a list of users" \
        attach \
            -file "cat.jpg" \
         attach \
            -file "flower.jpg" \
            -inline \
         -list "list.txt"
```

## Add Custom Headers

Use the command "header" to add custom headers. The command "header" can be
repeated.

```
    mailsend-go -sub "Testing custom headers" \
    -smtp smtp.gmail.com -port 587 \
    auth \
     -user example@gmail.com -pass "secret" \
     -to jdoe@example.com \
     -f "example@gmail.com" \
     body -msg "Testing adding Custom headers"
     header \
         -name "X-MyHeader-1" -value "Value of X-MyHeader-1" \
     header \
         -name "X-MyHeader-2" -value "Value of X-MyHeader-2"

```

## Write logs to a file

Use the flag `-log path_of_log_file.txt`

```
    mailsend-go -sub "test log" \
     -smtp smtp.example.com -port 587 \
     auth \
      -user example@gmail.com -pass "secret" \
      -to jdoe@example.com \
      -f "example@gmail.com" \
      body -msg "Testing log file" \
      -log "/tmp/mailsend-go.log"
```

## Specify a different character set

The default character set is utf-8

```
    mailsend-go -sub "test character set" \
     -smtp smtp.example.com -port 587 \
     auth \
      -user example@gmail.com -pass "secret" \
      -to jdoe@example.com \
      -from "example@gmail.com" \
      -subject "Testing Big5 Charset" \
      -cs "Big5" \
      body -msg "中文測試"

```

---

(Generated from docs/examples.md)

---
