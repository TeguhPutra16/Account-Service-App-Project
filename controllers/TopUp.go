package controllers

import (
	"be13/project/entities"
	"database/sql"
	"fmt"
	"log"
)

func BalanceTopUp(db *sql.DB, id int, topupamount int) string {
	result0 := db.QueryRow("select telp_number from users where id = ?", id)

	var topupuserNumber entities.Users
	errScan0 := result0.Scan(&topupuserNumber.Telp_number)
	if errScan0 != nil {
		if errScan0 == sql.ErrNoRows {
			log.Fatal("Sender Id does not exist")
		} else {
			log.Fatal("eror scan", errScan0.Error())
		}
	}
	// cek id dan balance dari user sebelum topup dengan telp number
	result := db.QueryRow("select id, balance from users where telp_number = ?", topupuserNumber.Telp_number)
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
		fmt.Println("")
	}

	// penambahan saldo setelah di topup
	var query = "update users set balance = ? where telp_number = ?"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("error prepare", errPrepare.Error())
	}

	balanceAfterTopUp := topupUser.Balance + topupamount

	result2, errExec := statement.Exec(balanceAfterTopUp, topupuserNumber.Telp_number)
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
	var query2 = "insert into transactions_tp (user_id, transaction_name) values (?,?)"
	// now := time.Now()
	statement2, errPrepare2 := db.Prepare(query2)
	if errPrepare2 != nil {
		log.Fatal("error prepare insert", errPrepare2.Error())
	}

	result3, errExec2 := statement2.Exec(topupUser.Id, "Top_Up")
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
		log.Fatal("error exec insert", errExec3.Error())
	} else {
		row, _ := result4.RowsAffected()
		if row > 0 {
			fmt.Println("")
		} else {
			fmt.Println("")
		}
	}
	return ""
}
