This release fixes [Issue #76](https://github.com/muquit/mailsend-go/issues/76).
The options `-printCerts` and `-verifyCert` now take effect while sending 
mail, not just with `-info`. Added `-verbose` as a shortcut for
`-debug -printCerts -verifyCert` together.

No need to update unless you need certificate details or verification
while sending mail.

Please look at [ChangeLog](ChangeLog.md) for details.

The binaries are cross-compiled with 
[go-xbuild-go](https://github.com/muquit/go-xbuild-go)

Do not forget to check checksums of the archives before using.

Thanks.
