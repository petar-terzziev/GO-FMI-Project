# Golang course app
A project for the Go course at FMI which is the backend to a website for the course itself
## Run

- Setup a PostgreSQL db
- You need to setup your connection string in the two files, vulnerable-db.go, and migrations.go
- Start migration by commenting migration in main.go and commenting API
- Type  (that will migrate to your db):
```
go run main.go
```
- Comment migration and uncomment api
- Type:
```
go run main.go
```