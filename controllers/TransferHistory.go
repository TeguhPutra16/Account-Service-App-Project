package controllers

import (
	"be13/project/entities"
	"database/sql"
	"fmt"
	"log"
)

func TransferHistory(db *sql.DB, user entities.Users) {
	result, errSelect := db.Query("select users.id ,users.name,transactions.transaction_name,transactions.id,transfers.transfer_amount,u.name,transactions.created_at from users inner join transactions on users.id=transactions.user_id inner join transfers on transactions.id=transfers.transaction_id inner join users u on u.id=transfers.user_id where users.id=?", user.Id)
	if errSelect != nil {
		log.Fatal("error select", errSelect.Error())
		// return nil, errSelect
	}

	// var dataUser []entities.Users
	for result.Next() {
		var penerima entities.Users
		var pengirim entities.Users
		var transaksi entities.Transaction
		var transfer entities.Transfer
		errScan := result.Scan(&pengirim.Id, &pengirim.Name, &transaksi.Transaction_name, &transaksi.Id, &transfer.Transfer_amount, &penerima.Name, &transaksi.Created_at)
		if errScan != nil {
			log.Fatal("eror scan", errScan.Error())
			// return nil, errScan
		}
		fmt.Printf("->> Id Pengirim:%d  Nama Pengirim:%s  Nama transaksi: %s  Id Transaksi:%d  Jumlah Transfer:%d Nama Penerima:%s Waktu Transaksi:%s \n\n", pengirim.Id, pengirim.Name, transaksi.Transaction_name, transaksi.Id, transfer.Transfer_amount, penerima.Name, transaksi.Created_at)

	}

}
