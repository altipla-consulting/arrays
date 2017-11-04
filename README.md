
# arrays

[![GoDoc](https://godoc.org/github.com/altipla-consulting/arrays?status.svg)](https://godoc.org/github.com/altipla-consulting/arrays)

> Models for integer and string arrays in MySQL.


### Install

```shell
go get github.com/altipla-consulting/arrays
```

This library has no external dependencies outside the Go standard library.


### Usage

You can use the types of this package in your models structs when working with `database/sql`:

```go
type MyModel struct {
  ID    int64             `db:"id,omitempty"`
  Foo   arrays.Integers32 `db:"foo"`
  Bar   arrays.Integers64 `db:"bar"`
  Codes arrays.Strings    `db:"codes"`
}
```


### Contributing

You can make pull requests or create issues in GitHub. Any code you send should be formatted using ```gofmt```.


### Running tests

Start the test database:

```shell
docker-compose up -d database
```

Install test libs:

```shell
go get github.com/stretchr/testify
go get upper.io/db.v3
```

Run the tests:

```shell
go test
```

Shutdown the database when finished testing:

```shell
docker-compose stop database
```


### License

[MIT License](LICENSE)
