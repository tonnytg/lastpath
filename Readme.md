# LastPath
 
This is a simple way to save the path and restore again next time.<br/>
Back easily to where are you working.


# Install

`go build`

`cp lastpath /usr/local/bin/lastpath`

`alias lp='lastpath'`


###  Save

`lp save`

This commando will save current folder to back easilly next time.


### Move

`lp`

Without parameter it will move to last folder saved.

### Info

`lp info`

This commando will show you the last folder saved.

### TShoot

Any problem? Try to remove the file `~/.lastpath.txt` and try again.

At this momment it is the same like `cat ~/.filepath.txt`