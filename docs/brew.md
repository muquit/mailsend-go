## Installing using Homebrew on Mac

You will need to install @BREW@ first. Note: @BREW@ formula will be avilable
only for released versions of `mailsend-go`

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
