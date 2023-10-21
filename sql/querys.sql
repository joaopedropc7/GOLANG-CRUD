CREATE DATABASE IF NOT EXISTS academia;
use academia;

CREATE TABLE cliente(
    id int auto_increment primary key,
    nome varchar(100) not null,
    email varchar(100) not null,
    numerotelefone varchar(10) not null,
    criadoem timestamp default current_timestamp()
)ENGINE=INNODB;