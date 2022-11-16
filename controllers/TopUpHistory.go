package controllers

import (
	"be13/project/entities"
	"database/sql"
	"fmt"
	"log"
)

func TopUpHistory(db *sql.DB, user entities.Users) {
	//lihat history topup berdasarkan id yg melakukan topup
	var transaction entities.Transaction
	var topup entities.Top_up
	result := db.QueryRow("select transactions.id, transactions.transaction_name, transactions.created_at, top_up.top_up_amount, users.balance, users.name as topupUser, users.id as topupUserId from users inner join transactions on transactions.user_id = users.id inner join top_up on top_up.transaction_id = transactions.id where users.id = ?", user.Id)
	errScan := result.Scan(&transaction.Id, &transaction.Transaction_name, &transaction.Created_at, &topup.Top_up_amount, &user.Balance, &user.Name, &user.Id)
	if errScan != nil {
		log.Fatal("top up history not found", errScan.Error())
	}
	fmt.Printf("\nTransaction Id: %d, Transaction Name: %s, Transaction Date: %s, TopUp Amount: %d, Current Balance: %d, TopUpUser Name: %s, TopUpUser Id: %d", transaction.Id, transaction.Transaction_name, transaction.Created_at, topup.Top_up_amount, user.Balance, user.Name, user.Id)
}
