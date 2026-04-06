# Downloading and Installing

Pre-compiled `mailsend-go` binaries are available for the following platforms:

* Windows - 32 and 64 bit
* Linux - 64 bit (tgz)
* MacOS - 64 bit (tgz, Homebrew)
* Raspberry pi - 32 bit (tgz)

Please download the binaries from the @RELEASES@
page.  

Please add an @ISSUES@ if you would need binaries for any other         platforms.

Before installing, please make sure to verify the checksum.

When the tgz or zip archives are extracted they create a directory `mailsend-go-dir/` with the 
content.

**Example**

```bash
➤ tar -tvf bin/mailsend-go-v1.0.11-linux-amd64.d.tar.gz
-rw-r--r--  0 muquit staff    1084 Jan 16 20:10 mailsend-go-v1.0.11-linux-amd64.d/LICENSE.txt
-rw-r--r--  0 muquit staff   33880 Jan 16 20:10 mailsend-go-v1.0.11-linux-amd64.d/README.md
-rwxr-xr-x  0 muquit staff 5427384 Jan 16 20:10 mailsend-go-v1.0.11-linux-amd64.d/mailsend-go-v1.0.11-linux-amd64
-rw-r--r--  0 muquit staff   34185 Jan 16 20:10 mailsend-go-v1.0.11-linux-amd64.d/mailsend-go.1
-rw-r--r--  0 muquit staff     903 Jan 16 20:10 mailsend-go-v1.0.11-linux-amd64.d/platforms.txt
```

```bash
➤ unzip -l bin/mailsend-go-v1.0.11-windows-amd64.d.zip
Archive:  bin/mailsend-go-v1.0.11-windows-amd64.d.zip
  Length      Date    Time    Name
---------  ---------- -----   ----
     1084  01-16-2026 20:10   mailsend-go-v1.0.11-windows-amd64.d/LICENSE.txt
    33880  01-16-2026 20:10   mailsend-go-v1.0.11-windows-amd64.d/README.md
  5563904  01-16-2026 20:10   mailsend-go-v1.0.11-windows-amd64.d/mailsend-go-v1.0.11-windows-amd64.exe
    34185  01-16-2026 20:10   mailsend-go-v1.0.11-windows-amd64.d/mailsend-go.1
      903  01-16-2026 20:10   mailsend-go-v1.0.11-windows-amd64.d/platforms.txt
---------                     -------
  5633956                     5 files
```

After extracting the archive, copy the binary somewhere in your PATH. 
Example:
```bash
sudo /bin/cp -fv \
         mailsend-go-v1.0.11-linux-amd64.d/mailsend-go-v1.0.11-linux-amd64 \
         /usr/local/bin/mailsend-go
sudo /bin/cp -fv \
         mailsend-go-v1.0.11-linux-amd64.d/mailsend-go.1 \
         /usr/share/main/man1
```

## Installing using Homebrew on Mac

You will need to install @BREW@ first. Note: @BREW@ formula will be availbale
only for released version of `mailsend-go`

### Installing

First install the custom tap.

```
brew tap muquit/formulae
brew install mailsend-go
```
Or use auto-tap (installs in one command):
```bash
brew install muquit/formulae/mailsend-go
```

**Note:** If you previously used the old dedicated tap (`muquit/mailsend-go`),
 you may get an ambiguity error. Migrate to the new tap with:

```bash
 brew uninstall mailsend-go
 brew untap muquit/mailsend-go
 brew install muquit/formulae/mailsend-go
```

### Updating
```bash
brew upgrade mailsend-go
```

### Uninstalling
```bash
brew uninstall mailsend-go
```

To remove the tap:
```bash
brew untap muquit/formulae
```
