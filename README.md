# depinst
depinst is supported tool for [golang/dep](https://github.com/golang/dep).

It builds to required packages in Gopkg.toml and saves binaries (mainly for CLI vendoring).

## Installation
```bash
go get -u github.com/orisano/depinst
```

## How to use
```bash
$ depinst -help
Usage of ./depinst:
  -dir string
    	bin directory (default "bin")
  -list
    	show required list
  -q	turn off output

$ cat Gopkg.toml
required = [
    "golang.org/x/tools/cmd/stringer",
    "github.com/rakyll/statik",
    "github.com/rubenv/sql-migrate/sql-migrate",
]

$ depinst
depinst: running [go build -o bin/stringer ./vendor/golang.org/x/tools/cmd/stringer] ...
depinst: running [go build -o bin/statik ./vendor/github.com/rakyll/statik] ...
depinst: running [go build -o bin/sql-migrate ./vendor/github.com/rubenv/sql-migrate/sql-migrate] ...

$ ls -1 ./bin
sql-migrate*
statik*
stringer*
```

## Author
Nao Yonashiro (@orisano)

## License
MIT