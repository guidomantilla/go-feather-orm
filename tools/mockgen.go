package tools

//go:generate mockgen -package=dao 			-source ../pkg/dao/types.go 		-destination ../pkg/dao/mocks.go
//go:generate mockgen -package=datasource 	-source ../pkg/datasource/types.go 	-destination ../pkg/datasource/mocks.go
//go:generate mockgen -package=transaction	-source ../pkg/transaction/types.go -destination ../pkg/transaction/mocks.go
