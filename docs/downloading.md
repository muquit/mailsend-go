# Downloading and Installing

Pre-compiled `mailsend-go` binaries for Windows, Linux and MacOS can be
downloaded from the [releases](https://github.com/muquit/mailsend-go/releases)
page.  

| Files | Platform |
| :-------| :--------|
| mailsend-go_1.0.1_checksums.txt| SHA256 checksum files for the binaries|
| mailsend-go_1.0.1_windows_64-bit.zip | Windows 64 bit |
| mailsend-go_1.0.1_linux_64-bit.tar.gz | Linux 64 bit|
| mailsend-go_1.0.1_mac_64-bit.tar.gz |  Mac OS X 64 bit |
| mailsend-go_linux_64-bit.rpm | RPM for Linux 64 bit |
| mailsend-go_linux_64-bit.deb | Debian package for Linux 64 bit |

Before installing, please make sure to verify the checksum.

When the tgz or zip archives are extracted they create a directory `mailsend-go-dir/` with the 
content.

**Example**

```
    $ tar -tvf mailsend-go_1.0.1_linux_64-bit.tar.gz
	-rw-r--r--  0 muquit staff    1081 Jan 26 15:21 mailsend-go-dir/LICENSE.txt
	-rw-r--r--  0 muquit staff   14242 Jan 27 13:47 mailsend-go-dir/README.md
	-rw-r--r--  0 muquit staff   16866 Jan 27 13:47 mailsend-go-dir/docs/mailsend-go.1
	-rwxr-xr-x  0 muquit staff 5052992 Feb  9 19:23 mailsend-go-dir/mailsend-go
```

```
	$ unzip -l mailsend-go_1.0.1_windows_64-bit.zip
	Archive:  mailsend-go_1.0.1_windows_64-bit.zip
	  Length      Date    Time    Name
	---------  ---------- -----   ----
		 1081  01-26-2019 15:21   mailsend-go-dir/LICENSE.txt
		14242  01-27-2019 13:47   mailsend-go-dir/README.md
		16866  01-27-2019 13:47   mailsend-go-dir/docs/mailsend-go.1
	  4933632  02-09-2019 19:23   mailsend-go-dir/mailsend-go.exe
	---------                     -------
	  4965821                     4 files
```

## Installing using Homebrew on Mac

You will need to install [Homebrew](https://brew.sh/) first.

### Install

First install the custom tap.

```
    $ brew tap muquit/mailsend-go https://github.com/muquit/mailsend-go.git
    $ brew install mailsend-go
```

### Uninstall
```
    $ brew uninstall mailsend-go
```


## Installing the debian package on Ubuntu/Debian

### Inspect the package content
```
    $ dpkg -c mailsend-go_linux_64-bit.deb
	drwxr-xr-x 0/0               0 2019-02-10 20:17 usr/
	drwxr-xr-x 0/0               0 2019-02-10 20:17 usr/local/
	drwxr-xr-x 0/0               0 2019-02-10 20:17 usr/local/share/
	drwxr-xr-x 0/0               0 2019-02-10 20:17 usr/local/share/docs/
	drwxr-xr-x 0/0               0 2019-02-10 20:17 usr/local/share/docs/mailsend-go/
	-rw-r--r-- 0/0            1081 2019-02-10 20:17 usr/local/share/docs/mailsend-go/LICENSE.txt
	drwxr-xr-x 0/0               0 2019-02-10 20:17 usr/local/bin/
	-rwxr-xr-x 0/0         5052992 2019-02-10 20:17 usr/local/bin/mailsend-go
	drwxr-xr-x 0/0               0 2019-02-10 20:17 usr/local/share/man/
	drwxr-xr-x 0/0               0 2019-02-10 20:17 usr/local/share/man/man1/
	-rw-r--r-- 0/0           20896 2019-02-10 20:17 usr/local/share/man/man1/mailsend-go.1
	-rw-r--r-- 0/0           19236 2019-02-10 20:17 usr/local/share/docs/mailsend-go/README.md
```

### Install

```
    $ sudo dpkg -i mailsend-go_linux_64-bit.deb 
	Selecting previously unselected package mailsend-go.
	(Reading database ... 4039 files and directories currently installed.)
	Preparing to unpack mailsend-go_linux_64-bit.deb ...
	Unpacking mailsend-go (1.0.1) ...
	Setting up mailsend-go (1.0.1) ...
    $ mailsend-go -V
    @(#) mailsend-go v1.0.1
```

### Uninstall

```
    $ sudo dpkg -r mailsend-go
```

## Install the RPM package

### Inspect the package content
```
    $ rpm -qlp mailsend-go_linux_64-bit.rpm
    /usr/local/bin/mailsend-go
    /usr/local/share/docs/mailsend-go/LICENSE.txt
    /usr/local/share/docs/mailsend-go/README.md
    /usr/local/share/man/man1/mailsend-go.1
```
### Install/Upgrade
```
    # rpm -Uvh mailsend-go_linux_64-bit.rpm
    # mailsend-go -V
    @(#) mailsend-go v1.0.1
```
### Uninstall
```
    # rpm -ev mailsend-go
```

## Install from archive

### Inspect the content
```
    $ tar -tvf mailsend-go_1.0.1_linux_64-bit.tar.gz
    -rw-r--r--  0 muquit staff    1081 Jan 26 15:21 mailsend-go-dir/LICENSE.txt
    -rw-r--r--  0 muquit staff   14242 Jan 27 13:47 mailsend-go-dir/README.md
    -rw-r--r--  0 muquit staff   16866 Jan 27 13:47 mailsend-go-dir/docs/mailsend-go.1
    -rwxr-xr-x  0 muquit staff 5052992 Feb  9 19:23 mailsend-go-dir/mailsend-go
```

```
    $ unzip -l mailsend-go_1.0.1_windows_64-bit.zip
    Archive:  mailsend-go_1.0.1_windows_64-bit.zip
      Length      Date    Time    Name
    ---------  ---------- -----   ----
     1081  01-26-2019 15:21   mailsend-go-dir/LICENSE.txt
    14242  01-27-2019 13:47   mailsend-go-dir/README.md
    16866  01-27-2019 13:47   mailsend-go-dir/docs/mailsend-go.1
      4933632  02-09-2019 19:23   mailsend-go-dir/mailsend-go.exe
    ---------                     -------
      4965821                     4 files
```

### Install Linux
```
    $ tar -xf mailsend-go_1.0.1_linux_64-bit.tar.gz
    $ sudo cp mailsend-go-dir/mailsend-go /usr/local/bin
    $ sudo cp mailsend-go-dir/doc/mailsend-go.1 /usr/local/share/man/man1
```

### Install Windows

Unzip mailsend-go_1.0.1_windows_64-bit.zip and copy `mailsend-go-dir\mailsend-go.exe` somewhere in your PATH or run it from the directory.
