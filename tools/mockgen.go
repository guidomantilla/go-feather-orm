package tools

//go:generate mockgen -source ../pkg/feather-relational-dao/types.go -destination ../pkg/feather-relational-dao/mocks.go -package=feather_relational_dao
//go:generate mockgen -source ../pkg/feather-relational-datasource/types.go -destination ../pkg/feather-relational-datasource/mocks.go -package=feather_relational_datasource
//go:generate mockgen -source ../pkg/feather-relational-tx/types.go -destination ../pkg/feather-relational-tx/mocks.go -package=feather_relational_tx
