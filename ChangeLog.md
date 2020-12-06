## Contents
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


