# concurrent-ssh
easily execute ssh commands against multiple hosts

## usage
you can pass target hosts as comma separated inputs with `--targets` or `-t` shorthand (brace expansion is supported):
```bash
cssh --targets host{1..3},host10
cssh -t host1,host2,host3,host10
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

override default concurrency limit of 25 by passing in `--limit` or `-l` shorthand
```bash
cssh --targets host1,host2 --limit 1 -- date
```

## todo
- tests (lol)
- benchmark against pssh (anecdotally this is marginally faster on my machine with 1-10 hosts)
- use goreleaser to build/publish releases
