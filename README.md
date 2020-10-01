# concurrent-ssh
easily execute ssh commands against multiple hosts

### usage
you can pass target hosts as comma separated inputs with `--targets` or `-t` shorthand:
```bash
cssh --targets host1,host2
cssh -t host1,host2
```
or as a newline separated file with `--file` or `-f` shorthand:
```bash
cssh --file /some/file/path
cssh -f /some/file/path
```

you can also pass args to ssh itself like so:
```bash
cssh --targets host1,host2 -- -o StrictHostKeyChecking=no whoami
```

You can override default concurrency limit of 25 by passing in `--limit` or `-l` shorthand
```bash
cssh --targets host1,host2 --limit 1 -- date
```

### todo
- refactor to enable other tools (scp, rsync, etc)
- expansion of targets (something like `--targets 127.0.0.{1..3}`)
- better output (prepend hosts for clarity / blocks / colors)
- tests (lol)
- benchmark against pssh (anecdotally this is marginally faster on my machine with 1-10 hosts)
- use goreleaser to build/publish releases
