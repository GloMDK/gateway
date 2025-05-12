CREATE SCHEMA transactions;
CREATE TABLE transactions.transactions(
    id integer PRIMARY KEY,
    currency_code integer,
    amount float,
    bank_name varchar(300),
    status smallint,
    created_at timestamp default current_timestamp
);
