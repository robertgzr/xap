# xap

*For now* a simple alternative to MPD using mpv.

## xap cli

command line tool to interact with mpv via unix socket

- [x] basic player ctl
- [x] basic queue ctl
- [ ] player settings
    - [x] switch audio device (via `xap settings audio-device`)
    - [x] softvol adjustments (via `xap vol`)
    - [x] access to raw mpv properties (via `xap raw`)
- [x] start/stop background mpv process
- [x] use with custom mpv socket (e.g. to control [IINA](https://github.com/lhc70000/iina))

For documentation of the available commands check `xap --help`.

## _libxap \*planned\*_

wrapper around mpv's commands, options and properties (via [blang/mpv](https://github.com/blang/mpv))

- [x] basic functionality in `pkg/com`
- [ ] spawn/stop mpv instance
- [ ] connect to mpv via unix/tcp socket (find out if `blang/mpv` supports a way to do this)

## Installation

With a working [Go environment](https://golang.org/doc/install):
```
go get -u github.com/robertgzr/xap
```

## more

### tricks
- add tracks/albums from your beets library: `beet ls -p <query> | xap add -`
- simple shuffle with beets: `beet random -p -n <num> [-e] [<query>] | xap add -`

### additional features

- [ ] [beets](beets.io) integration: `shuffle` subcommand
- [ ] xap web ui?
- [ ] android app using libxap and a tcp socket?
