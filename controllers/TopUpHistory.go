package controllers

import (
	"be13/project/entities"
	"database/sql"
	"fmt"
	"log"
)

func TopUpHistory(db *sql.DB, user int) []entities.Users {
	//lihat history topup berdasarkan id yg melakukan topup
	result, errSelect := db.Query("select transactions_tp.id, transactions_tp.transaction_name, transactions_tp.created_at, top_up.top_up_amount, users.balance, users.name as topupUser, users.id as topupUserId from users inner join transactions_tp on transactions_tp.user_id = users.id inner join top_up on top_up.transaction_tp_id = transactions_tp.id where users.id = ?", user)
	if errSelect != nil {
		log.Fatal("error select", errSelect.Error())
	}
	var history []entities.Users

	for result.Next() {
		var user entities.Users
		var transaction entities.Transaction_tp
		var topup entities.Top_up
		errScan := result.Scan(&transaction.Id, &transaction.Transaction_name, &transaction.Created_at, &topup.Top_up_amount, &user.Balance, &user.Name, &user.Id)
		if errScan != nil {
			log.Fatal("error scan", errScan.Error())
		}
		history = append(history, user)
		fmt.Printf("\nTransaction Id: %d\nTransaction Name: %s\nTransaction Date: %s\nTopUp Amount: %d\nCurrent Balance: %d\nTopUpUser Name: %s\nTopUpUser Id: %d\n", transaction.Id, transaction.Transaction_name, transaction.Created_at, topup.Top_up_amount, user.Balance, user.Name, user.Id)
	}
	return history
}
