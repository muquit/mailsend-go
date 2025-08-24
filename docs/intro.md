# Introduction

`mailsend-go` is a command line tool to send mail via SMTP protocol. This is the
[golang](https://golang.org/) incarnation of my C version of
[mailsend](https://github.com/muquit/mailsend/). However, this version is much
simpler and all the heavy lifting is done by the package
[gomail.v2](https://gopkg.in/gomail.v2). 

**Note:** this package is not maintained anymore. Therefore, I forked it to
[gomail](https://github.com/muquit/gomail) (starting from mailsend-go v1.0.11, Feb-14-2025).
The main purpose of this fork is to add XOAUTH2 support (Bug #68, TODO).

If you use [mailsend](https://github.com/muquit/mailsend), please consider
using mailsend-go as no new features will be added to 
[mailsend](https://github.com/muquit/mailsend).

If you have any question, request or suggestion, please enter it in the 
[Issues](https://github.com/muquit/mailsend-go/issues) with appropriate label.

Announcement: (Apr-11-2025) - SMTP XOAUTH2 support is added, will be released
as soon as I get some time.
