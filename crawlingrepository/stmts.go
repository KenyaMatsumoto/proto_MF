package crawlingrepository

func OfficeReadStmt() string {
	s := `SELECT
			office_name, updated_at
		FROM
			Users
		WHERE
			user_id = ?
	`
	return s
}

func DistinctBankIdAndBankNameStmt() string {
	s := `SELECT
				DISTINCT bank_id,
				bank_name,
				bank_status,
				last_commit_date
			FROM
				Banks
			WHERE
				user_id = ? and office_name = ?;`
	return s
}

func DetailStmt(startDay string, lastDay string) string {
	if startDay == "" {
		startDay = "2010-01-01"
	}
	if lastDay == "" {
		lastDay = "2030-12-31"
	}

	s := `SELECT
			bank_name,
			trading_content,
			payment,
			amount,
			trading_date
				FROM
					Details
				WHERE
					user_id = ? and bank_id = ?  and trading_date > '` + startDay + `' and trading_date < '` + lastDay + `'`
	return s
}

func DistinctCardIdAndCardNameStmt() string {
	s := `SELECT
				DISTINCT card_id,
				card_name,
				card_status,
				last_commit_date
			FROM
				Cards
			WHERE
				user_id = ? and office_name = ?`
	return s
}
