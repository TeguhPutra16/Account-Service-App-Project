package controllers

import (
	"be13/project/entities"
	"database/sql"
	"fmt"
	"log"
)

func TransferHistory(db *sql.DB, user int) []entities.Users {
	result, errSelect := db.Query("select users.id ,users.name,transactions_tf.transaction_name,transactions_tf.id,transfers.transfer_amount,u.name,transactions_tf.created_at from users inner join transactions_tf on users.id=transactions_tf.user_id inner join transfers on transactions_tf.id=transfers.transaction_tf_id inner join users u on u.id=transfers.user_id where users.id=?", user)
	if errSelect != nil {
		// return errSelect
		// return nil, errSelect
		log.Fatal("error select", errSelect.Error())
	}
	var users []entities.Users

	for result.Next() {
		var penerima entities.Users
		var pengirim entities.Users
		var transaksi entities.Transaction_tf
		var transfer entities.Transfer

		errScan := result.Scan(&pengirim.Id, &pengirim.Name, &transaksi.Transaction_name, &transaksi.Id, &transfer.Transfer_amount, &penerima.Name, &transaksi.Created_at)
		if errScan != nil {
			log.Fatal("error scan", errScan.Error())

			// return nil, errScan
		}
		users = append(users, pengirim, penerima)

		fmt.Printf("->> Sender Id:%d  Sender Name:%s  Transaction Name: %s  Transaction Id:%d  Transfer Amount:%d Recipient Name:%s Transaction Time:%s \n\n", pengirim.Id, pengirim.Name, transaksi.Transaction_name, transaksi.Id, transfer.Transfer_amount, penerima.Name, transaksi.Created_at)

	}
	return users

}
