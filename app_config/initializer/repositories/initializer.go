package repositories

import (
	"transactions_reader_stori/repository/factories/account_repository_factory"
	"transactions_reader_stori/repository/factories/transaction_repository_factory"
	"transactions_reader_stori/repository/init_repositories"
)

type AppRepositoriesCommandsComponentsInitializerI interface {
	InitDatabaseRepoCommands() *init_repositories.DatabaseRepoCommands
}

type AppRepositoriesCommandsComponentsInitializer struct {
}

func (initializer AppRepositoriesCommandsComponentsInitializer) InitDatabaseRepoCommands() *init_repositories.DatabaseRepoCommands {
	repo := init_repositories.NewDatabaseRepo()

	transactionDatabaseRepo := transaction_repository_factory.NewTransactionDatabaseRepo(repo)
	accountDatabaseRepo := account_repository_factory.NewAccountDatabaseRepo(repo)

	commands := &init_repositories.DatabaseRepoCommands{
		TransactionRepository: transactionDatabaseRepo,
		AccountRepository:     accountDatabaseRepo,
	}

	return commands
}
