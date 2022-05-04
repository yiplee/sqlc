# sqlc: A simple dynamic query builder for [kyleconroy/sqlc](https://github.com/kyleconroy/sqlc)

## How to use

```go
package sqlc_test

import (
    "context"
    "database/sql"
    "log"

    "github.com/yiplee/sqlc"
    "github.com/yiplee/sqlc/example"
)

func Example_build() {
    ctx := context.Background()

    db, err := sql.Open("postgres", "dsn")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Wrap the db in order to hook the query
    query := example.New(sqlc.Wrap(db))

    /*
    the original query must be simple and not contain any WHERE/ORDER/LIMIT/OFFSET clause

    -- name: ListAuthors :many
    SELECT * FROM authors;
    */

    // customize the builder
    authors, err := query.ListAuthors(sqlc.Build(ctx, func(b *sqlc.Builder) {
        b.Where("age > $1", 10)
        b.Where("name = $2", "foo")
        b.Order("age,name DESC")
        b.Limit(10)
    }))

    if err != nil {
        log.Fatalln("ListAuthors", err)
    }

    log.Printf("list %d authors", len(authors))
}
```


Not perfect, but enough for now.
