create database account_service_app;

use account_service_app;

create table users(
id int not null primary key auto_increment,
name varchar(50),
gender enum("M","F"),
address varchar(100),
email varchar(100),
telp_number varchar(50),
password varchar(50),
balance int,
created_at datetime default current_timestamp,
updated_at datetime default current_timestamp
);

create table top_up(
transaction_id int not null primary key auto_increment,
top_up_amount int,
created_at datetime default current_timestamp,
updated_at datetime default current_timestamp,
constraint fk_Top_upTransactions foreign key (transaction_id) references transactions(id)
);

create table transactions(
id int not null primary key auto_increment,
user_id int,
transaction_name varchar(50),
transaction_date datetime,
created_at datetime default current_timestamp,
updated_at datetime default current_timestamp,
constraint fk_UsersTransactions foreign key (user_id) references users(id)
);


create table transfers(
transaction_id int not null primary key auto_increment,
user_id int,
transfer_amount int,
created_at datetime default current_timestamp,
updated_at datetime default current_timestamp,
constraint fk_TransferTransactions foreign key (transaction_id) references transactions(id),
constraint fk_TransferUsers foreign key (user_id) references users(id)
);



