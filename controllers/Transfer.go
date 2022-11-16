package controllers

import (
	"be13/project/entities"
	"database/sql"
	"fmt"
	"log"
)

func Transfer(db *sql.DB, noTelp_kirim, noTelp_terima string, jumlahTf int) string {
	///////////////////////////////////SALDO PENGIRIM//////////////////////////////////////////////////////////

	result := db.QueryRow("SELECT id,balance FROM users where telp_number=?", noTelp_kirim)

	var pengirim entities.Users
	errScan := result.Scan(&pengirim.Id, &pengirim.Balance)
	if errScan != nil {
		if errScan == sql.ErrNoRows {
			log.Fatal("Id pengirim tidak ada")
		} else {
			log.Fatal("eror scan", errScan.Error())
		}

	}

	/////////////////////////////////////////////////////////////////////////////////////////
	if pengirim.Balance-jumlahTf < 0 {
		log.Fatal("Saldo tidak mencukupi")
	}
	////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////////SALDO PENERIMA//////////////////////////////////
	result1 := db.QueryRow("SELECT id,balance FROM users where telp_number=?", noTelp_terima)

	var penerima entities.Users
	errScan1 := result1.Scan(&penerima.Id, &penerima.Balance)
	if errScan1 != nil {
		if errScan1 == sql.ErrNoRows {
			log.Fatal("Id penerima tidak ada")
		} else {
			log.Fatal("eror scan", errScan.Error())
		}

	}

	////////////////////////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////TAMBAH SALDO PENERIMA////////////////////////////

	var query = "UPDATE users set balance=? where telp_number=?"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("error prepare", errPrepare.Error())

	}

	current_balance := penerima.Balance + jumlahTf

	result2, errExec := statement.Exec(current_balance, noTelp_terima)
	if errExec != nil {
		log.Fatal("error exec insert", errExec.Error())
		// return errExec
	} else {
		row, _ := result2.RowsAffected()
		if row > 0 {
			// fmt.Println("Transfer berhasil diterima oleh Nomor :", noTelp_terima)
			// return ""
		} else {
			fmt.Println("Tambah saldo Gagal")
		}
	}

	/////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////KURANG SALDO PENGIRIM///////////////////////////////
	var query1 = "UPDATE users set balance=? where telp_number=?"
	statement1, errPrepare1 := db.Prepare(query1)
	if errPrepare1 != nil {
		log.Fatal("error prepare1", errPrepare1.Error())

	}

	current_balance1 := pengirim.Balance - jumlahTf

	result3, errExec := statement1.Exec(current_balance1, noTelp_kirim)
	if errExec != nil {
		log.Fatal("error exec insert", errExec.Error())
		// return errExec
	} else {
		row, _ := result3.RowsAffected()
		if row > 0 {
			// fmt.Println("kurang saldo pengirim")

		} else {
			log.Fatal("Transfer Gagal")
		}
	}

	/////////////////////////////////////////////////////INSERT KE TRANSAKSI///////////////////////
	var query2 = "insert into transactions (user_id,transaction_name) values (?,?)"
	statement2, errPrepare2 := db.Prepare(query2)
	if errPrepare2 != nil {
		log.Fatal("error prepare insert", errPrepare.Error())
		// return errPrepare2
	}

	result4, errExec := statement2.Exec(pengirim.Id, "Transfer")
	if errExec != nil {
		log.Fatal("error exec insert", errExec.Error())
		// return errExec
	} else {
		row, _ := result4.RowsAffected()
		if row > 0 {
			// fmt.Println("Transfer berhasil diterima oleh Nomor :", noTelp_terima)

		} else {
			fmt.Println("Failed to Add Account")
		}
	}
	////////////////////////////////////////////////////Insert ke Tabel Transfer///////////////////////////
	var query3 = "insert into transfers (user_id,transfer_amount) values (?,?)"
	statement3, errPrepare3 := db.Prepare(query3)
	if errPrepare3 != nil {
		log.Fatal("error prepare insert", errPrepare3.Error())
		// return errPrepare2
	}

	result5, errExec := statement3.Exec(penerima.Id, jumlahTf)
	if errExec != nil {
		log.Fatal("error exec insert", errExec.Error())
		// return errExec
	} else {
		row, _ := result5.RowsAffected()
		if row > 0 {
			fmt.Println("Transfer berhasil diterima oleh Nomor :", noTelp_terima)
		} else {
			fmt.Println("Failed to Add Account")
		}
	}

	return ""

}
