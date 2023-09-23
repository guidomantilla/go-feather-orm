module github.com/guidomantilla/go-feather-sql

go 1.21.1

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/golang/mock v1.6.0
	github.com/guidomantilla/go-feather-commons v0.5.2
	github.com/sijms/go-ora/v2 v2.7.18
)

require (
	golang.org/x/mod v0.12.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
	golang.org/x/tools v0.13.0 // indirect
)

replace github.com/guidomantilla/go-feather-commons => ../go-feather-commons
