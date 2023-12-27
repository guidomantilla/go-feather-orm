module github.com/guidomantilla/go-feather-sql

go 1.21.1

require (
	github.com/DATA-DOG/go-sqlmock v1.5.1
	github.com/golang/mock v1.6.0
	github.com/guidomantilla/go-feather-commons v0.70.0
	github.com/jmoiron/sqlx v1.3.5
	github.com/sijms/go-ora/v2 v2.8.3
)

require (
	go.uber.org/mock v0.4.0 // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/tools v0.16.1 // indirect
)

replace github.com/guidomantilla/go-feather-commons => ../go-feather-commons
