# httptop
Grab http stats from running server.

# usage
```bash
$ ngrep -d <dev> -w Byline port <port> | httptop
URL COUNT
URL COUNT
...
```

# todos
- [x] count request
- [x] tests
- [ ] include statuscode
- [ ] include count by statuscode
- [ ] check out terminal control chars
- [ ] replace regex w byte peek
