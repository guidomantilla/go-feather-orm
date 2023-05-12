module github.com/guidomantilla/go-feather-sql

go 1.18

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/golang/mock v1.6.0
	github.com/guidomantilla/go-feather-commons v0.2.3
	github.com/sijms/go-ora/v2 v2.7.5
	go.uber.org/zap v1.24.0
)

require (
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/mod v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/tools v0.9.1 // indirect
)

replace github.com/guidomantilla/go-feather-commons => ../go-feather-commons
