# go-linkr

`go-linkr` is a small URL redirect service backed by SQLite. Request a link by its slug and the service redirects to its stored target.

> **Note:** This is a simple tool built to suit my personal need. It includes absolutely no fancy features, nor an admin interface

## Run

From the repository root:

```sh
go run main.go
```

By default, the server listens on port `8383` on all network interfaces and stores its database in `go-linkr-db.db` in the current directory. Set either value with command-line flags:

```sh
go run main.go -port 8080 -db ./data/links.db
```

The parent directory for the database file is created automatically. The `links` table and its slug index are also created on startup if they do not already exist.

## Add Links

There is no HTTP endpoint for creating or editing links. Start the service once to create the database and schema, then insert rows directly with the `sqlite3` command-line tool:

```sh
sqlite3 go-linkr-db.db "INSERT INTO links (slug, target) VALUES ('docs', 'https://go.dev/doc/');"
```

The table has this shape, with a unique index on slug:

```sql
CREATE TABLE links (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  slug TEXT NOT NULL,
  target TEXT NOT NULL
);
```

## Routes

- `GET /{slug}` looks up the slug and responds with a temporary (`302`) redirect. Targets beginning with `http://` or `https://` are used as-is; other targets are prefixed with `https://`.
- `GET /` responds with `400 Bad Request`.
- A slug that is not present in the database also responds with `400 Bad Request`.

For example, after adding the `docs` row above, visit `http://localhost:8383/docs` to redirect to the Go documentation. If you run the service with a custom `-db` path, use that same path for the SQLite command.

## Build

```sh
go build -o go-linkr .
```

## Helper Scripts

The `helper-scripts/` directory contains some scripts I use to run this inside a jail on my Truenas CORE (FreeBSD) installation.
