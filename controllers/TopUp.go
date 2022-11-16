package controllers

import (
	"be13/project/entities"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func BalanceTopUp(db *sql.DB, topupuserNumber string, topupamount int) {
	// cek id dan balance dari user sebelum topup dengan telp number
	result := db.QueryRow("select id, balance from users where telp_number = ?", topupuserNumber)
	var topupUser entities.Users
	errScan := result.Scan(&topupUser.Id, &topupUser.Balance)
	if errScan != nil {
		if errScan == sql.ErrNoRows {
			log.Fatal("topupUser not found")
		} else {
			log.Fatal("error scan", errScan.Error())
		}
	}
	if topupUser.Balance-topupamount < 0 {
		fmt.Println("no balance")
	}

	// penambahan saldo setelah di topup
	var query = "update users set balance = ? where telp_number = ?"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("error prepare", errPrepare.Error())
	}

	balanceAfterTopUp := topupUser.Balance + topupamount

	result2, errExec := statement.Exec(balanceAfterTopUp, topupuserNumber)
	if errExec != nil {
		log.Fatal("error exec insert", errExec.Error())
	} else {
		row, _ := result2.RowsAffected()
		if row > 0 {
			fmt.Println("topup successful")
		} else {
			fmt.Println("topup failed")
		}
	}

	// melengkapi tabel transaction
	var query2 = "insert into transactions (user_id, transaction_name, transaction_date) values (?,?,?)"
	now := time.Now()
	statement2, errPrepare2 := db.Prepare(query2)
	if errPrepare2 != nil {
		log.Fatal("error prepare insert", errPrepare2.Error())
	}

	result3, errExec2 := statement2.Exec(topupUser.Id, "topup", now)
	if errExec2 != nil {
		log.Fatal("error exec insert", errExec2.Error())
	} else {
		row, _ := result3.RowsAffected()
		if row > 0 {
			fmt.Println("")
		} else {
			fmt.Println("")
		}
	}

	// melengkapi tabel topup
	var query3 = "insert into top_up (top_up_amount) values (?)"
	statement3, errPrepare3 := db.Prepare(query3)
	if errPrepare3 != nil {
		log.Fatal("error prepare insert", errPrepare3.Error())
	}

	result4, errExec3 := statement3.Exec(topupamount)
	if errExec3 != nil {
		log.Fatal("error exec insert", errExec.Error())
	} else {
		row, _ := result4.RowsAffected()
		if row > 0 {
			fmt.Println("")
		} else {
			fmt.Println("")
		}
	}
}
