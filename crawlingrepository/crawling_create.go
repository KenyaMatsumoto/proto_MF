package crawlingrepository

// import (
// 	"database/sql"
// 	"log"
// 	"time"

// 	_ "github.com/go-sql-driver/mysql"
// )

// type DB interface {
// 	UserCreate(users []*User, userId string, updatedAt *time.Time) error
// 	BankCreate(userId string, banks []*Bank, today *time.Time) error
// 	DetailCreate(userId string, details []*Detail, today *time.Time) error
// }

// type db struct {
// 	Client *sql.DB
// }

// func NewDatabase() DB {
// 	client, err := sql.Open("mysql", "root@/freee")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return &db{
// 		Client: client,
// 	}
// }

// func (d *db) UserCreate(user []*User, userId string, updatedAt *time.Time) (err error) {
// 	for _, v := range user {
// 		v.UserId = userId

// 		updateStmt, err := d.Client.Prepare("UPDATE Users set userId = ?, lastId = ?,updatedAt = ? where userId = ?")
// 		if err != nil {
// 			return err
// 		}
// 		result, err := updateStmt.Exec(v.UserId, v.LastId, updatedAt)
// 		if err != nil {
// 			return err
// 		}

// 		rowsAffect, err := result.RowsAffected()
// 		if err != nil {
// 			return err
// 		}

// 		if rowsAffect == 0 {
// 			insertStmt, err := d.Client.Prepare("INSERT INTO Users(userId,lastId,updatedAt) VALUES(?, ?, ?)")
// 			if err != nil {
// 				return err
// 			}
// 			_, err = insertStmt.Exec(v.UserIdOfficeName, v.UserId, v.OfficeName, v.LastId, updatedAt)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}
// 	return nil
// }

// func (d *db) BankCreate(userId string, banks []*Bank, today *time.Time) error {
// 	for _, v := range banks {
// 		if v.Kind == "銀行口座" {

// 			insertStmt, err := d.Client.Prepare("INSERT INTO Banks(userId,bankId,LastCommitDate,officeName,bankName,amount,TopSyncStatus,updatedAt) VALUES(?, ?, ?, ?, ?, ?, ?, ?)")
// 			if err != nil {
// 				return err
// 			}
// 			_, err = insertStmt.Exec(userId, v.BankId, v.LastCommit, v.OfficeName, v.BankName, v.Amount, v.TopSyncStatus, today)
// 			if err != nil {
// 				return err
// 			}

// 		} else if v.Kind == "クレジットカード" {
// 			insertStmt, err := d.Client.Prepare("INSERT INTO Cards(userId,cardId,LastCommitDate,officeName,cardName,amount,TopSyncStatus,updatedAt) VALUES(?, ?, ?, ?, ?, ?, ?, ?)")
// 			if err != nil {
// 				return err
// 			}
// 			_, err = insertStmt.Exec(userId, v.BankId, v.LastCommit, v.OfficeName, v.BankName, v.Amount, v.TopSyncStatus, today)
// 			if err != nil {
// 				return err
// 			}
// 		} else {
// 			insertStmt, err := d.Client.Prepare("INSERT INTO Others(userId,otherId,LastCommitDate,officeName,otherName,amount, TopSyncStatus,updatedAt) VALUES(?, ?, ?, ?, ?, ?, ?, ?)")
// 			if err != nil {
// 				return err
// 			}
// 			_, err = insertStmt.Exec(userId, v.BankId, v.LastCommit, v.OfficeName, v.BankName, v.Amount, v.TopSyncStatus, today)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}
