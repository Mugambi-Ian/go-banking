CREATE DATABASE banking;
USE banking;

DROP TABLE IF EXISTS `customers`;
CREATE TABLE `customers` (
  `customer_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `date_of_birth` date NOT NULL,
  `city` varchar(100) NOT NULL,
  `zipcode` varchar(10) NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`customer_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2006 DEFAULT CHARSET=latin1;


DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts` (
  `account_id` int(11) NOT NULL AUTO_INCREMENT,
  `customer_id` int(11) NOT NULL,
  `opening_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `account_type` varchar(10) NOT NULL,
  `amount` decimal(10,2) NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`account_id`),
  KEY `accounts_FK` (`customer_id`),
  CONSTRAINT `accounts_FK` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`customer_id`)
) ENGINE=InnoDB AUTO_INCREMENT=95471 DEFAULT CHARSET=latin1;


DROP TABLE IF EXISTS `transactions`;
CREATE TABLE `transactions` (
  `transaction_id` int(11) NOT NULL AUTO_INCREMENT,
  `account_id` int(11) NOT NULL,
  `amount` decimal(10,2) NOT NULL,
  `transaction_type` varchar(10) NOT NULL,
  `transaction_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`transaction_id`),
  KEY `transactions_FK` (`account_id`),
  CONSTRAINT `transactions_FK` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `username` varchar(20) NOT NULL,
  `password` varchar(20) NOT NULL,
  `role` varchar(20) NOT NULL,
  `customer_id` int(11) DEFAULT NULL,
  `created_on` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `refresh_token_store`;

CREATE TABLE `refresh_token_store` (
    `refresh_token` varchar(300) NOT NULL,
    created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`refresh_token`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;