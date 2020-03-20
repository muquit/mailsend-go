# v1.0.8
* One line message can have embedded new line with \n. If \n is found, it will
be repalced with real new line. Example: 
  body -msg "This is line1.\nThis is line2." 
The message will look like:
This is line1.
This is line2.

Fix Issue #22

(Mar-17-2020)

# v1.0.7
* If -q was used with -info, the messages were still printed on stdout.

Fix Issue #19

(Feb-16-2020)

# v1.0.6
* Add the flag -cs charser to specify a character set for text or HTML.
The default character set is utf-8.

Fix Issue #12

(Oct-27-2019)

# v1.0.5
* Add the flag -log filePath to write log messages to this path.

Fix Issue #5

(Jul-06-2019)

# v1.0.4

* The To address specified with -t was ignored when a list file was specified with -list. 

* Start a line with # to specify comment in the adderss list file.

Fix Issue #6

* Add binaries for Raspberr pi. Tested on Raspberry pi 3 b+

* Add the flag -rt to specify reply-to header. Thaks to Dominik G.

(Mar-26-2019)

# v1.0.3

* Code for comma separated to, cc and bcc was not there.
Fix Issue #2

* Remove -tname option for now. It creates trouble if multiple recipients are specifed.

(Feb-20-2019)

# v1.0.2

* Supply compiled binary for 32 bit Windows. No code change.

(Feb-14-2019)

# v1.0.1

* Released

(Feb-20-2019)

