# Table Of Contents
- [v1.0.12](#v1012)
- [v1.0.11](#v1011)
- [v1.0.11-b2](#v1011-b2)
- [v1.0.11-b1](#v1011-b1)
- [v1.0.10](#v1010)
- [v1.0.9](#v109)
- [v1.0.8](#v108)
- [v1.0.7](#v107)
- [v1.0.6](#v106)
- [v1.0.5](#v105)
- [v1.0.4](#v104)
- [v1.0.3](#v103)
- [v1.0.2](#v102)
- [v1.0.1](#v101)

# v1.0.12

* Add STARTTLS downgrade protection. A MITM attacker can strip STARTTLS
from the server's EHLO response, causing the client to send credentials in
plaintext. The bug is not in `mailsend-go` itself, rather it is in the
[gomail.v2](https://gopkg.in/gomail.v2).  The fix is in my fork of [gomail](https://github.com/muquit/gomail) via the new `RequireSTARTTLS` 
field on the `Dialer`. `mailsend-go` sets it when credentials are provided 
and SSL is not in use, so the connection returns an error instead of 
falling back to plaintext.

* Remove unconditional debug print that leaked SMTP password to stdout.

* A warning is now printed to stderr when TLS certificate verification is
disabled. Use `-verifyCert` if desired.

(May-16-2026)

# v1.0.11

* Fixed duplicate charset in Content-Type header (Issue #73).  Resolved 
issue where http.DetectContentType() was adding charset=utf-8 to the MIME 
type, causing duplicate charset parameters in headers.  The MIME type 
detection code now strips charset params, letting gomail handle charset 
assignment based on the `-cs` flag. Added more test scripts in test/ dir.

(Jan-16-2026)


* Update document about [oauth-helper](https://github.com/muquit/oauth-helper). Switch to [markdown-toc-go](https://github.com/muquit/markdown-toc-go)
from [markdown_helper](https://github.com/BurdetteLamar/markdown_helper) for document generation. Add [oauth-helper](https://github.com/muquit/oauth-helper) link at the end of usage.

(Apr-01-2026)

# v1.0.11-b2

* `ServerName` was missing from `tls.Config`, needed for `-verifyCert` option.
Bug #71.

* Added flag `-printCerts` to print details certificate chain during
displaying  SMTP information with `-info` for SSL or StartTLS. The code is
in `cert_info.go`, created with with Assistance from Claude AI Sonnet 4,
much better than my original cert info. My version was also wrong, it was 
printing a intermediate cert as server cert.

* Use build-time version injection via ldflags instead of hardcoded 
version in main.go using -X ldflags and VERSION file.

(Sep-26-2025)

# v1.0.11-b1
* Initial support to send mail via XOAUTH2. The flags are:
```
   auth -oauth2 -token "access token"
```
Or specify the Access token with environament variable **SMTP_OAUTH_TOKEN**

Note: mailsend-go itself does not do any OAUTH flow. It just needs the 
OAUTH2 access token. You've to get it from your SMTP email provider and use it 
with mailsend-go to send mail.

Please look at the following projects on how an Access Token is Obtained - 
Instructions are for Google and gmail at this time.  But it is similar for 
other providers.

* [smtp-oauth-setup-guide](https://github.com/muquit/smtp-oauth-setup-guide)
* [oauth-helper](https://github.com/muquit/oauth-helper)

(Aug-23-2025)

* Since gomail.v2 is no longer maintained, I forked it to 
https://github.com/muquit/gomail. The main purpose of this fork is to add XOAUTH2 support 
Bug #68)

* Initialize EHLO domian to localhost for smtp info

(Feb-14-2025)

# v1.0.10
* Add flag -use <mail provider> to specify default values for -smtp, -port and
-ssl for well known mail providers. This works for gmail, yahoo, outlook, 
gmx, zoho and aol. Thanks to Nikolas Garofil for pull request.

* Add Docket build info. Fix typos in ChangeLog.md.
  Thanks to Nikolas Garofil for pull request.

* Port is supposed to be optional with default value 587 but was required. 
  Fix Issue #33.

(Dec-06-2020)

# v1.0.9
* The implementation of -name for attachment name was missing.

Fix Issue #26

(Apr-08-2020)

# v1.0.8
* One line message can have embedded new line with \n. If \n is found, it will
be replaced with real new line. Example: 
  body -msg "This is line1.\nThis is line2." 
The message will look like:
```
This is line1.
This is line2.
```

Fix Issue #22

(Mar-17-2020)

# v1.0.7
* If -q was used with -info, the messages were still printed on stdout.

Fix Issue #19

(Feb-16-2020)

# v1.0.6
* Add the flag -cs charset to specify a character set for text or HTML.
The default character set is utf-8.

Fix Issue #12

(Oct-27-2019)

# v1.0.5
* Add the flag -log filePath to write log messages to this path.

Fix Issue #5

(Jul-06-2019)

# v1.0.4

* The To address specified with -t was ignored when a list file was specified with -list. 

* Start a line with # to specify comment in the address list file.

Fix Issue #6

* Add binaries for Raspberry pi. Tested on Raspberry pi 3 b+

* Add the flag -rt to specify reply-to header. Thanks to Dominik G.

(Mar-26-2019)

# v1.0.3

* Code for comma separated to, cc and bcc was not there.
Fix Issue #2

* Remove -tname option for now. It creates trouble if multiple recipients are specified.

(Feb-20-2019)

# v1.0.2

* Supply compiled binary for 32 bit Windows. No code change.

(Feb-14-2019)

# v1.0.1

* Released

(Feb-20-2019)

