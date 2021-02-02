CREATE DATABASE IF NOT EXISTS user_db;
use user_db;

CREATE TABLE IF NOT EXISTS user (
    id int(10) unsigned not null auto_increment,
    name varchar(255) not null,
    mail_address varchar(1000) not null,
    created_time datetime not null default current_timestamp,
    updated_time datetime not null default current_timestamp on update current_timestamp,
    primary key (id)
);
