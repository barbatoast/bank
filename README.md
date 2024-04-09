# bank
Example CRUD App

## Database Migration

```bash
cd migration
touch db.sqlite3
sqlite3 db.sqlite3 < create.sql
sqlite3 db.sqlite3 < insert.sql
```

## Run

```bash
go install
go build
./bank
```

## REST Endpoint

```bash
curl -s http://localhost:8080/api/accounts/test
```
