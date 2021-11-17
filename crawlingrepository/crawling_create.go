package crawlingrepository

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DB interface {
	UserCreate(user []*User, userId string, updatedAt *time.Time) error
	BankCreate(userId string, banks []*Bank, today *time.Time) error
	DetailCreate(userId string, details []*Detail, today *time.Time) error
}

type db struct {
	Client *sql.DB
}

func NewDatabase() DB {
	client, err := sql.Open("mysql", "root@/mf")
	if err != nil {
		log.Fatal(err)
	}
	return &db{
		Client: client,
	}
}

func (d *db) UserCreate(user []*User, userId string, updatedAt *time.Time) (err error) {
	for _, v := range user {
		v.UserId = userId

		updateStmt, err := d.Client.Prepare("UPDATE Users set office_name = ?, user_id = ?, updated_at = ? where office_name = ?")
		if err != nil {
			return err
		}
		result, err := updateStmt.Exec(v.officeName, v.UserId, updatedAt, v.officeName)
		if err != nil {
			return err
		}

		rowsAffect, err := result.RowsAffected()
		log.Println(rowsAffect)
		if err != nil {
			return err
		}

		if rowsAffect == 0 {
			insertStmt, err := d.Client.Prepare("INSERT INTO Users(office_name, user_id, updated_at) VALUES(?, ?, ?)")
			if err != nil {
				return err
			}
			_, err = insertStmt.Exec(v.officeName, v.UserId, updatedAt)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (d *db) BankCreate(userId string, banks []*Bank, today *time.Time) error {
	for _, v := range banks {
		if v.Kind == "銀行" {

			insertStmt, err := d.Client.Prepare("INSERT INTO Banks(user_id,bank_id,bank_name,office_name, amount,bank_status, last_commit_date,updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?)")
			if err != nil {
				return err
			}
			_, err = insertStmt.Exec(userId, v.BankId, v.BankName, v.OfficeName, v.Amount, v.BankStatus, v.ResitrationDate, today)
			if err != nil {
				return err
			}

		} else {
			insertStmt, err := d.Client.Prepare("INSERT INTO Cards(user_id,card_id,card_name,office_name, amount,card_status, last_commit_date,updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?)")
			if err != nil {
				return err
			}
			_, err = insertStmt.Exec(userId, v.BankId, v.BankName, v.OfficeName, v.Amount, v.BankStatus, v.ResitrationDate, today)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (d *db) DetailCreate(userId string, details []*Detail, today *time.Time) error {
	for _, v := range details {

		insertStmt, err := d.Client.Prepare("INSERT INTO Details(user_id, bank_id, bank_name, trading_date, trading_content, payment, amount, status, transaction_number, edit, crawling) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			return err
		}

		_, err = insertStmt.Exec(userId, v.BankId, v.BankName, v.TradingDate, v.TradingContent, v.Payment, v.Amount, v.Status, v.TransactionNumber, v.edit, today)
		if err != nil {
			return err
		}
	}
	return nil
}
