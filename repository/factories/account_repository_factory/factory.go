package account_repository_factory

import (
	account_repository2 "transactions_reader_stori/repository/commands/account_repository"
	"transactions_reader_stori/repository/init_repositories"
)

func NewAccountDatabaseRepo(db *init_repositories.DatabaseRepo) account_repository2.AccountRepository {
	return &account_repository2.AccountDatabaseRepo{Db: db.Db}
}
