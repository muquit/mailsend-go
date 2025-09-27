# Introduction

`mailsend-go` is a command line tool to send mail via SMTP protocol. This is the
[golang](https://golang.org/) incarnation of my C version of
[mailsend](https://github.com/muquit/mailsend/). However, this version is much
simpler and all the heavy lifting is done by the package
[gomail.v2](https://gopkg.in/gomail.v2). However, this package is not maintained anymore. Therefore, I forked it to
[gomail](https://github.com/muquit/gomail) (starting from mailsend-go v1.0.11-b1 Aug-24-2025).
The main purpose of this fork is to add XOAUTH2 support (Bug #68)

If you use [mailsend](https://github.com/muquit/mailsend), please consider
using mailsend-go as no new features will be added to 
[mailsend](https://github.com/muquit/mailsend).

If you have any question, request or suggestion, please enter it in the 
[Issues](https://github.com/muquit/mailsend-go/issues) with appropriate label.

**NOTE:** XOAUTH2 support is available in v1.0.11-b1 (Released on Aug-24-2025)

Please look at [ChangeLog](ChangeLog.md) for what has changed in the current version.
