package controllers

import (
	"be13/project/entities"
	"database/sql"
	"fmt"
	"log"
)

func Transfer(db *sql.DB, id int, noTelp_terima, noTelp_kirim string, jumlahTf int) string {

	//////Mengambil data no telepon pengirim dengan menggunakan user id dari fungsi login///////////////////////////////
	result0 := db.QueryRow("SELECT telp_number FROM users where id=?", id)

	var no_kirim entities.Users
	errScan0 := result0.Scan(&no_kirim.Telp_number)
	if errScan0 != nil {
		if errScan0 == sql.ErrNoRows {
			log.Fatal("Sender Id does not exist")
		} else {
			log.Fatal("eror scan", errScan0.Error())
		}
	}
	/////Konfirmasi nomor pemgirim jika salah akan kembali ke menu////////////////////////////////////////////////////////////////////////////
	if noTelp_kirim == no_kirim.Telp_number {

	} else {

		fmt.Println("==============================\nyour phone number is incorrect\nPlease Try again\n==============================")
		return ""
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////
	///////////////////////////////////MENGAMBIL DATA SALDO PENGIRIM//////////////////////////////////////////////////////////

	result := db.QueryRow("SELECT balance FROM users where id=?", id)

	var pengirim entities.Users
	errScan := result.Scan(&pengirim.Balance)
	if errScan != nil {
		if errScan == sql.ErrNoRows {
			log.Fatal("Sender Id does not exist")
		} else {
			log.Fatal("eror scan", errScan.Error())
		}
	}

	/////////////////////////////////////////////////////////////////////////////////////////
	if pengirim.Balance-jumlahTf < 0 {
		log.Fatal("Not Enough Balance")
	}
	////////////////////////////////////////////////////////////////////////////////////////
	///////////////////////////////////AMBIL DATA SALDO PENERIMA//////////////////////////////////
	result1 := db.QueryRow("SELECT id,balance FROM users where telp_number=?", noTelp_terima)

	var penerima entities.Users
	errScan1 := result1.Scan(&penerima.Id, &penerima.Balance)
	if errScan1 != nil {
		if errScan1 == sql.ErrNoRows {
			log.Fatal("Sender Id does not exist")
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
			//lanjut"
		} else {
			fmt.Println("Add Balance Failed")
		}
	}

	/////////////////////////////////////////////////////////////////////////////////////////
	////////////////////////////////////KURANG SALDO PENGIRIM///////////////////////////////
	var query1 = "UPDATE users set balance=? where id=?"
	statement1, errPrepare1 := db.Prepare(query1)
	if errPrepare1 != nil {
		log.Fatal("error prepare1", errPrepare1.Error())

	}

	current_balance1 := pengirim.Balance - jumlahTf

	result3, errExec := statement1.Exec(current_balance1, id)
	if errExec != nil {
		log.Fatal("error exec insert", errExec.Error())
		// return errExec
	} else {
		row, _ := result3.RowsAffected()
		if row > 0 {
			//lanjut

		} else {
			log.Fatal("Transfer Failed")
		}
	}

	///////////////////////////INSERT user id(pengirim) dan jenis transaksi KE tabel TRANSAKSI///////////////////////
	var query2 = "insert into transactions_tf (user_id,transaction_name) values (?,?)"
	statement2, errPrepare2 := db.Prepare(query2)
	if errPrepare2 != nil {
		log.Fatal("error prepare insert", errPrepare.Error())
		// return errPrepare2
	}

	result4, errExec := statement2.Exec(id, "Transfer")
	if errExec != nil {
		log.Fatal("error exec insert", errExec.Error())
		// return errExec
	} else {
		row, _ := result4.RowsAffected()
		if row > 0 {
			//lanjut

		} else {
			fmt.Println("Failed to Add Account")
		}
	}
	///////////////////////////////Insert user id(penerima) ke Tabel Transfer///////////////////////////
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

			fmt.Println("\nTransfer successfully received by Number :", noTelp_terima)
		} else {
			fmt.Println("Failed to Add Account")
		}
	}
	return ""
}
