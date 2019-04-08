# xap

bring some *mpd* to *mpv*

## xap cli

command line tool to interact with mpv via unix socket

- [x] player ctl
- [x] queue ctl
- [x] player settings
    - [x] softvol adjustments (via `xap vol`)
    - [x] access to raw mpv properties (via `xap raw [get|set|exec]`)
- [x] start/stop background mpv process (via `xap player [start|stop|show]`)
- [x] use with custom mpv socket (e.g. to control [IINA](https://github.com/lhc70000/iina),
    [gnome-mpv](https://github.com/gnome-mpv/gnome-mpv), etc.)

For documentation of the available commands check `xap --help`.

## installation

With a working [Go environment](https://golang.org/doc/install):
```
go get -u github.com/robertgzr/xap
```

## dynamic git-style subcommands

xap supports running external subcommands to extend its functionality.
For an example of how this can be done see [plugins/xap-rad-io](https://github.com/robertgzr/xap/master/plugins/xap-rad-io)

## more

### using xap with [beet](https://beets.io)
- add tracks/albums from your beets library: `beet ls -p <query> | xap add -`
- simple shuffle with beets: `beet random -p -n <num> [-e] [<query>] | xap add -`
