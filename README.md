# concurrent-ssh
![CI](https://github.com/rileymichael/concurrent-ssh/workflows/CI/badge.svg)
![Latest Release](https://img.shields.io/github/v/release/rileymichael/concurrent-ssh)

easily execute ssh commands against multiple hosts

## usage
pass in target hosts via one of the below flags, then pass any ssh options if needed & desired command to run
```bash
$ cssh --targets "root@host{1,2};root@host5" -- -o StrictHostKeyChecking=no date
[root@host5]
error: exit status 255
[root@host1]
Tue Jul 21 18:59:51 UTC 2020
[root@host2]
Tue Jul 21 18:59:51 UTC 2020
```

| flag           | description                                                         | example                                          |
|----------------|---------------------------------------------------------------------|--------------------------------------------------|
| --targets / -t | semicolon separated list of hosts (brace expansion supported)       | `cssh --targets "host{1..3};host10" -- date`     |
| --file / -f    | path to newline separated file of hosts (brace expansion supported) | `cssh --file /some/file/path -- date`            |
| --limit / -l   | concurrency limit (default 25)                                      | `cssh --targets "host1;host2" --limit 1 -- date` |
| --help / -f    | help                                                                | `cssh --help`                                    |

you must be able to ssh to all hosts without a password (it's using the BatchMode ssh option) 
so be sure to setup your ssh keys / ssh-agent appropriately

## todo
- tests (lol)
- benchmark against pssh (anecdotally this is marginally faster on my machine with 1-10 hosts)
- publish homebrew tap recipe w/goreleaser
