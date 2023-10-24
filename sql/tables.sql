
CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

CREATE TABLE IF NOT EXISTS usuarios (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY, 
    name VARCHAR(60) NOT NULL, 
    nick VARCHAR(30) NOT NULL UNIQUE, 
    email VARCHAR(60) NOT NULL UNIQUE, 
    password VARCHAR(20) NOT NULL, 
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP()
) ENGINE=INNODB;