DROP SCHEMA IF EXISTS `test`;

CREATE SCHEMA `test`;

DROP TABLE IF EXISTS `test`.`orders`;

CREATE TABLE `test`.`orders` (
  `id` varchar(100) NOT NULL,
  `owner_id` varchar(100) DEFAULT NULL,
  `info` text,
  `date` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `test`.`users`;

CREATE TABLE `test`.`users` (
  `id` varchar(100) NOT NULL,
  `mail` varchar(45) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

