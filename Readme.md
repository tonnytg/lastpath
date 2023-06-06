# LastPath
 
This is a simple way to save the path and restore again next time.<br/>
Back easily to where are you working.

It is not possible to change of current folder with code, because children process can't change state of something control by operating system.
In this case, we create a alias called `lp` with alias of binary file `lastpath` and we can use it to save and restore the path.


# Install

```
go build -o lastpath cmd/main.go
cp lastpath /usr/local/bin/lastpath
alias lsp='cd $(lastpath)'
```

### SAVE

`lastpath save`

This commando will save current folder to back easilly next time.


### Move

`lp`

Without parameter it will move to last folder saved.

### Info

`lastpath info`

This commando will show you the last folder saved.

### TShoot

Any problem? Try to remove the file `~/.lastpath.txt` and try again.

At this momment it is the same like `cat ~/.filepath.txt`
