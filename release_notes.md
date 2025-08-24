First cut to support SMTP XOAUTH2. This version is in beta. Please test and
give feedback if it works for you.

`mailsend-go` itself does not do any OAUTH2 flow. It only takes 
the access token to authenticate before sending mail.
You have to obtain the access token externally. 

Please look at the following projects on how an Access Token is Obtained -
Instructions are for Google and gmail at this time.  But it is similar for
other providers.

* [smtp-oauth-setup-guide](https://github.com/muquit/smtp-oauth-setup-guide)
* [oauth-helper](https://github.com/muquit/oauth-helper)

Please look at [ChangeLog](ChangeLog.md) for details on what has changed 
in the current version. 

The binaries are cross-compiled with 
[go-xbuild-go](https://github.com/muquit/go-xbuild-go.)

Do not forget to check checksums of the archives before using.

Thanks.
