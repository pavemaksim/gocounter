# Gocounter

Very simple counter written in Go. It allows you to track total count for some unique ID.

# Requirements

Go 1.7+
MySQL 5.6+

# Deploy

- Grab source code `git clone`
- Create config file: `cp config.example.toml config.toml`
- Set up your DB connection via `DBDsn` parameter in `config.toml`
- `go run main.go` and you're done!

# API

Track count for ID:

`GET /counter/:id`

Response

```
{
    "status" : "ok"
}
```
