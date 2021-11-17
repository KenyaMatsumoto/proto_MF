create TABLE Users (
    id int auto_increment primary key not null,
    office_name char(255) not null,
    user_id char(255) not null,
    updated_at timestamp not null
);

create TABLE Banks(
    id int auto_increment primary key not null,
    user_id char(255) not null,
    bank_id char(255) not null,
    bank_name char(255) not null,
    office_name char(255) not null,
    amount int not null,
    bank_status char(255) not null,
    last_commit_date char(255) not null,
    updated_at timestamp not null
);

create TABLE Cards(
    id int auto_increment primary key not null,
    user_id char(255) not null,
    card_id char(255) not null,
    card_name char(255) not null,
    office_name char(255) not null,
    amount int not null,
    card_status char(255) not null,
    last_commit_date char(255) not null,
    updated_at timestamp not null
);


CREATE TABLE Details (
    id int auto_increment primary key not null,
    user_id char(255) not null,
    bank_name char(255) not null,
    trading_date timestamp not null,
    trading_content char(255),
    payment int,
    amount int,
    status char(255),
    transaction_number char(255),
    edit char(255), 
    crawling timestamp not null
);
