# xap - complete mpv controller

For now intended to be a lightweight alternative to mpd using mpv.

## components (design)

### libxap

a wrapper around mpv's commands, options and properties

TODO: need a list of implemented "endpoints" of mpv functions

    - [x] basic functionality in `pkg/com` (via `blang/mpv`)
    - [ ] mpv doc parser?
    - [ ] spawn/stop mpv instance
    - [ ] connect to mpv via unix/tcp socket

### xap cli

command line tool to interact with mpv via unix socket

    - [x] basic player ctl
    - [x] basic queue ctl
    - [ ] player settings
        - [x] switch audio device
        - [ ] softvol adjustments
    - [x] start/stop background mpv process
    - [x] use with custom socket

### more

- platform integration using libxap
- xap web ui

## additional features

- [ ] [beets](beets.io) integration - shuffle library
- ...
