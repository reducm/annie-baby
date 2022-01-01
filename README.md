# Annie Baby is a Annie wrapper

## dev

### run

```bash
# docker pg
$ docker-compose -f deploy/dev_compose.yml up -d

$ go get
$ go run main.go
```

### model schema generate

```bash
$ go generate ./ent   
```
