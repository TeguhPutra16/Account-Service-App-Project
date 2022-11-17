package controllers

import (
	"be13/project/entities"
	"database/sql"
	"fmt"
	"log"
)

func TransferHistory(db *sql.DB, user int) []entities.Users {

	///////////Proses Mengambil data-data mengenai Transaksi transfer//////////////////
	result, errSelect := db.Query("select users.id ,users.name,transactions_tf.transaction_name,transactions_tf.id,transfers.transfer_amount,u.name,transactions_tf.created_at from users inner join transactions_tf on users.id=transactions_tf.user_id inner join transfers on transactions_tf.id=transfers.transaction_tf_id inner join users u on u.id=transfers.user_id where users.id=?", user)
	if errSelect != nil {

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
		}
		users = append(users, pengirim, penerima)

		fmt.Printf("->> \nSender Id:%d\nSender Name:%s\nTransaction Name: %s\nTransaction Id:%d\nTransfer Amount:%d\nRecipient Name:%s\nTransaction Time:%s \n->>", pengirim.Id, pengirim.Name, transaksi.Transaction_name, transaksi.Id, transfer.Transfer_amount, penerima.Name, transaksi.Created_at)

	}
	return users

}
