package repository

const queryDDL = `
	CREATE TABLE IF NOT EXISTS ACCOUNTS (
		id INTEGER PRIMARY KEY,
		name TEXT,
		balance FLOAT
	);

	CREATE TABLE IF NOT EXISTS TRANSACTIONS (
		id INTEGER PRIMARY KEY,
		account_id INTEGER,
		date TEXT,
		amount FLOAT,
		is_credit INTEGER,
		FOREIGN KEY(account_id) REFERENCES ACCOUNTS(id)
	);
`
