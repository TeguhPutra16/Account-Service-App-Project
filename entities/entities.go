package entities

import "time"

type Users struct {
	Id          int
	Name        string
	Email       string
	Password    string
	Address     string
	Telp_number string
	Balance     int
	Gender      string
	Created_at  time.Time
	Updated_at  time.Time
}

type Transaction_tp struct {
	Id               int
	User_id          int
	Transaction_name string
	Created_at       time.Time
}
type Transaction_tf struct {
	Id               int
	User_id          int
	Transaction_name string
	Created_at       time.Time
}

type Top_up struct {
	Transaction_id int
	Top_up_amount  int
}

type Transfer struct {
	Transaction_id  int
	User_id         int
	Transfer_amount int
}
